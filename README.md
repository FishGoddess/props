# ⚙ props

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/props)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![Coverage](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/props/actions/workflows/test.yml/badge.svg)

**props** 是一个用于解析 properties 配置文件的库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md)

### ☁️ 功能特性

* 解析 properties 配置，提供多种类型转换方式，方便配置的类型转换

### 🔨 安装方式

```bash
$ go get -u github.com/FishGoddess/props
```

### 💡 参考案例

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

### 👀 贡献者

如果您觉得 props 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。
