# 📜 props

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/props)
[![License](_icons/license.svg)](https://opensource.org/licenses/MIT)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**props** 是一个用于解析 properties 配置文件的库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md)

### ✏ 功能特性

* 解析 properties 配置，提供多种类型转换方式，方便配置的类型转换
* 使用默认值机制替代 ok 判断机制，代码更简洁
* 支持从字符串、io.Reader、文件中加载 properties 数据
* 支持将 properties 数据回写到字符串、io.Writer、文件中

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

	// LoadFromFile loads a properties from file
	//
	// basic.properties:
	//    # database properties
	//    db.user = props
	//    db.password = 123456
	//
	//    # web properties
	//    web.port = 8080
	//    web.directories = html, css, js
	properties, err := props.LoadFromFile("basic.properties")
	if err != nil {
		panic(err)
	}

	// You can use Get method to get whatever in properties
	// We notice that types in properties is insufficient, so we provides some type conversions for you
	// We also notice that the key passed to Get can be absent, but we don't want you to use if-else to check if the key exists or not
	// So we add a defaultValue to all type conversions and the defaultValue will be returned if the key doesn't exist
	dbUser := properties.Get("db.user").String("unknown")
	webPort := properties.Get("web.port").Int(80)
	webDirectories := properties.Get("web.directories").Strings(",", []string{"a", "b", "c", "d"})
	notExisted := properties.Get("notExistedKey").Int(123)
	absent := properties.Get("notExistedKey").Absent()

	// Now check what we got
	fmt.Printf("db.user is %s\n", dbUser)                  // output: db.user is props
	fmt.Printf("web.port is %d\n", webPort)                // output: web.port is 8080
	fmt.Printf("web.directories is %+v\n", webDirectories) // output: web.directories is [html  css  js]
	fmt.Printf("notExistedKey is %d\n", notExisted)        // output: notExistedKey is 123
	fmt.Printf("notExistedKey is absent? %+v\n", absent)   // output: notExistedKey is absent? true

	// Also, we provide a way to get all entries in properties
	properties.Traverse(func(key string, value *props.Value) bool {
		fmt.Printf("key:{%s}, value:{%s}\n", key, value.String(""))
		return true // return true to keep traversing
	})
}
```

* [basic](./_examples/basic.go)
* [load](./_examples/load.go)
* [store](./_examples/store.go)

### 👀 贡献者

如果您觉得 props 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 🎁 使用 logit 的项目

* 待更新。。。