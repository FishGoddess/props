// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strings"

	"github.com/FishGoddess/props"
)

func main() {

	// As you know, we provide a LoadFromFile function to load properties from file
	//
	//    properties, err := props.LoadFromFile("xxx.properties")
	//
	// However, not all properties are stored in a file, some of them may be stored in database, even just a string
	// So we provide other ways to load properties
	reader := strings.NewReader("key1=value1")
	properties, err := props.LoadFrom(reader)
	if err != nil {
		panic(err)
	}

	value := properties.Get("key1").String("")
	fmt.Println(value) // output: value1

	properties, err = props.Load("key2=value2")
	if err != nil {
		panic(err)
	}

	value = properties.Get("key2").String("")
	fmt.Println(value) // output: value2
}
