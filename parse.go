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
	cr = "\r" // Carriage return
	lf = "\n" // Line break

	commentPrefix     = "#" // The prefix of comment line
	keyValueSeparator = "=" // The separator between key and value
)

// isBlank returns line is blank ling or not.
func isBlank(line string) bool {
	return line == ""
}

// isComment returns line is comment line or not.
func isComment(line string) bool {
	return strings.HasPrefix(line, commentPrefix)
}

// needToBeIgnored returns line needs to be ignored or not.
func needToBeIgnored(line string) bool {
	return isBlank(line) || isComment(line)
}

// kvOf parses key and value from line.
// It returns an error if failed.
func kvOf(line string) (string, string, error) {

	kv := strings.SplitN(line, keyValueSeparator, 2)
	if len(kv) < 2 {
		return "", "", fmt.Errorf("line {%s} doesn't have key and value", line)
	}
	return strings.TrimSpace(kv[0]), strings.TrimSpace(strings.Trim(kv[1], cr)), nil
}

// parseFromString parses properties from string and returns an error if failed.
func parseFromString(str string) (*Properties, error) {

	properties := NewProperties()

	lines := strings.Split(str, lf)
	for _, line := range lines {

		line = strings.TrimSpace(line)
		if needToBeIgnored(line) {
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

// parseFromProperties parses string from properties.
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
