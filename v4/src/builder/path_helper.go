package builder

// Paths inside a builder use "/" segments, and a numeric segment points at an EXISTING list element —
// the same rule the PHP Builder::set applies. Without it, "item/0/quantity/value" would store a literal
// map key "0" where FHIR expects a list, which is valid JSON and rejected as the wrong type by the server.

func isIndexSegment(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func setPath(data map[string]interface{}, path string, v interface{}) {
	setIn(data, splitPath(path), v)
}

func setIn(node interface{}, seg []string, v interface{}) {
	if len(seg) == 0 {
		return
	}
	head := seg[0]

	if len(seg) == 1 {
		switch n := node.(type) {
		case map[string]interface{}:
			n[head] = v
		case []interface{}:
			if i, ok := indexOf(head); ok && i >= 0 && i < len(n) {
				n[i] = v
			}
		}
		return
	}

	switch n := node.(type) {
	case map[string]interface{}:
		child, ada := n[head]
		if !ada || child == nil {
			// A list element cannot be created by index: the caller must push it first, as PHP demands.
			if isIndexSegment(seg[1]) {
				return
			}
			child = map[string]interface{}{}
			n[head] = child
		}
		setIn(child, seg[1:], v)
	case []interface{}:
		i, ok := indexOf(head)
		if !ok || i < 0 || i >= len(n) {
			return
		}
		setIn(n[i], seg[1:], v)
	}
}

func pushPath(data map[string]interface{}, path string, v interface{}) {
	pushIn(data, splitPath(path), v)
}

func pushIn(node interface{}, seg []string, v interface{}) {
	if len(seg) == 0 {
		return
	}
	head := seg[0]

	if len(seg) == 1 {
		switch n := node.(type) {
		case map[string]interface{}:
			list, _ := n[head].([]interface{})
			n[head] = append(list, v)
		}
		return
	}

	switch n := node.(type) {
	case map[string]interface{}:
		child, ada := n[head]
		if !ada || child == nil {
			if isIndexSegment(seg[1]) {
				return
			}
			child = map[string]interface{}{}
			n[head] = child
		}
		pushIn(child, seg[1:], v)
	case []interface{}:
		i, ok := indexOf(head)
		if !ok || i < 0 || i >= len(n) {
			return
		}
		pushIn(n[i], seg[1:], v)
	}
}

func indexOf(s string) (int, bool) {
	if !isIndexSegment(s) {
		return 0, false
	}
	n := 0
	for _, r := range s {
		n = n*10 + int(r-'0')
	}
	return n, true
}

func splitPath(path string) []string {
	out := []string{}
	mulai := 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			out = append(out, path[mulai:i])
			mulai = i + 1
		}
	}
	return append(out, path[mulai:])
}
