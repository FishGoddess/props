// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import "strconv"

type Value string

func (v Value) String() string {
	return string(v)
}

func (v Value) Int() (int, error) {
	return strconv.Atoi(v.String())
}

func (v Value) Int8() (int8, error) {
	parsed, err := strconv.ParseInt(v.String(), 10, 8)
	return int8(parsed), err
}

func (v Value) Int16() (int16, error) {
	parsed, err := strconv.ParseInt(v.String(), 10, 16)
	return int16(parsed), err
}

func (v Value) Int32() (int32, error) {
	parsed, err := strconv.ParseInt(v.String(), 10, 32)
	return int32(parsed), err
}

func (v Value) Int64() (int64, error) {
	parsed, err := strconv.ParseInt(v.String(), 10, 64)
	return int64(parsed), err
}

func (v Value) Uint() (uint, error) {
	parsed, err := strconv.ParseUint(v.String(), 10, 64)
	return uint(parsed), err
}

func (v Value) Uint8() (uint8, error) {
	parsed, err := strconv.ParseUint(v.String(), 10, 8)
	return uint8(parsed), err
}

func (v Value) Uint16() (uint16, error) {
	parsed, err := strconv.ParseUint(v.String(), 10, 16)
	return uint16(parsed), err
}

func (v Value) Uint32() (uint32, error) {
	parsed, err := strconv.ParseUint(v.String(), 10, 32)
	return uint32(parsed), err
}

func (v Value) Uint64() (uint64, error) {
	parsed, err := strconv.ParseUint(v.String(), 10, 64)
	return uint64(parsed), err
}

func (v Value) Float32() (float32, error) {
	parsed, err := strconv.ParseFloat(v.String(), 32)
	return float32(parsed), err
}

func (v Value) Float64() (float64, error) {
	parsed, err := strconv.ParseFloat(v.String(), 64)
	return float64(parsed), err
}
