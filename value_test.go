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
	properties.Set("stringsKey", "string1,string2,string3")
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
