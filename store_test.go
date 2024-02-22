// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestStoreMap$
func TestStoreMap(t *testing.T) {
	props := New()
	props.Set("key1", "value1")
	props.Set("key2", "value2")
	props.Set("key3", "value3")

	got := StoreMap(props)
	if len(got) != props.Count() {
		t.Fatalf("len(got) %d != props.Count() %d", len(got), props.Count())
	}

	props.List(func(key string, value Value) (finished bool) {
		gotValue, ok := got[key]
		if !ok {
			t.Fatalf("key %s not found", key)
		}

		if gotValue != value {
			t.Fatalf("gotValue %s != value %s", gotValue, value)
		}

		return false
	})
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestStoreWriter$
func TestStoreWriter(t *testing.T) {
	props := New()
	props.Set("key1", "value1")
	props.Set("key2", "value2")
	props.Set("key3", "value3")

	buffer := bytes.NewBuffer(make([]byte, 0, 512))
	if err := StoreWriter(props, buffer); err != nil {
		t.Fatal(err)
	}

	got := buffer.String()
	strs := strings.Split(got, tokenLF)
	strs = strs[:props.Count()]

	sort.Strings(strs)

	got = strings.Join(strs, tokenLF)
	got = got + tokenLF

	want := "key1=value1\nkey2=value2\nkey3=value3\n"
	if got != want {
		t.Fatalf("got %q != want %q", got, want)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestStoreFile$
func TestStoreFile(t *testing.T) {
	props := New()
	props.Set("key1", "value1")
	props.Set("key2", "value2")
	props.Set("key3", "value3")

	file := filepath.Join(t.TempDir(), t.Name())
	if err := StoreFile(props, file); err != nil {
		t.Fatal(err)
	}

	bs, err := os.ReadFile(file)
	if err != nil {
		t.Log(err)
	}

	got := string(bs)
	strs := strings.Split(got, tokenLF)
	strs = strs[:props.Count()]

	sort.Strings(strs)

	got = strings.Join(strs, tokenLF)
	got = got + tokenLF

	want := "key1=value1\nkey2=value2\nkey3=value3\n"
	if got != want {
		t.Fatalf("got %q != want %q", got, want)
	}
}
