// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package conv

import (
	"strings"
	"testing"

	"github.com/issue9/assert"
)

type a1 struct {
	Id   int
	Name string
}

type b1 struct {
	a1
	Password string
}

type c1 struct {
	sub      *a1
	Sub      *b1
	Password string
}

type C struct {
	SUB      *b1
	PASSWORD string
}

func ToUpperFieldConv(str string) string {
	return strings.ToUpper(str)
}

func TestMap2Obj(t *testing.T) {
	as := assert.New(t)

	// 一般
	m := map[string]interface{}{
		"Id":   5,
		"Name": "admin",
	}

	// 包含匿名元素
	obja := &a1{}
	err := Map2Obj(m, obja, nil)
	as.Nil(err)
	as.Equal(obja.Id, 5)
	as.Equal(obja.Name, "admin")

	m = map[string]interface{}{
		"Id":       5,
		"Name":     "admin",
		"Password": "password",
		"lower":    "lower",
	}
	objb := &b1{}
	err = Map2Obj(m, objb, nil)
	as.Nil(err)
	as.Equal(objb.Id, 5)
	as.Equal(objb.Name, "admin")
	as.Equal(objb.Password, "password")

	// 包含子元素
	objc := &c1{Sub: &b1{}}
	m = map[string]interface{}{
		"Password": "password",
		"Sub": map[string]interface{}{
			"Id":       6,
			"Name":     "test",
			"Password": "sub-password",
		},
	}
	err = Map2Obj(m, objc, nil)
	as.Nil(err)
	as.Equal(objc.Password, "password")
	as.Equal(objc.Sub.Id, 6)
	as.Equal(objc.Sub.Password, "sub-password")

	// 带转换函数
	objC := &C{SUB: &b1{}}
	err = Map2Obj(m, objC, ToUpperFieldConv)
	as.Nil(err)
	as.Equal(objC.PASSWORD, "password")
	as.NotNil(objC.SUB)
}

func TestObj2Map(t *testing.T) {
	as := assert.New(t)

	// 普通
	obja := &a1{6, "admin"}
	m, err := Obj2Map(obja, nil)
	as.Nil(err)
	as.Equal(m["Id"], 6)
	as.Equal(m["Name"], "admin")

	// 包含匿名字段
	objb := &b1{a1{6, "admin"}, "password"}
	m, err = Obj2Map(objb, nil)
	as.Nil(err)
	as.Equal(m["Id"], 6)
	as.Equal(m["Name"], "admin")
	as.Equal(m["Password"], "password")

	// 包含子元素
	objc := &c1{sub: &a1{6, "admin"}, Sub: &b1{a1{5, "test"}, "b-password"}, Password: "password"}
	m, err = Obj2Map(objc, nil)
	as.Nil(err)
	as.Equal(m["Password"], "password")
	sub := m["Sub"].(map[string]interface{})
	as.Equal(sub["Password"], "b-password")
	as.Equal(sub["Id"], 5)
	as.Equal(sub["Name"], "test")

	// 带转换函数
	m, err = Obj2Map(objc, ToUpperFieldConv)
	as.Nil(err)
	as.Equal(m["PASSWORD"], "password")
	sub = m["SUB"].(map[string]interface{})
	as.Equal(sub["ID"], 5)
}
