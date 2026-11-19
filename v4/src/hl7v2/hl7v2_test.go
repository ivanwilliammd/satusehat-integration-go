package hl7v2

import (
	"encoding/json"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const adt = "MSH|^~\\&|SIMRS|RS-01|SATUSEHAT|KEMKES|20261118120000||ADT^A01^ADT_A01|MSG00001|P|2.5.1\r" +
	"EVN|A01|20261118115900\r" +
	"PID|1||0001234567890^^^NIK^NIK||ANGGRAENI^SITI^^^\r" +
	"PV1|1|I|WARD01^101^A\r" +
	"OBX|1|NM|8867-4^Heart rate^LN||78|/min|||||F\r"

func mustParse(t *testing.T, er7 string) *Message {
	t.Helper()
	m, err := Parse(er7)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return m
}

func TestParseSegmentsAndPaths(t *testing.T) {
	m := mustParse(t, adt)
	if len(m.Segments) != 5 {
		t.Fatalf("segments = %d, want 5", len(m.Segments))
	}
	for _, c := range []struct{ path, want string }{
		{"MSH.9.1", "ADT"},
		{"MSH.9.2", "A01"},
		{"MSH.10", "MSG00001"},
		{"PID.5.1", "ANGGRAENI"},
		{"OBX.5", "78"},
	} {
		if got := m.Get(c.path); got != c.want {
			t.Errorf("%s = %q, want %q", c.path, got, c.want)
		}
	}
}

func TestRoundTripIsByteForByte(t *testing.T) {
	if got := mustParse(t, adt).String(); got != adt {
		t.Fatalf("round trip mismatch:\n got %q\nwant %q", got, adt)
	}
}

func TestEscapesAreDecodedAndReencoded(t *testing.T) {
	m := mustParse(t, "MSH|^~\\&|A|B|C|D|20260101||ADT^A08|1|P|2.5.1\rPID|1||123||SMITH\\S\\JOHN\\F\\JR\\T\\II\r")
	if got := m.Get("PID.5.1"); got != "SMITH^JOHN|JR&II" {
		t.Fatalf("decoded = %q", got)
	}
	m.Set("PID.5.2", "O'BRIEN&SONS")
	if got := mustParse(t, m.String()).Get("PID.5.2"); got != "O'BRIEN&SONS" {
		t.Fatalf("re-encoded = %q", got)
	}
}

func TestMessageWithoutMSHIsRejected(t *testing.T) {
	if _, err := Parse("PID|1||123\r"); err != ErrNoMSH {
		t.Fatalf("err = %v, want ErrNoMSH", err)
	}
	if _, err := Parse("   "); err != ErrEmpty {
		t.Fatalf("err = %v, want ErrEmpty", err)
	}
}

func TestCharsetIsReportedNotApplied(t *testing.T) {
	if got := mustParse(t, adt).Charset(); got != "ASCII" {
		t.Fatalf("charset = %q", got)
	}
	if got := CharsetOf("UNICODE UTF-8"); got != "UTF-8" {
		t.Fatalf("charset = %q", got)
	}
}

func TestMllpFrameAndUnframe(t *testing.T) {
	if !strings.Contains(Unframe(Frame(adt)), "ADT^A01") {
		t.Fatal("unframed message lost its content")
	}
}

func TestMllpNeverEmitsTruncatedMessage(t *testing.T) {
	stream := Frame(adt) + Frame(strings.Replace(adt, "ADT^A01", "ADT^A08", 1))
	complete := map[string]bool{adt: true, strings.Replace(adt, "ADT^A01", "ADT^A08", 1): true}
	// The property that matters is that a partial stream never yields a shortened message; a cut that
	// lands after the second FS but before its CR legitimately yields both complete messages.
	for cut := 1; cut < len(stream); cut++ {
		for _, m := range Split(stream[:cut]) {
			if !complete[m] {
				t.Fatalf("cut %d produced a shortened message (%d bytes)", cut, len(m))
			}
		}
	}
	if len(Split(stream)) != 2 {
		t.Fatal("two frames must yield two messages")
	}
}

func TestReaderKeepsSecondFrameInSameChunk(t *testing.T) {
	r := NewReader()
	if err := r.Push(Frame("MSG-A") + Frame("MSG-B")); err != nil {
		t.Fatal(err)
	}
	if m, ok := r.Take(); !ok || m != "MSG-A" {
		t.Fatalf("first = %q ok=%v", m, ok)
	}
	if m, ok := r.Take(); !ok || m != "MSG-B" {
		t.Fatalf("second = %q ok=%v", m, ok)
	}
	if _, ok := r.Take(); ok {
		t.Fatal("buffer must be empty")
	}
}

func TestAckEchoesControlID(t *testing.T) {
	ack := BuildAck(mustParse(t, adt), Accept, "")
	if ack.MessageType() != "ACK" || ack.Get("MSA.1") != Accept || ack.Get("MSA.2") != "MSG00001" {
		t.Fatalf("ack = %q", ack.String())
	}
	nak := BuildNak(mustParse(t, adt), "PID missing")
	if nak.Get("MSA.1") != Error || nak.Get("MSA.3") != "PID missing" {
		t.Fatalf("nak = %q", nak.String())
	}
}

type manifest struct {
	Valid   map[string]struct{ Profile string }        `json:"valid"`
	Invalid map[string]struct {
		Profile   string   `json:"profile"`
		Issues    []string `json:"issues"`
		Exception string   `json:"exception"`
	} `json:"invalid"`
}

func loadManifest(t *testing.T) manifest {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("testdata", "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var m manifest
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	return m
}

func TestCorpusAcceptsWhatHospitalsSend(t *testing.T) {
	mf := loadManifest(t)
	if len(mf.Valid) == 0 {
		t.Fatal("empty corpus: refusing to pass with zero cases")
	}
	for name, spec := range mf.Valid {
		raw, err := os.ReadFile(filepath.Join("testdata", "valid", name+".hl7"))
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		msg := mustParse(t, string(raw))
		for _, issue := range Profiles[spec.Profile]().Validate(msg) {
			if issue.Severity == SeverityError {
				t.Errorf("%s: unexpected error %s", name, issue)
			}
		}
	}
}

func TestCorpusRefusesWhatMustNotReachADatabase(t *testing.T) {
	mf := loadManifest(t)
	if len(mf.Invalid) == 0 {
		t.Fatal("empty corpus: refusing to pass with zero cases")
	}
	for name, spec := range mf.Invalid {
		raw, err := os.ReadFile(filepath.Join("testdata", "invalid", name+".hl7"))
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		msg, err := Parse(string(raw))
		if spec.Exception != "" {
			if err == nil {
				t.Errorf("%s: expected an error", name)
			}
			continue
		}
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		paths := map[string]bool{}
		for _, issue := range Profiles[spec.Profile]().Validate(msg) {
			paths[issue.Path] = true
		}
		for _, want := range spec.Issues {
			if !paths[want] {
				t.Errorf("%s: expected issue at %s, got %v", name, want, paths)
			}
		}
	}
}

func TestClientOverRealSocket(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				reader := NewReader()
				buf := make([]byte, 8192)
				for {
					n, err := c.Read(buf)
					if n > 0 {
						_ = reader.Push(string(buf[:n]))
						for {
							raw, ok := reader.Take()
							if !ok {
								break
							}
							in := mustParse(t, raw)
							_, _ = c.Write([]byte(Frame(BuildAck(in, Accept, "").String())))
						}
					}
					if err != nil {
						return
					}
				}
			}(conn)
		}
	}()

	client := NewClient(ln.Addr().String(), 2*time.Second, 1)
	res, err := client.Send(mustParse(t, adt))
	if err != nil {
		t.Fatalf("send: %v", err)
	}
	if res.Code != Accept || res.ControlID != "MSG00001" || !res.Accepted {
		t.Fatalf("unexpected result %+v", res)
	}

	var msgs []*Message
	for _, id := range []string{"S1", "S2", "S3"} {
		msgs = append(msgs, mustParse(t, strings.Replace(adt, "MSG00001", id, 1)))
	}
	results, err := client.Session(msgs)
	if err != nil {
		t.Fatalf("session: %v", err)
	}
	if len(results) != 3 || results[2].ControlID != "S3" {
		t.Fatalf("unexpected results %+v", results)
	}
}

func TestClientRejectsStaleAck(t *testing.T) {
	msg := mustParse(t, adt)
	stale := BuildAck(msg, Accept, "").Set("MSA.2", "OTHER123").String()
	if _, err := (&Client{}).Interpret(msg, stale); err == nil {
		t.Fatal("stale ACK must be rejected")
	}
}

func TestClientFailsClearlyWhenNothingListens(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := ln.Addr().String()
	ln.Close()

	_, err = NewClient(addr, 300*time.Millisecond, 1).Send(mustParse(t, adt))
	if err == nil || !strings.Contains(err.Error(), "MLLP send failed") {
		t.Fatalf("err = %v", err)
	}
}
