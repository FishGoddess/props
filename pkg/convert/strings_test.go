// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package convert

import "testing"

// go test -v -cover -run=^TestStringsInts$
func TestStringsInts(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Ints()
	if err == nil {
		t.Fatalf("convertToInts returns wrong err of strs %+v", strs)
	}

	strs = Strings([]string{"1", "2", "3"})
	result, err = strs.Ints()
	if err != nil {
		t.Fatalf("convertToInts returns wrong err of strs %+v", strs)
	}

	rights := []int{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %d != %d", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsInt8s$
func TestStringsInt8s(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Int8s()
	if err == nil {
		t.Fatalf("convertToInt8s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Int8s()
	if err != nil {
		t.Fatalf("convertToInt8s returns wrong err of strs %+v", strs)
	}

	rights := []int8{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %d != %d", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsInt16s$
func TestStringsInt16s(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Int16s()
	if err == nil {
		t.Fatalf("convertToInt16s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Int16s()
	if err != nil {
		t.Fatalf("convertToInt16s returns wrong err of strs %+v", strs)
	}

	rights := []int16{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %d != %d", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsInt32s$
func TestStringsInt32s(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Int32s()
	if err == nil {
		t.Fatalf("convertToInt32s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Int32s()
	if err != nil {
		t.Fatalf("convertToInt32s returns wrong err of strs %+v", strs)
	}

	rights := []int32{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %d != %d", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsInt64s$
func TestStringsInt64s(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Int64s()
	if err == nil {
		t.Fatalf("convertToInt64s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Int64s()
	if err != nil {
		t.Fatalf("convertToInt64s returns wrong err of strs %+v", strs)
	}

	rights := []int64{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %d != %d", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsUints$
func TestStringsUints(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Uints()
	if err == nil {
		t.Fatalf("convertToUints returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Uints()
	if err != nil {
		t.Fatalf("convertToUints returns wrong err of strs %+v", strs)
	}

	rights := []uint{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %d != %d", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsUint8s$
func TestStringsUint8s(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Uint8s()
	if err == nil {
		t.Fatalf("convertToUint8s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Uint8s()
	if err != nil {
		t.Fatalf("convertToUint8s returns wrong err of strs %+v", strs)
	}

	rights := []uint8{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %d != %d", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsUint16s$
func TestStringsUint16s(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Uint16s()
	if err == nil {
		t.Fatalf("convertToUint16s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Uint16s()
	if err != nil {
		t.Fatalf("convertToUint16s returns wrong err of strs %+v", strs)
	}

	rights := []uint16{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %d != %d", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsUint32s$
func TestStringsUint32s(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Uint32s()
	if err == nil {
		t.Fatalf("convertToUint32s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Uint32s()
	if err != nil {
		t.Fatalf("convertToUint32s returns wrong err of strs %+v", strs)
	}

	rights := []uint32{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %d != %d", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsUint64s$
func TestStringsUint64s(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Uint64s()
	if err == nil {
		t.Fatalf("convertToUint64s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Uint64s()
	if err != nil {
		t.Fatalf("convertToUint64s returns wrong err of strs %+v", strs)
	}

	rights := []uint64{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %d != %d", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsFloat32s$
func TestStringsFloat32s(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Float32s()
	if err == nil {
		t.Fatalf("convertToFloat32s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Float32s()
	if err != nil {
		t.Fatalf("convertToFloat32s returns wrong err of strs %+v", strs)
	}

	rights := []float32{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %f != %f", i, i, result[i], rights[i])
		}
	}
}

// go test -v -cover -run=^TestStringsFloat64s$
func TestStringsFloat64s(t *testing.T) {
	strs := Strings([]string{"a", "b", "c"})
	result, err := strs.Float64s()
	if err == nil {
		t.Fatalf("convertToFloat64s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = strs.Float64s()
	if err != nil {
		t.Fatalf("convertToFloat64s returns wrong err of strs %+v", strs)
	}

	rights := []float64{1, 2, 3}
	for i := 0; i < len(strs); i++ {
		if result[i] != rights[i] {
			t.Fatalf("result[%d] != rights[%d], %f != %f", i, i, result[i], rights[i])
		}
	}
}
