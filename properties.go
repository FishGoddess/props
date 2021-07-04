// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/04 23:00:16

package props

import (
	"io"
	"os"
)

var (
	emptyValue = (*Value)(nil)
)

type Properties struct {
	keys []string
	data map[string]*Value
}

func NewProperties() *Properties {
	return &Properties{
		keys: make([]string, 0, 64),
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
		p.keys = append(p.keys, key)
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

func (p *Properties) Traverse(fn func(key string, value *Value) bool) {

	for _, key := range p.keys {
		if !fn(key, p.Get(key).clone()) {
			break
		}
	}
}

func (p *Properties) Store() string {
	return parseFromProperties(p)
}

func (p *Properties) StoreTo(writer io.Writer) error {
	_, err := io.WriteString(writer, p.Store())
	return err
}

func (p *Properties) StoreToFile(file string) error {

	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return p.StoreTo(f)
}
