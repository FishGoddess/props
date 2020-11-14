// Author: fish
// Email: yezijie@bilibili.com
// Created at 2020/11/13 14:06

package props

import (
	"math"
	"strconv"
	"testing"
)

// prepareTestProperties prepares one properties for testing.
func prepareTestProperties() *Properties {

	properties := newProperties()
	data := properties.data
	data["key"] = "value"
	data["stringKey"] = "string"
	data["stringsKey"] = "string1,string2,string3"
	data["intKey"] = "123456789"
	data["int8Key"] = strconv.FormatInt(int64(math.MinInt8), 10)
	data["int16Key"] = strconv.FormatInt(int64(math.MinInt16), 10)
	data["int32Key"] = strconv.FormatInt(int64(math.MinInt32), 10)
	data["int64Key"] = strconv.FormatInt(int64(math.MinInt64), 10)
	data["uint8Key"] = strconv.FormatUint(uint64(math.MaxUint8), 10)
	data["uint16Key"] = strconv.FormatUint(uint64(math.MaxUint16), 10)
	data["uint32Key"] = strconv.FormatUint(uint64(math.MaxUint32), 10)
	data["uint64Key"] = strconv.FormatUint(uint64(math.MaxUint64), 10)
	data["float32Key"] = "32.32"
	data["float64Key"] = "64.64"
	return properties
}

// go test -v -cover -run=^TestPropertiesGet$
func TestPropertiesGet(t *testing.T) {

	properties := prepareTestProperties()
	if value, ok := properties.Get("notExistedKey"); ok {
		t.Fatalf("get a not existed key is %s", value)
	}

	if value, ok := properties.Get("key"); !ok || value != "value" {
		t.Fatalf("value of key (%s) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetString$
func TestPropertiesGetString(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetString("notExistedKey", "defaultValue"); value != "defaultValue" {
		t.Fatalf("get a not existed key is %s", value)
	}

	if value := properties.GetString("stringKey", "defaultValue"); value != "string" {
		t.Fatalf("value of key (%s) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetStrings$
func TestPropertiesGetStrings(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetStrings("notExistedKey", ",", []string{}); len(value) != 0 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetStrings("stringsKey", ",", []string{})
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

// go test -v -cover -run=^TestPropertiesGetInt$
func TestPropertiesGetInt(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetInt("notExistedKey", 1); value != 1 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetInt("intKey", 1)
	if value != 123456789 {
		t.Fatalf("value of intKey (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetInt8$
func TestPropertiesGetInt8(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetInt8("notExistedKey", math.MinInt8); value != math.MinInt8 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetInt8("int8Key", 1)
	if value != math.MinInt8 {
		t.Fatalf("value of int8Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetInt16$
func TestPropertiesGetInt16(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetInt16("notExistedKey", math.MinInt16); value != math.MinInt16 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetInt16("int16Key", 1)
	if value != math.MinInt16 {
		t.Fatalf("value of int16Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetInt32$
func TestPropertiesGetInt32(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetInt32("notExistedKey", math.MinInt32); value != math.MinInt32 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetInt32("int32Key", 1)
	if value != math.MinInt32 {
		t.Fatalf("value of int32Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetInt64$
func TestPropertiesGetInt64(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetInt64("notExistedKey", math.MinInt64); value != math.MinInt64 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetInt64("int64Key", 1)
	if value != math.MinInt64 {
		t.Fatalf("value of int64Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetUint8$
func TestPropertiesGetUint8(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetUint8("notExistedKey", math.MaxUint8); value != math.MaxUint8 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetUint8("uint8Key", 1)
	if value != math.MaxUint8 {
		t.Fatalf("value of uint8Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetUint16$
func TestPropertiesGetUint16(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetUint16("notExistedKey", math.MaxUint16); value != math.MaxUint16 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetUint16("uint16Key", 1)
	if value != math.MaxUint16 {
		t.Fatalf("value of uint16Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetUint32$
func TestPropertiesGetUint32(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetUint32("notExistedKey", math.MaxUint32); value != math.MaxUint32 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetUint32("uint32Key", 1)
	if value != math.MaxUint32 {
		t.Fatalf("value of uint32Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetUint64$
func TestPropertiesGetUint64(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetUint64("notExistedKey", math.MaxUint64); value != math.MaxUint64 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetUint64("uint64Key", 1)
	if value != math.MaxUint64 {
		t.Fatalf("value of uint64Key (%d) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetFloat32$
func TestPropertiesGetFloat32(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetFloat32("notExistedKey", math.MaxFloat32); value != math.MaxFloat32 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetFloat32("float32Key", 1)
	if value != 32.32 {
		t.Fatalf("value of float32Key (%.3f) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesGetFloat64$
func TestPropertiesGetFloat64(t *testing.T) {

	properties := prepareTestProperties()
	if value := properties.GetFloat64("notExistedKey", math.MaxFloat64); value != math.MaxFloat64 {
		t.Fatalf("get a not existed key is %v", value)
	}

	value := properties.GetFloat64("float64Key", 1)
	if value != 64.64 {
		t.Fatalf("value of float64Key (%.3f) is wrong", value)
	}
}
