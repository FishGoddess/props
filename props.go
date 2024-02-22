// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

const (
	defaultMapCapacity = 16
)

type Props struct {
	entries map[string]Value
}

func New() *Props {
	props := &Props{
		entries: make(map[string]Value, defaultMapCapacity),
	}

	return props
}

func NewFromMap(entries map[string]Value) *Props {
	props := New()
	for key, value := range entries {
		props.Set(key, value)
	}

	return props
}

func (p *Props) Get(key string) Value {
	return p.entries[key]
}

func (p *Props) Set(key string, value Value) {
	p.entries[key] = value
}
