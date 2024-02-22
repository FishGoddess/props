// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/FishGoddess/props"
)

func main() {
	// Load properties from a file.
	properties, err := props.LoadFile("load.properties")
	if err != nil {
		panic(err)
	}

	// Get one key from properties and covert it into a integer.
	fmt.Println(properties.Get("key1"))
	fmt.Println(properties.Get("key1").Int())

	// List all entries in properties.
	// Returns false if you want to continue listing.
	properties.List(func(key string, value props.Value) (finished bool) {
		fmt.Println(key, value)
		return false
	})

	// Use StoreFile to store a props to a file.
	// What's more, we provide StoreMap and StoreWriter to store a props to map and writer.
	err = props.StoreFile(properties, "store.properties")
	if err != nil {
		panic(err)
	}
}
