// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/07 00:35:59

package props

import "strconv"

// convertToInts converts []string to []int and returns an error if failed.
func convertToInts(strs []string) ([]int, error) {

	result := make([]int, len(strs))
	for i, str := range strs {
		one, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		result[i] = one
	}
	return result, nil
}

// convertToInt8s converts []string to []int8 and returns an error if failed.
func convertToInt8s(strs []string) ([]int8, error) {

	result := make([]int8, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseInt(str, 10, 8)
		if err != nil {
			return nil, err
		}
		result[i] = int8(one)
	}
	return result, nil
}

// convertToInt16s converts []string to []int16 and returns an error if failed.
func convertToInt16s(strs []string) ([]int16, error) {

	result := make([]int16, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseInt(str, 10, 16)
		if err != nil {
			return nil, err
		}
		result[i] = int16(one)
	}
	return result, nil
}

// convertToInt32s converts []string to []int32 and returns an error if failed.
func convertToInt32s(strs []string) ([]int32, error) {

	result := make([]int32, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return nil, err
		}
		result[i] = int32(one)
	}
	return result, nil
}

// convertToInt64s converts []string to []int64 and returns an error if failed.
func convertToInt64s(strs []string) ([]int64, error) {

	result := make([]int64, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		result[i] = int64(one)
	}
	return result, nil
}

// convertToUints converts []string to []uint and returns an error if failed.
func convertToUints(strs []string) ([]uint, error) {

	result := make([]uint, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return nil, err
		}
		result[i] = uint(one)
	}
	return result, nil
}

// convertToUint8s converts []string to []uint8 and returns an error if failed.
func convertToUint8s(strs []string) ([]uint8, error) {

	result := make([]uint8, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseUint(str, 10, 8)
		if err != nil {
			return nil, err
		}
		result[i] = uint8(one)
	}
	return result, nil
}

// convertToUint16s converts []string to []uint16 and returns an error if failed.
func convertToUint16s(strs []string) ([]uint16, error) {

	result := make([]uint16, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseUint(str, 10, 16)
		if err != nil {
			return nil, err
		}
		result[i] = uint16(one)
	}
	return result, nil
}

// convertToUint32s converts []string to []uint32 and returns an error if failed.
func convertToUint32s(strs []string) ([]uint32, error) {

	result := make([]uint32, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return nil, err
		}
		result[i] = uint32(one)
	}
	return result, nil
}

// convertToUint64s converts []string to []uint64 and returns an error if failed.
func convertToUint64s(strs []string) ([]uint64, error) {

	result := make([]uint64, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return nil, err
		}
		result[i] = one
	}
	return result, nil
}

// convertToFloat32s converts []string to []float32 and returns an error if failed.
func convertToFloat32s(strs []string) ([]float32, error) {

	result := make([]float32, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return nil, err
		}
		result[i] = float32(one)
	}
	return result, nil
}

// convertToFloat64s converts []string to []float64 and returns an error if failed.
func convertToFloat64s(strs []string) ([]float64, error) {

	result := make([]float64, len(strs))
	for i, str := range strs {
		one, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		result[i] = one
	}
	return result, nil
}
