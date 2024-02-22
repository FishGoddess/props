// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import "testing"

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNew$
func TestNew(t *testing.T) {
	props := New()
	if props.entries == nil {
		t.Fatal("props entries is nil")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNewFromMap$
func TestNewFromMap(t *testing.T) {
	entries := map[string]Value{
		"k1": Value("v1"),
		"k2": Value("v2"),
		"k3": Value("v3"),
	}

	props := NewFromMap(entries)
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

// go test -v -cover -count=1 -test.cpu=1 -run=^TestPropsGet$
func TestPropsGet(t *testing.T) {
	key := "k"
	value := Value("v")

	props := New()
	props.entries[key] = value

	missed := props.Get("missed")
	if missed != "" {
		t.Fatal("got non-empty missed value")
	}

	got := props.Get(key)
	if got != value {
		t.Fatalf("got %s != value %s", got, value)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestPropsSet$
func TestPropsSet(t *testing.T) {
	key := "k"
	value := Value("v")

	props := New()
	props.Set(key, value)

	if len(props.entries) != 1 {
		t.Fatalf("entries len %d is wrong", len(props.entries))
	}

	got, ok := props.entries[key]
	if !ok {
		t.Fatalf("key %s not found", key)
	}

	if got != value {
		t.Fatalf("got %s != value %s", got, value)
	}
}
