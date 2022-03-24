// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package convert

import "strconv"

type Strings []string

// Ints converts []string to []int and returns an error if failed.
func (ss Strings) Ints() ([]int, error) {
	result := make([]int, len(ss))
	for i, str := range ss {
		one, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		result[i] = one
	}
	return result, nil
}

// Int8s converts []string to []int8 and returns an error if failed.
func (ss Strings) Int8s() ([]int8, error) {
	result := make([]int8, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseInt(str, 10, 8)
		if err != nil {
			return nil, err
		}

		result[i] = int8(one)
	}
	return result, nil
}

// Int16s converts []string to []int16 and returns an error if failed.
func (ss Strings) Int16s() ([]int16, error) {
	result := make([]int16, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseInt(str, 10, 16)
		if err != nil {
			return nil, err
		}

		result[i] = int16(one)
	}
	return result, nil
}

// Int32s converts []string to []int32 and returns an error if failed.
func (ss Strings) Int32s() ([]int32, error) {
	result := make([]int32, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return nil, err
		}

		result[i] = int32(one)
	}
	return result, nil
}

// Int64s converts []string to []int64 and returns an error if failed.
func (ss Strings) Int64s() ([]int64, error) {
	result := make([]int64, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}

		result[i] = one
	}
	return result, nil
}

// Uints converts []string to []uint and returns an error if failed.
func (ss Strings) Uints() ([]uint, error) {
	result := make([]uint, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return nil, err
		}

		result[i] = uint(one)
	}
	return result, nil
}

// Uint8s converts []string to []uint8 and returns an error if failed.
func (ss Strings) Uint8s() ([]uint8, error) {
	result := make([]uint8, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseUint(str, 10, 8)
		if err != nil {
			return nil, err
		}

		result[i] = uint8(one)
	}
	return result, nil
}

// Uint16s converts []string to []uint16 and returns an error if failed.
func (ss Strings) Uint16s() ([]uint16, error) {
	result := make([]uint16, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseUint(str, 10, 16)
		if err != nil {
			return nil, err
		}

		result[i] = uint16(one)
	}
	return result, nil
}

// Uint32s converts []string to []uint32 and returns an error if failed.
func (ss Strings) Uint32s() ([]uint32, error) {
	result := make([]uint32, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return nil, err
		}

		result[i] = uint32(one)
	}
	return result, nil
}

// Uint64s converts []string to []uint64 and returns an error if failed.
func (ss Strings) Uint64s() ([]uint64, error) {
	result := make([]uint64, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return nil, err
		}

		result[i] = one
	}
	return result, nil
}

// Float32s converts []string to []float32 and returns an error if failed.
func (ss Strings) Float32s() ([]float32, error) {
	result := make([]float32, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return nil, err
		}

		result[i] = float32(one)
	}
	return result, nil
}

// Float64s converts []string to []float64 and returns an error if failed.
func (ss Strings) Float64s() ([]float64, error) {
	result := make([]float64, len(ss))
	for i, str := range ss {
		one, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}

		result[i] = one
	}
	return result, nil
}
