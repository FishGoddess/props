// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/04 23:00:16

package props

var (
	emptyValue = (*Value)(nil)
)

type Properties struct {
	data map[string]*Value
}

func newProperties() *Properties {
	return &Properties{
		data: make(map[string]*Value, 64),
	}
}

func (p *Properties) Get(key string) *Value {
	if value, ok := p.data[key]; ok {
		return value
	}
	return emptyValue
}

func (p *Properties) Set(key string, value string) bool {

	oldValue, ok := p.data[key]
	if ok {
		p.data[key] = oldValue.set(value)
	} else {
		p.data[key] = newValue(value)
	}
	return true
}

func (p *Properties) SetExisted(key string, value string) bool {
	if _, ok := p.data[key]; !ok {
		return false
	}
	return p.Set(key, value)
}

func (p *Properties) SetNotExisted(key string, value string) bool {
	if _, ok := p.data[key]; ok {
		return false
	}
	return p.Set(key, value)
}
