// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import (
	"io"
	"os"
)

var (
	// emptyValue is for getting a not existed key.
	emptyValue = (*Value)(nil)
)

// Properties stores all entries in properties.
type Properties struct {

	// keys stores all keys in properties in order.
	keys []string

	// data stores all entries inside.
	data map[string]*Value
}

// NewProperties returns a new properties for use.
func NewProperties() *Properties {
	return &Properties{
		keys: make([]string, 0, 64),
		data: make(map[string]*Value, 64),
	}
}

// Get returns the value of key.
// You can convert value to basic types by using Value.
func (p *Properties) Get(key string) *Value {
	if value, ok := p.data[key]; ok {
		return value
	}
	return emptyValue
}

// Set sets key and value to properties and returns true forever.
// A new value will be added if key doesn't exist.
// The old value will be replaced if key exists.
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

// SetExisted sets key-value only if the key exists.
// If not, it will return false.
func (p *Properties) SetExisted(key string, value string) bool {
	if _, ok := p.data[key]; !ok {
		return false
	}
	return p.Set(key, value)
}

// SetNotExisted sets key-value only if the key doesn't exist.
// If exists, it will return false.
func (p *Properties) SetNotExisted(key string, value string) bool {
	if _, ok := p.data[key]; ok {
		return false
	}
	return p.Set(key, value)
}

// Traverse passes all entries in properties to fn.
// Stop traversing by returning false in fn.
func (p *Properties) Traverse(fn func(key string, value *Value) bool) {

	for _, key := range p.keys {
		if !fn(key, p.Get(key).clone()) {
			break
		}
	}
}

// Store stores p to a string.
func (p *Properties) Store() string {
	return parseFromProperties(p)
}

// StoreTo stores p to io.Writer.
func (p *Properties) StoreTo(writer io.Writer) error {
	_, err := io.WriteString(writer, p.Store())
	return err
}

// StoreToFile stores p to file.
func (p *Properties) StoreToFile(file string) error {

	f, err := os.OpenFile(file, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return p.StoreTo(f)
}
