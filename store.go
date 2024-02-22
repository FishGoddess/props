// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import (
	"io"
	"io/fs"
	"os"
)

var (
	FilePerm fs.FileMode = 0644
)

func StoreMap(props *Props) map[string]Value {
	m := make(map[string]Value, props.Count())
	props.List(func(key string, value Value) (finished bool) {
		m[key] = value
		return false
	})

	return m
}

func StoreWriter(props *Props, writer io.Writer) error {
	var err error

	line := func(key string, value Value) string {
		return key + tokenSeparator + value.String() + tokenLF
	}

	props.List(func(key string, value Value) (finished bool) {
		bs := []byte(line(key, value))
		if _, err = writer.Write(bs); err != nil {
			return true
		}

		return false
	})

	return err
}

func StoreFile(props *Props, file string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, FilePerm)
	if err != nil {
		return err
	}

	defer f.Close()
	return StoreWriter(props, f)
}
