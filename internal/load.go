// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package internal

import (
	"io"
	"io/ioutil"
	"os"
)

// Load loads props from io.Reader and maps it to v.
func Load(reader io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	parsed, err := decodeString(string(data))
	if err != nil {
		return err
	}

	return mapToStruct(parsed, v)
}

// LoadFile loads props from file and maps it to v.
func LoadFile(source string, v interface{}) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	return Load(file, v)
}
