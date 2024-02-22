// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import (
	"fmt"
	"strings"
)

const (
	tokenCR        = "\r"
	tokenLF        = "\n"
	tokenComment   = "#"
	tokenSeparator = "="
)

func parseLine(i int, line string) (key string, value Value, err error) {
	kv := strings.SplitN(line, tokenSeparator, 2)
	if len(kv) < 2 {
		return "", "", fmt.Errorf("props: parse line %d without key or value", i)
	}

	key = strings.TrimSpace(kv[0])
	value = Value(strings.TrimSpace(kv[1]))

	return key, value, nil
}

func parseLines(lines []string) (map[string]Value, error) {
	entries := make(map[string]Value, len(lines))

	for i, line := range lines {
		line = strings.TrimSuffix(line, tokenCR)
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, tokenComment) {
			continue
		}

		key, value, err := parseLine(i+1, line)
		if err != nil {
			return nil, err
		}

		entries[key] = value
	}

	return entries, nil
}

func parse(str string) (map[string]Value, error) {
	lines := strings.Split(str, tokenLF)
	return parseLines(lines)
}
