// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import (
	"strconv"
	"strings"
)

// Value stores the data of key in properties.
type Value struct {
	// data stores the real value inside.
	data string
}

// newValue returns a new value storing data.
func newValue(data string) *Value {
	return &Value{data: data}
}

// get returns the real value inside.
func (v *Value) get() string {
	if v != nil {
		return v.data
	}
	return ""
}

// set set data to Value and returns itself.
func (v *Value) set(data string) *Value {
	if v != nil {
		v.data = data
	}
	return v
}

// clone return a copy of v.
func (v *Value) clone() *Value {
	if v == nil {
		return nil
	}
	return newValue(v.data)
}

// Absent returns if v doesn't have data or not.
func (v *Value) Absent() bool {
	if v == nil {
		return true
	}
	return false
}

// String returns data as string type and defaultValue will be returned if data is absent.
func (v *Value) String(defaultValue string) string {
	if v.Absent() {
		return defaultValue
	}
	return v.data
}

// Int returns data as int type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Int(defaultValue int) int {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.Atoi(v.data)
	if err != nil {
		return defaultValue
	}
	return result
}

// Int8 returns data as int8 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Int8(defaultValue int8) int8 {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseInt(v.data, 10, 8)
	if err != nil {
		return defaultValue
	}
	return int8(result)
}

// Int16 returns data as int16 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Int16(defaultValue int16) int16 {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseInt(v.data, 10, 16)
	if err != nil {
		return defaultValue
	}
	return int16(result)
}

// Int32 returns data as int32 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Int32(defaultValue int32) int32 {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseInt(v.data, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int32(result)
}

// Int64 returns data as int64 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Int64(defaultValue int64) int64 {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseInt(v.data, 10, 64)
	if err != nil {
		return defaultValue
	}
	return result
}

// Uint returns data as uint type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Uint(defaultValue uint) uint {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseUint(v.data, 10, 64)
	if err != nil {
		return defaultValue
	}
	return uint(result)
}

// Uint8 returns data as uint8 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Uint8(defaultValue uint8) uint8 {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseUint(v.data, 10, 8)
	if err != nil {
		return defaultValue
	}
	return uint8(result)
}

// Uint16 returns data as uint16 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Uint16(defaultValue uint16) uint16 {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseUint(v.data, 10, 16)
	if err != nil {
		return defaultValue
	}
	return uint16(result)
}

// Uint32 returns data as uint32 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Uint32(defaultValue uint32) uint32 {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseUint(v.data, 10, 32)
	if err != nil {
		return defaultValue
	}
	return uint32(result)
}

// Uint64 returns data as uint64 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Uint64(defaultValue uint64) uint64 {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseUint(v.data, 10, 64)
	if err != nil {
		return defaultValue
	}
	return result
}

// Float32 returns data as float32 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Float32(defaultValue float32) float32 {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseFloat(v.data, 32)
	if err != nil {
		return defaultValue
	}
	return float32(result)
}

// Float64 returns data as float64 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Float64(defaultValue float64) float64 {

	if v.Absent() {
		return defaultValue
	}

	result, err := strconv.ParseFloat(v.data, 64)
	if err != nil {
		return defaultValue
	}
	return result
}

// Strings returns data as []string type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Strings(separator string, defaultValue []string) []string {
	if v.Absent() {
		return defaultValue
	}
	return strings.Split(v.data, separator)
}

// Ints returns data as []int type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Ints(separator string, defaultValue []int) []int {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToInts(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Int8s returns data as []int8 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Int8s(separator string, defaultValue []int8) []int8 {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToInt8s(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Int16s returns data as []int16 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Int16s(separator string, defaultValue []int16) []int16 {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToInt16s(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Int32s returns data as []int32 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Int32s(separator string, defaultValue []int32) []int32 {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToInt32s(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Int64s returns data as []int64 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Int64s(separator string, defaultValue []int64) []int64 {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToInt64s(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Uints returns data as []uint type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Uints(separator string, defaultValue []uint) []uint {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToUints(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Uint8s returns data as []uint8 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Uint8s(separator string, defaultValue []uint8) []uint8 {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToUint8s(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Uint16s returns data as []uint16 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Uint16s(separator string, defaultValue []uint16) []uint16 {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToUint16s(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Uint32s returns data as []uint32 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Uint32s(separator string, defaultValue []uint32) []uint32 {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToUint32s(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Uint64s returns data as []uint64 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Uint64s(separator string, defaultValue []uint64) []uint64 {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToUint64s(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Float32s returns data as []float32 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Float32s(separator string, defaultValue []float32) []float32 {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToFloat32s(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}

// Float64s returns data as []int64 type and defaultValue will be returned if data is absent or something wrong happens.
func (v *Value) Float64s(separator string, defaultValue []float64) []float64 {

	if v.Absent() {
		return defaultValue
	}

	result, err := convertToFloat64s(strings.Split(v.data, separator))
	if err != nil {
		return defaultValue
	}
	return result
}
