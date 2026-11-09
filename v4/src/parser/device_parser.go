package parser

import (
	"fmt"
	"strings"
)

// DeviceParser extracts fields from a FHIR Device resource map.
type DeviceParser struct {
	device map[string]any
}

// NewDeviceParser creates a new DeviceParser instance.
func NewDeviceParser(device map[string]any) *DeviceParser {
	if device == nil {
		device = map[string]any{}
	}
	return &DeviceParser{device: device}
}

func (p *DeviceParser) GetIdentifiers() any {
	return p.device["identifier"]
}

func (p *DeviceParser) GetStatus() string {
	return strVal(p.device["status"])
}

func (p *DeviceParser) GetManufacturer() string {
	return strVal(p.device["manufacturer"])
}

func (p *DeviceParser) GetDeviceNames() any {
	return p.device["deviceName"]
}

func (p *DeviceParser) GetType() any {
	return p.device["type"]
}

// Shared helper functions across package parser.

func removePrefix(val string, prefix string) string {
	if val != "" && strings.HasPrefix(val, prefix) {
		return strings.TrimPrefix(val, prefix)
	}
	return val
}

func strVal(v any) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}

func mapVal(v any) map[string]any {
	if v == nil {
		return nil
	}
	if m, ok := v.(map[string]any); ok {
		return m
	}
	return nil
}

func sliceMaps(v any) []map[string]any {
	if v == nil {
		return nil
	}
	if s, ok := v.([]map[string]any); ok {
		return s
	}
	if s, ok := v.([]any); ok {
		out := make([]map[string]any, 0, len(s))
		for _, item := range s {
			if m, ok := item.(map[string]any); ok {
				out = append(out, m)
			}
		}
		return out
	}
	return nil
}

func nestedStr(root map[string]any, key1, key2 string) string {
	if root == nil {
		return ""
	}
	m := mapVal(root[key1])
	if m == nil {
		return ""
	}
	return strVal(m[key2])
}
