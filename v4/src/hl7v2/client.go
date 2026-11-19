package hl7v2

import (
	"errors"
	"fmt"
	"net"
	"time"
)

// Result is the outcome of one message, whether the transport worked or not.
type Result struct {
	Code       string
	ControlID  string
	Text       string
	Attempts   int
	Latency    time.Duration
	Accepted   bool
}

// Client sends messages to an analyzer or LIS over MLLP.
//
// Two rules that matter operationally:
//   - transport failures retry, application verdicts do not: retrying an AE turns one bad message
//     into a thousand;
//   - MSA-2 must echo MSH-10, otherwise the ACK is stale or misrouted and must not count as success.
type Client struct {
	Address string
	Timeout time.Duration
	Retries int
}

func NewClient(address string, timeout time.Duration, retries int) *Client {
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	return &Client{Address: address, Timeout: timeout, Retries: retries}
}

// Send is Session for one message.
func (c *Client) Send(m *Message) (Result, error) {
	results, err := c.Session([]*Message{m})
	if err != nil {
		return Result{}, err
	}
	return results[0], nil
}

// Session sends several messages over one connection, one result per message, in order.
func (c *Client) Session(messages []*Message) ([]Result, error) {
	var results []Result
	var conn net.Conn
	var reader *Reader

	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()

	for _, msg := range messages {
		var lastErr error
		for attempt := 1; attempt <= c.Retries+1; attempt++ {
			if conn == nil {
				dialed, err := net.DialTimeout("tcp", c.Address, c.Timeout)
				if err != nil {
					lastErr = err
					time.Sleep(time.Duration(attempt) * 200 * time.Millisecond)
					continue
				}
				conn = dialed
				reader = NewReader()
			}
			res, err := c.exchange(conn, reader, msg)
			if err == nil {
				res.Attempts = attempt
				results = append(results, res)
				lastErr = nil
				break
			}
			lastErr = err
			_ = conn.Close()
			conn = nil
			if attempt <= c.Retries {
				time.Sleep(time.Duration(attempt) * 200 * time.Millisecond)
			}
		}
		if lastErr != nil {
			return results, fmt.Errorf("MLLP send failed at message %d after %d retries: %w", len(results)+1, c.Retries, lastErr)
		}
	}
	return results, nil
}

// Interpret parses an ACK received out-of-band; daemons that own their connection can reuse it.
func (c *Client) Interpret(sent *Message, ackRaw string) (Result, error) {
	ack, err := Parse(ackRaw)
	if err != nil {
		return Result{}, err
	}
	if ack.MessageType() != "ACK" {
		return Result{}, fmt.Errorf("reply is not an ACK but %s", ack.MessageType())
	}
	echo := ack.Get("MSA.2")
	if echo != "" && echo != sent.ControlID() {
		return Result{}, fmt.Errorf("MSA-2 (%s) does not match MSH-10 (%s): stale or misrouted ACK", echo, sent.ControlID())
	}
	code := ack.Get("MSA.1")
	if code == "" {
		code = Error
	}
	return Result{Code: code, ControlID: echo, Text: ack.Get("MSA.3"), Accepted: code == Accept}, nil
}

func (c *Client) exchange(conn net.Conn, reader *Reader, msg *Message) (Result, error) {
	started := time.Now()
	if err := conn.SetDeadline(time.Now().Add(c.Timeout)); err != nil {
		return Result{}, err
	}
	if _, err := conn.Write([]byte(Frame(msg.String()))); err != nil {
		return Result{}, err
	}
	ackRaw, err := c.readAck(conn, reader)
	if err != nil {
		return Result{}, err
	}
	res, err := c.Interpret(msg, ackRaw)
	if err != nil {
		return Result{}, err
	}
	res.Latency = time.Since(started)
	return res, nil
}

func (c *Client) readAck(conn net.Conn, reader *Reader) (string, error) {
	if msg, ok := reader.Take(); ok {
		return msg, nil
	}
	buf := make([]byte, 8192)
	for {
		n, err := conn.Read(buf)
		if n > 0 {
			if perr := reader.Push(string(buf[:n])); perr != nil {
				return "", perr
			}
			if msg, ok := reader.Take(); ok {
				return msg, nil
			}
		}
		if err != nil {
			return "", errors.New("no ACK received: " + err.Error())
		}
	}
}
