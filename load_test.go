// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestLoadMap$
func TestLoadMap(t *testing.T) {
	entries := map[string]Value{
		"k1": Value("v1"),
		"k2": Value("v2"),
		"k3": Value("v3"),
	}

	props := LoadMap(entries)
	if len(props.entries) != len(entries) {
		t.Fatalf("len(props.entries) %d != len(entries) %d", len(props.entries), len(entries))
	}

	for key, value := range entries {
		got := props.entries[key]
		if got != value {
			t.Fatalf("got %s != value %s", got, value)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestLoadReader$
func TestLoadReader(t *testing.T) {
	str := "# xxx\r\n  k1    =     v1  \r\nk2    =     v2\r\nk3=v3\r\nk4=v4\nk5=v5"
	reader := strings.NewReader(str)

	want := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
		"k5": "v5",
	}

	props, err := LoadReader(reader)
	if err != nil {
		t.Fatal(err)
	}

	if len(props.entries) != len(want) {
		t.Fatalf("props.entries len %d != want len %d", len(props.entries), len(want))
	}

	for wantKey, wantValue := range want {
		value, ok := props.entries[wantKey]
		if !ok {
			t.Fatalf("key %s not found", wantKey)
		}

		if value.String() != wantValue {
			t.Fatalf("value %s is wrong", value)
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestLoadFile$
func TestLoadFile(t *testing.T) {
	str := "# xxx\r\n  k1    =     v1  \r\nk2    =     v2\r\nk3=v3\r\nk4=v4\nk5=v5"
	file := filepath.Join(t.TempDir(), t.Name())

	if err := os.WriteFile(file, []byte(str), 0644); err != nil {
		t.Fatal(err)
	}

	want := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
		"k5": "v5",
	}

	props, err := LoadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	if len(props.entries) != len(want) {
		t.Fatalf("props.entries len %d != want len %d", len(props.entries), len(want))
	}

	for wantKey, wantValue := range want {
		value, ok := props.entries[wantKey]
		if !ok {
			t.Fatalf("key %s not found", wantKey)
		}

		if value.String() != wantValue {
			t.Fatalf("value %s is wrong", value)
		}
	}
}
