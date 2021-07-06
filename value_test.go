// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/04 23:00:16

package props

import (
	"math"
	"strconv"
	"testing"
)

// prepareTestProperties prepares one properties for testing.
func prepareTestProperties() *Properties {

	properties := NewProperties()
	properties.Set("key", "value")
	properties.Set("stringKey", "string")
	properties.Set("intKey", "123456789")
	properties.Set("int8Key", strconv.FormatInt(int64(math.MinInt8), 10))
	properties.Set("int16Key", strconv.FormatInt(int64(math.MinInt16), 10))
	properties.Set("int32Key", strconv.FormatInt(int64(math.MinInt32), 10))
	properties.Set("int64Key", strconv.FormatInt(int64(math.MinInt64), 10))
	properties.Set("uintKey", strconv.FormatUint(uint64(999), 10))
	properties.Set("uint8Key", strconv.FormatUint(uint64(math.MaxUint8), 10))
	properties.Set("uint16Key", strconv.FormatUint(uint64(math.MaxUint16), 10))
	properties.Set("uint32Key", strconv.FormatUint(uint64(math.MaxUint32), 10))
	properties.Set("uint64Key", strconv.FormatUint(uint64(math.MaxUint64), 10))
	properties.Set("float32Key", "32.32")
	properties.Set("float64Key", "64.64")
	properties.Set("stringsKey", "string1,string2,string3")
	properties.Set("intsKey", "123,-456,789")
	properties.Set("int8sKey", "1,-2,3")
	properties.Set("int16sKey", "443,-3306,6379")
	properties.Set("int32sKey", "123456,-234567,345678")
	properties.Set("int64sKey", "123456789,-234567890,345678901")
	properties.Set("uintsKey", "123,456,789")
	properties.Set("uint8sKey", "1,2,3")
	properties.Set("uint16sKey", "443,3306,6379")
	properties.Set("uint32sKey", "123456,234567,345678")
	properties.Set("uint64sKey", "123456789,234567890,345678901")
	properties.Set("float32sKey", "1.23,4.56,7.89")
	properties.Set("float64sKey", "123.456789,234.567891,567.890123")
	return properties
}

// go test -v -cover -run=^TestValueString$
func TestValueString(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").String("defaultValue"); value != "defaultValue" {
		t.Fatalf("get a not existed key is %s", value)
	}

	if value := properties.Get("stringKey").String("defaultValue"); value != "string" {
		t.Fatalf("value of key (%s) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueInt$
func TestValueInt(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Int(1); value != 1 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("intKey").Int(1)
	if value != 123456789 {
		t.Fatalf("value of intKey (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueInt8$
func TestValueInt8(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Int8(math.MinInt8); value != math.MinInt8 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("int8Key").Int8(1)
	if value != math.MinInt8 {
		t.Fatalf("value of int8Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueInt16$
func TestValueInt16(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Int16(math.MinInt16); value != math.MinInt16 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("int16Key").Int16(1)
	if value != math.MinInt16 {
		t.Fatalf("value of int16Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueInt32$
func TestValueInt32(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Int32(math.MinInt32); value != math.MinInt32 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("int32Key").Int32(1)
	if value != math.MinInt32 {
		t.Fatalf("value of int32Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueInt64$
func TestValueInt64(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Int64(math.MinInt64); value != math.MinInt64 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("int64Key").Int64(1)
	if value != math.MinInt64 {
		t.Fatalf("value of int64Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueUint$
func TestValueUint(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Uint(999); value != 999 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("uintKey").Uint(1)
	if value != 999 {
		t.Fatalf("value of uintKey (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueUint8$
func TestValueUint8(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Uint8(math.MaxUint8); value != math.MaxUint8 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("uint8Key").Uint8(1)
	if value != math.MaxUint8 {
		t.Fatalf("value of uint8Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueUint16$
func TestValueUint16(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Uint16(math.MaxUint16); value != math.MaxUint16 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("uint16Key").Uint16(1)
	if value != math.MaxUint16 {
		t.Fatalf("value of uint16Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueUint32$
func TestValueUint32(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Uint32(math.MaxUint32); value != math.MaxUint32 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("uint32Key").Uint32(1)
	if value != math.MaxUint32 {
		t.Fatalf("value of uint32Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueUint64$
func TestValueUint64(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Uint64(math.MaxUint64); value != math.MaxUint64 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("uint64Key").Uint64(1)
	if value != math.MaxUint64 {
		t.Fatalf("value of uint64Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueFloat32$
func TestValueFloat32(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Float32(math.MaxFloat32); value != math.MaxFloat32 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("float32Key").Float32(1)
	if value != 32.32 {
		t.Fatalf("value of float32Key (%.3f) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueFloat64$
func TestValueFloat64(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Float64(math.MaxFloat64); value != math.MaxFloat64 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("float64Key").Float64(1)
	if value != 64.64 {
		t.Fatalf("value of float64Key (%.3f) is wrong", value)
	}
}

// go test -v -cover -run=^TestValueStrings$
func TestValueStrings(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Strings(",", []string{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("stringsKey").Strings(",", []string{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []string{"string1", "string2", "string3"}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%s) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueInts$
func TestValueInts(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Ints(",", []int{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("intsKey").Ints(",", []int{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []int{123, -456, 789}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%d) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueInt8s$
func TestValueInt8s(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Int8s(",", []int8{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("int8sKey").Int8s(",", []int8{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []int8{1, -2, 3}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%d) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueInt16s$
func TestValueInt16s(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Int16s(",", []int16{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("int16sKey").Int16s(",", []int16{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []int16{443, -3306, 6379}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%d) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueInt32s$
func TestValueInt32s(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Int32s(",", []int32{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("int32sKey").Int32s(",", []int32{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []int32{123456, -234567, 345678}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%d) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueInt64s$
func TestValueInt64s(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Int64s(",", []int64{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("int64sKey").Int64s(",", []int64{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []int64{123456789, -234567890, 345678901}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%d) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueUints$
func TestValueUints(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Uints(",", []uint{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("uintsKey").Uints(",", []uint{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []uint{123, 456, 789}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%d) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueUint8s$
func TestValueUint8s(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Uint8s(",", []uint8{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("uint8sKey").Uint8s(",", []uint8{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []uint8{1, 2, 3}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%d) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueUint16s$
func TestValueUint16s(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Uint16s(",", []uint16{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("uint16sKey").Uint16s(",", []uint16{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []uint16{443, 3306, 6379}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%d) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueUint32s$
func TestValueUint32s(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Uint32s(",", []uint32{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("uint32sKey").Uint32s(",", []uint32{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []uint32{123456, 234567, 345678}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%d) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueUint64s$
func TestValueUint64s(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Uint64s(",", []uint64{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("uint64sKey").Uint64s(",", []uint64{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []uint64{123456789, 234567890, 345678901}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%d) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueFloat32s$
func TestValueFloat32s(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Float32s(",", []float32{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("float32sKey").Float32s(",", []float32{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []float32{1.23, 4.56, 7.89}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%f) is wrong", i, value)
		}
	}
}

// go test -v -cover -run=^TestValueFloat64s$
func TestValueFloat64s(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.Get("notExistedKey").Float64s(",", []float64{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.Get("float64sKey").Float64s(",", []float64{})
	if len(value) != 3 {
		t.Fatalf("len of value (%d) is wrong", len(value))
	}

	result := []float64{123.456789, 234.567891, 567.890123}
	for i := 0; i < 3; i++ {
		if value[i] != result[i] {
			t.Fatalf("value of key %d (%f) is wrong", i, value)
		}
	}
}
