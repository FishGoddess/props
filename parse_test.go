// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import "testing"

// go test -v -cover -run=^TestParseLine$
func TestParseLine(t *testing.T) {
	key, value, err := parseLine(1, "  error  ")
	if err == nil {
		t.Fatal("parse error line returns nil")
	}

	key, value, err = parseLine(1, "  key      =                value  \r\n")
	if err != nil {
		t.Fatal(err)
	}

	if key != "key" {
		t.Fatalf("key %s is wrong", key)
	}

	if value != "value" {
		t.Fatalf("value %s is wrong", value)
	}
}

// go test -v -cover -run=^TestParseLine$
func TestParseLines(t *testing.T) {
	lines := []string{
		"# xxx",
		"  k1    =     v1  \r\n",
		"k2    =     v2\r\n",
		"k3=v3\r\n",
		"k4=v4\n",
		"k5=v5",
	}

	want := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
		"k5": "v5",
	}

	entries, err := parseLines(lines)
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) != len(want) {
		t.Fatalf("entries len %d != want len %d", len(entries), len(want))
	}

	for wantKey, wantValue := range want {
		value, ok := entries[wantKey]
		if !ok {
			t.Fatalf("key %s not found", wantKey)
		}

		if value.String() != wantValue {
			t.Fatalf("value %s is wrong", value)
		}
	}
}

// go test -v -cover -run=^TestParse$
func TestParse(t *testing.T) {
	str := "# xxx\r\n  k1    =     v1  \r\nk2    =     v2\r\nk3=v3\r\nk4=v4\nk5=v5"

	want := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
		"k5": "v5",
	}

	entries, err := parse(str)
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) != len(want) {
		t.Fatalf("entries len %d != want len %d", len(entries), len(want))
	}

	for wantKey, wantValue := range want {
		value, ok := entries[wantKey]
		if !ok {
			t.Fatalf("key %s not found", wantKey)
		}

		if value.String() != wantValue {
			t.Fatalf("value %s is wrong", value)
		}
	}
}
