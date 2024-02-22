# ⚙ props

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/props)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/props/actions/workflows/test.yml/badge.svg)

**props** is a properties lib for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md)

### ☁️ Features

* Parsing properties configurations, type conversion supports

### 🔨 Installation

```bash
$ go get -u github.com/FishGoddess/props
```

### 💡 Examples

```go
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
```

### 👀 Contributors

If you find that something is not working as expected please open an _**issue**_.
