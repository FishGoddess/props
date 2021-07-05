// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/06 00:14:04

package main

import (
	"fmt"
	"os"

	"github.com/FishGoddess/props"
)

func main() {

	// We provide some ways to store properties to somewhere
	// At first, create a empty properties
	properties := props.NewProperties()

	// Add some properties to it
	// Set sets key and value to properties and returns true forever
	// A new value will be added if key doesn't exist
	// The old value will be replaced if key exists
	added := properties.Set("key1", "value") // The value of key1 is value this moment
	fmt.Println(added)                       // output: true
	added = properties.Set("key1", "value1") // Now, the value of key1 is value1
	fmt.Println(added)                       // output: true

	// SetExisted sets key-value only if the key exists
	// If not, it will return false
	added = properties.SetExisted("key2", "value2")
	fmt.Println(added) // output: false

	// SetNotExisted sets key-value only if the key doesn't exist
	// If exists, it will return false
	added = properties.SetNotExisted("key1", "value3")
	fmt.Println(added) // output: false

	// If you want to store properties to string, try this:
	str := properties.Store()
	fmt.Print(str) // output: key1=value1

	// If you want to store properties to io.Writer, try this:
	err := properties.StoreTo(os.Stdout) // output: key1=value1
	if err != nil {
		panic(err)
	}

	// If you want to store properties to a file, try this:
	err = properties.StoreToFile("store.properties")
	if err != nil {
		panic(err)
	}
}
