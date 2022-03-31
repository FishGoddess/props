// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"strings"
)

const (
	tokenCR           = "\r" // Carriage return
	tokenLF           = "\n" // Line break
	tokenComment      = "#"  // Prefix of comment line
	tokenSeparator    = "="  // Separator between keys and value
	tokenKeySeparator = "."  // Separator between keys and value
)

// isBlank returns if line is blank or not.
func isBlank(line string) bool {
	return line == ""
}

// isComment returns if line is comment or not.
func isComment(line string) bool {
	return strings.HasPrefix(line, tokenComment)
}

// newKeyString returns a new key constructed from strs.
func newKeyString(strs []string) string {
	return strings.Join(strs, tokenKeySeparator)
}

// encodeKeyAndValue encodes key and value to target string and returns an error if failed.
func encodeKeyAndValue(key string, value string) (target string) {
	return key + tokenSeparator + value
}

// encodeEntries encodes entries to target string and returns an error if failed.
func encodeEntries(entries map[string][]string) (target string, err error) {
	var builder strings.Builder

	for key, values := range entries {
		for _, value := range values {
			builder.WriteString(encodeKeyAndValue(key, value))
		}
	}

	return builder.String(), err
}

// decodeLine decodes keys and value from line and returns an error if failed.
func decodeLine(i int, line string) (key string, value string, err error) {
	kv := strings.SplitN(line, tokenSeparator, 2)
	if len(kv) < 2 {
		return "", "", fmt.Errorf("props: line %d {%s} doesn't have keys or value", i, line)
	}

	key = strings.TrimSpace(kv[0])
	value = strings.TrimSpace(strings.TrimSuffix(kv[1], tokenCR)) // Trim cr / space
	return key, value, nil
}

// decodeLine decodes entries from lines and returns an error if failed.
func decodeLines(lines []string) (entries map[string][]string, err error) {
	entries = make(map[string][]string, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if isBlank(line) || isComment(line) {
			continue
		}

		key, value, err := decodeLine(i, line)
		if err != nil {
			return nil, err
		}

		entries[key] = append(entries[key], value)
	}

	return entries, nil
}

// decodeString decodes entries from source string and returns an error if failed.
func decodeString(source string) (entries map[string][]string, err error) {
	return decodeLines(strings.Split(source, tokenLF))
}
