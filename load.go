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
	"io/ioutil"
)

func Load(data string) (*Properties, error) {
	return parseFromString(data)
}

func LoadFrom(reader io.Reader) (*Properties, error) {

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return Load(string(data))
}

func LoadFromFile(file string) (*Properties, error) {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return Load(string(data))
}
