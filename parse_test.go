// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/04 23:00:16

package props

import "testing"

// go test -v -cover -run=^TestParseFromString$
func TestParseFromString(t *testing.T) {

	properties, err := parseFromString("key1=value1\nkey2=value2\nkey3=value3\n")
	if err != nil {
		t.Fatalf("parseFromString returns an error %+v", err)
	}

	value := properties.Get("key1").String("")
	if value != "value1" {
		t.Fatalf("get key1 returns wrong result %s", value)
	}

	value = properties.Get("key2").String("")
	if value != "value2" {
		t.Fatalf("get key2 returns wrong result %s", value)
	}

	value = properties.Get("key3").String("")
	if value != "value3" {
		t.Fatalf("get key3 returns wrong result %s", value)
	}
}

// go test -v -cover -run=^TestParseFromProperties$
func TestParseFromProperties(t *testing.T) {

	properties := NewProperties()
	properties.Set("key1", "value1")
	properties.Set("key2", "value2")
	properties.Set("key3", "value3")

	str := parseFromProperties(properties)
	if str != "key1=value1\nkey2=value2\nkey3=value3\n" {
		t.Fatalf("parseFromProperties returns wrong result %s", str)
	}
}
