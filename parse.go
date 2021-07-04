// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/04 23:00:16

package props

import (
	"fmt"
	"strings"
)

const (
	cr = "\r"
	lf = "\n"

	commentPrefix     = "#"
	keyValueSeparator = "="
)

func isBlank(line string) bool {
	return line == ""
}

func isComment(line string) bool {
	return strings.HasPrefix(line, commentPrefix)
}

func kvOf(line string) (string, string, error) {

	kv := strings.SplitN(line, keyValueSeparator, 2)
	if len(kv) < 2 {
		return "", "", fmt.Errorf("line {%s} doesn't have key and value", line)
	}
	return strings.TrimSpace(kv[0]), strings.TrimSpace(strings.Trim(kv[1], cr)), nil
}

func parseFromString(str string) (*Properties, error) {

	properties := NewProperties()

	lines := strings.Split(str, lf)
	for _, line := range lines {

		line = strings.TrimSpace(line)
		if isBlank(line) {
			continue
		}

		if isComment(line) {
			continue
		}

		key, value, err := kvOf(line)
		if err != nil {
			return nil, err
		}
		properties.Set(key, value)
	}
	return properties, nil
}

func parseFromProperties(properties *Properties) string {

	buffer := make([]byte, 0, 512)
	properties.Traverse(func(key string, value *Value) bool {
		buffer = append(buffer, key...)
		buffer = append(buffer, keyValueSeparator...)
		buffer = append(buffer, value.get()...)
		buffer = append(buffer, lf...)
		return true
	})
	return string(buffer)
}
