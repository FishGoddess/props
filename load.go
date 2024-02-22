// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import (
	"io"
	"os"
)

func LoadMap(entries map[string]Value) *Props {
	props := New()
	for key, value := range entries {
		props.Set(key, value)
	}

	return props
}

func LoadReader(reader io.Reader) (*Props, error) {
	bs, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	m, err := parse(string(bs))
	if err != nil {
		return nil, err
	}

	return LoadMap(m), nil
}

func LoadFile(file string) (*Props, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return LoadReader(f)
}
