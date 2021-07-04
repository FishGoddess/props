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

func parseFrom(str string) (*Properties, error) {

	properties := newProperties()

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
