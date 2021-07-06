// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/05 23:17:59

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
	//    web.directories = html,css,js
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
	fmt.Printf("web.directories is %+v\n", webDirectories) // output: web.directories is [html css js]
	fmt.Printf("notExistedKey is %d\n", notExisted)        // output: notExistedKey is 123
	fmt.Printf("notExistedKey is absent? %+v\n", absent)   // output: notExistedKey is absent? true

	// Also, we provide a way to get all entries in properties
	properties.Traverse(func(key string, value *props.Value) bool {
		fmt.Printf("key:{%s}, value:{%s}\n", key, value.String(""))
		return true // return true to keep traversing
	})
}
