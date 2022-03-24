// Copyright 2021 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package props

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// go test -v -cover -run=^TestPropertiesGet$
func TestPropertiesGet(t *testing.T) {

	properties := NewProperties()
	properties.Set("key", "value")

	if value := properties.Get("notExistedKey"); !value.Absent() {
		t.Fatalf("get a not existed key is %+v", value)
	}

	if value := properties.Get("key"); value.Absent() || value.String("") != "value" {
		t.Fatalf("value of key (%s) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesSet$
func TestPropertiesSet(t *testing.T) {

	properties := NewProperties()
	added := properties.Set("key", "value")
	if !added {
		t.Fatalf("set key returns %+v", added)
	}

	value := properties.Get("key").String("")
	if value != "value" {
		t.Fatalf("value of key (%s) is wrong", value)
	}

	added = properties.Set("key", "value123")
	if !added {
		t.Fatalf("set key returns %+v", added)
	}

	value = properties.Get("key").String("")
	if value != "value123" {
		t.Fatalf("value of key (%s) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesSetExisted$
func TestPropertiesSetExisted(t *testing.T) {

	properties := NewProperties()
	added := properties.SetExisted("key", "value")
	if added {
		t.Fatalf("set existed key returns %+v", added)
	}

	absent := properties.Get("key").Absent()
	if !absent {
		t.Fatalf("get absent key returns %+v", absent)
	}

	added = properties.Set("key", "value")
	if !added {
		t.Fatalf("set key returns %+v", added)
	}

	value := properties.Get("key").String("")
	if value != "value" {
		t.Fatalf("value of key (%s) is wrong", value)
	}

	added = properties.SetExisted("key", "value123")
	if !added {
		t.Fatalf("set key returns %+v", added)
	}

	value = properties.Get("key").String("")
	if value != "value123" {
		t.Fatalf("value of key (%s) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesSetNotExisted$
func TestPropertiesSetNotExisted(t *testing.T) {

	properties := NewProperties()
	added := properties.Set("key", "value")
	if !added {
		t.Fatalf("set key returns %+v", added)
	}

	value := properties.Get("key").String("")
	if value != "value" {
		t.Fatalf("value of key (%s) is wrong", value)
	}

	added = properties.SetNotExisted("key", "value")
	if added {
		t.Fatalf("set not existed key returns %+v", added)
	}

	added = properties.SetNotExisted("key1", "value1")
	if !added {
		t.Fatalf("set key returns %+v", added)
	}

	value = properties.Get("key1").String("")
	if value != "value1" {
		t.Fatalf("value of key1 (%s) is wrong", value)
	}
}

// go test -v -cover -run=^TestPropertiesTraverse$
func TestPropertiesTraverse(t *testing.T) {

	properties := NewProperties()
	properties.Set("key1", "value1")
	properties.Set("key2", "value2")
	properties.Set("key3", "value3")

	buffer := make([]byte, 0, 512)
	properties.Traverse(func(key string, value *Value) bool {
		buffer = append(buffer, key...)
		buffer = append(buffer, keyValueSeparator...)
		buffer = append(buffer, value.get()...)
		buffer = append(buffer, lf...)
		return true
	})

	str := string(buffer)
	if str != "key1=value1\nkey2=value2\nkey3=value3\n" {
		t.Fatalf("traverse returns wrong result %s", str)
	}
}

// go test -v -cover -run=^TestPropertiesStore$
func TestPropertiesStore(t *testing.T) {

	properties := NewProperties()
	properties.Set("key1", "value1")
	properties.Set("key2", "value2")
	properties.Set("key3", "value3")

	str := properties.Store()
	if str != "key1=value1\nkey2=value2\nkey3=value3\n" {
		t.Fatalf("store returns wrong result %s", str)
	}
}

// go test -v -cover -run=^TestPropertiesStoreTo$
func TestPropertiesStoreTo(t *testing.T) {

	properties := NewProperties()
	properties.Set("key1", "value1")
	properties.Set("key2", "value2")
	properties.Set("key3", "value3")

	buffer := bytes.NewBuffer(make([]byte, 0, 64))
	err := properties.StoreTo(buffer)
	if err != nil {
		t.Fatalf("store to returns an error %+v", err)
	}

	str := buffer.String()
	if str != "key1=value1\nkey2=value2\nkey3=value3\n" {
		t.Fatalf("store to returns wrong result %s", str)
	}
}

// go test -v -cover -run=^TestPropertiesStoreToFile$
func TestPropertiesStoreToFile(t *testing.T) {

	properties := NewProperties()
	properties.Set("key1", "value1")
	properties.Set("key2", "value2")
	properties.Set("key3", "value3")

	fileName := filepath.Join(os.TempDir(), "TestPropertiesStoreToFile_"+time.Now().Format("20060102150405000"+".properties"))

	t.Logf("TestPropertiesStoreToFile test file is %s", fileName)
	err := properties.StoreToFile(fileName)
	if err != nil {
		t.Fatalf("store to file returns an error %+v", err)
	}

	buffer, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatalf("read file returns an error %+v", err)
	}

	str := string(buffer)
	if str != "key1=value1\nkey2=value2\nkey3=value3\n" {
		t.Fatalf("store to file returns wrong result %s", str)
	}
}
