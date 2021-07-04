// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/04 23:00:16

package props

import (
	"strconv"
	"strings"
)

type Value struct {
	data string
}

func newValue(data string) *Value {
	return &Value{data: data}
}

func (v *Value) get() string {
	if v != nil {
		return v.data
	}
	return ""
}

func (v *Value) set(data string) *Value {
	if v != nil {
		v.data = data
	}
	return v
}

func (v *Value) clone() *Value {
	if v == nil {
		return nil
	}
	return newValue(v.data)
}

func (v *Value) Absent() bool {
	if v == nil {
		return true
	}
	return false
}

func (v *Value) String(defaultValue string) string {
	if v.Absent() {
		return defaultValue
	}
	return v.data
}

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

func (v *Value) Strings(separator string, defaultValue []string) []string {
	if v.Absent() {
		return defaultValue
	}
	return strings.Split(v.data, separator)
}
