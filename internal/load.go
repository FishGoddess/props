// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// parseLine parses keys and value from line and returns an error if failed.
func parseLine(i int, line string) (string, string, error) {
	kv := strings.SplitN(line, tokenSeparator, 2)
	if len(kv) < 2 {
		return "", "", fmt.Errorf("line %d {%s} doesn't have keys or value", i, line)
	}

	key := strings.TrimSpace(kv[0])
	value := strings.TrimSpace(strings.TrimSuffix(kv[1], tokenCR)) // Trim cr / space
	return key, value, nil
}

func parseLines(lines []string) (map[string][]string, error) {
	result := make(map[string][]string, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if isBlank(line) || isComment(line) {
			continue
		}

		key, value, err := parseLine(i, line)
		if err != nil {
			return nil, err
		}

		result[key] = append(result[key], value)
	}

	return result, nil
}

// LoadFrom loads props from io.Reader and returns an error if failed.
func LoadFrom(reader io.Reader) (interface{}, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return parseLines(strings.Split(string(data), tokenLF))
}

// LoadFromFile loads props from file and returns an error if failed.
func LoadFromFile(source string) (interface{}, error) {
	file, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return LoadFrom(file)
}
