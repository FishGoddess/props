// Author: fish
// Email: yezijie@bilibili.com
// Created at 2020/11/12 17:02

package props

import (
	"strconv"
	"strings"
)

type Properties struct {
	data map[string]string
}

func newProperties() *Properties {
	return &Properties{
		data: map[string]string{},
	}
}

func (p *Properties) Get(key string) (string, bool) {
	value, ok := p.data[key]
	return value, ok
}

func (p *Properties) GetString(key string, defaultValue string) string {
	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}
	return value
}

func (p *Properties) GetStrings(key string, separator string, defaultValue []string) []string {
	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}
	return strings.Split(value, separator)
}

func (p *Properties) GetInt(key string, defaultValue int) int {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}

func (p *Properties) GetInt8(key string, defaultValue int8) int8 {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return defaultValue
	}
	return int8(result)
}

func (p *Properties) GetInt16(key string, defaultValue int16) int16 {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return defaultValue
	}
	return int16(result)
}

func (p *Properties) GetInt32(key string, defaultValue int32) int32 {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int32(result)
}

func (p *Properties) GetInt64(key string, defaultValue int64) int64 {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}
	return result
}

func (p *Properties) GetUint8(key string, defaultValue uint8) uint8 {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return defaultValue
	}
	return uint8(result)
}

func (p *Properties) GetUint16(key string, defaultValue uint16) uint16 {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return defaultValue
	}
	return uint16(result)
}

func (p *Properties) GetUint32(key string, defaultValue uint32) uint32 {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return defaultValue
	}
	return uint32(result)
}

func (p *Properties) GetUint64(key string, defaultValue uint64) uint64 {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return defaultValue
	}
	return result
}

func (p *Properties) GetFloat32(key string, defaultValue float32) float32 {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defaultValue
	}
	return float32(result)
}

func (p *Properties) GetFloat64(key string, defaultValue float64) float64 {

	value, ok := p.Get(key)
	if !ok {
		return defaultValue
	}

	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return result
}
