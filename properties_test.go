// Copyright 2021 Ye Zi Jie.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/04 23:00:16

package props

import (
	"testing"
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
