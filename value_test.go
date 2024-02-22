// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import "testing"

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueString$
func TestValueString(t *testing.T) {
	want := "value"
	got := Value(want).String()

	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueInt$
func TestValueInt(t *testing.T) {
	got, err := Value("xxx").Int()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("54").Int()
	if err != nil {
		t.Fatal(err)
	}

	if got != 54 {
		t.Fatalf("got %d is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueInt8$
func TestValueInt8(t *testing.T) {
	got, err := Value("xxx").Int8()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("54").Int8()
	if err != nil {
		t.Fatal(err)
	}

	if got != 54 {
		t.Fatalf("got %d is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueInt16$
func TestValueInt16(t *testing.T) {
	got, err := Value("xxx").Int16()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("54").Int16()
	if err != nil {
		t.Fatal(err)
	}

	if got != 54 {
		t.Fatalf("got %d is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueInt32$
func TestValueInt32(t *testing.T) {
	got, err := Value("xxx").Int32()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("54").Int32()
	if err != nil {
		t.Fatal(err)
	}

	if got != 54 {
		t.Fatalf("got %d is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueInt64$
func TestValueInt64(t *testing.T) {
	got, err := Value("xxx").Int64()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("54").Int64()
	if err != nil {
		t.Fatal(err)
	}

	if got != 54 {
		t.Fatalf("got %d is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueUint$
func TestValueUint(t *testing.T) {
	got, err := Value("xxx").Uint()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("54").Uint()
	if err != nil {
		t.Fatal(err)
	}

	if got != 54 {
		t.Fatalf("got %d is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueUint8$
func TestValueUint8(t *testing.T) {
	got, err := Value("xxx").Uint8()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("54").Uint8()
	if err != nil {
		t.Fatal(err)
	}

	if got != 54 {
		t.Fatalf("got %d is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueUint16$
func TestValueUint16(t *testing.T) {
	got, err := Value("xxx").Uint16()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("54").Uint16()
	if err != nil {
		t.Fatal(err)
	}

	if got != 54 {
		t.Fatalf("got %d is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueUint32$
func TestValueUint32(t *testing.T) {
	got, err := Value("xxx").Uint32()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("54").Uint32()
	if err != nil {
		t.Fatal(err)
	}

	if got != 54 {
		t.Fatalf("got %d is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueUint64$
func TestValueUint64(t *testing.T) {
	got, err := Value("xxx").Uint64()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("54").Uint64()
	if err != nil {
		t.Fatal(err)
	}

	if got != 54 {
		t.Fatalf("got %d is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueFloat32$
func TestValueFloat32(t *testing.T) {
	got, err := Value("xxx").Float32()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("123.456").Float32()
	if err != nil {
		t.Fatal(err)
	}

	if got != 123.456 {
		t.Fatalf("got %.3f is wrong", got)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestValueFloat64$
func TestValueFloat64(t *testing.T) {
	got, err := Value("xxx").Float64()
	if err == nil {
		t.Fatal("parse error value returns nil")
	}

	got, err = Value("123.456").Float64()
	if err != nil {
		t.Fatal(err)
	}

	if got != 123.456 {
		t.Fatalf("got %.3f is wrong", got)
	}
}
