// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package internal

import "strings"

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
