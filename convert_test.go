// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import "testing"

// go test -v -cover -run=^TestConvertToInts$
func TestConvertToInts(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToInts(strs)
	if err == nil {
		t.Fatalf("convertToInts returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToInts(strs)
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

// go test -v -cover -run=^TestConvertToInt8s$
func TestConvertToInt8s(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToInt8s(strs)
	if err == nil {
		t.Fatalf("convertToInt8s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToInt8s(strs)
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

// go test -v -cover -run=^TestConvertToInt16s$
func TestConvertToInt16s(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToInt16s(strs)
	if err == nil {
		t.Fatalf("convertToInt16s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToInt16s(strs)
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

// go test -v -cover -run=^TestConvertToInt32s$
func TestConvertToInt32s(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToInt32s(strs)
	if err == nil {
		t.Fatalf("convertToInt32s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToInt32s(strs)
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

// go test -v -cover -run=^TestConvertToInt64s$
func TestConvertToInt64s(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToInt64s(strs)
	if err == nil {
		t.Fatalf("convertToInt64s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToInt64s(strs)
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

// go test -v -cover -run=^TestConvertToUints$
func TestConvertToUints(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToUints(strs)
	if err == nil {
		t.Fatalf("convertToUints returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToUints(strs)
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

// go test -v -cover -run=^TestConvertToUint8s$
func TestConvertToUint8s(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToUint8s(strs)
	if err == nil {
		t.Fatalf("convertToUint8s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToUint8s(strs)
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

// go test -v -cover -run=^TestConvertToUint16s$
func TestConvertToUint16s(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToUint16s(strs)
	if err == nil {
		t.Fatalf("convertToUint16s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToUint16s(strs)
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

// go test -v -cover -run=^TestConvertToUint32s$
func TestConvertToUint32s(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToUint32s(strs)
	if err == nil {
		t.Fatalf("convertToUint32s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToUint32s(strs)
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

// go test -v -cover -run=^TestConvertToUint64s$
func TestConvertToUint64s(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToUint64s(strs)
	if err == nil {
		t.Fatalf("convertToUint64s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToUint64s(strs)
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

// go test -v -cover -run=^TestConvertToFloat32s$
func TestConvertToFloat32s(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToFloat32s(strs)
	if err == nil {
		t.Fatalf("convertToFloat32s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToFloat32s(strs)
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

// go test -v -cover -run=^TestConvertToFloat64s$
func TestConvertToFloat64s(t *testing.T) {

	strs := []string{"a", "b", "c"}
	result, err := convertToFloat64s(strs)
	if err == nil {
		t.Fatalf("convertToFloat64s returns wrong err of strs %+v", strs)
	}

	strs = []string{"1", "2", "3"}
	result, err = convertToFloat64s(strs)
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
