// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import "sync"

const (
	defaultMapCapacity = 16
)

type Props struct {
	entries map[string]Value

	lock sync.RWMutex
}

func New() *Props {
	props := &Props{
		entries: make(map[string]Value, defaultMapCapacity),
	}

	return props
}

func (p *Props) Get(key string) Value {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.entries[key]
}

func (p *Props) Set(key string, value Value) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.entries[key] = value
}

func (p *Props) Delete(key string) {
	p.lock.Lock()
	defer p.lock.Unlock()

	delete(p.entries, key)
}

func (p *Props) List(f func(key string, value Value) (finished bool)) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	for key, value := range p.entries {
		if finished := f(key, value); finished {
			break
		}
	}
}

func (p *Props) Count() int {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return len(p.entries)
}
