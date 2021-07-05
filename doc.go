// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/06 00:00:52

/*
Package props provides an easy way to use foundation for your properties operations.

1. basic

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
	fmt.Printf("db.user is %s\n", dbUser) // output: db.user is props
	fmt.Printf("web.port is %d\n", webPort) // output: web.port is 8080
	fmt.Printf("web.directories is %+v\n", webDirectories) // output: web.directories is [html  css  js]
	fmt.Printf("notExistedKey is %d\n", notExisted) // output: notExistedKey is 123
	fmt.Printf("notExistedKey is absent? %+v\n", absent) // output: notExistedKey is absent? true

	// Also, we provide a way to get all entries in properties
	properties.Traverse(func(key string, value *props.Value) bool {
		fmt.Printf("key:{%s}, value:{%s}\n", key, value.String(""))
		return true // return true to keep traversing
	})

2. load

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

3. store

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
*/
package props // import "github.com/FishGoddess/props"

const (
	// Version is the version string representation of props.
	Version = "v0.2.0-alpha"
)
