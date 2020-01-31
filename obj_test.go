// SPDX-License-Identifier: MIT

package conv

import (
	"strings"
	"testing"

	"github.com/issue9/assert"
)

type A1 struct {
	ID   int
	Name string
}

type b1 struct {
	A1
	Password string
}

type c1 struct {
	sub      *A1
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
		"ID":   5,
		"Name": "admin",
	}

	// 包含匿名元素
	obja := &A1{}
	err := Map2Obj(m, obja, nil)
	as.Nil(err)
	as.Equal(obja.ID, 5)
	as.Equal(obja.Name, "admin")

	m = map[string]interface{}{
		"ID":       5,
		"Name":     "admin",
		"Password": "password",
		"lower":    "lower",
	}
	objb := &b1{}
	err = Map2Obj(m, objb, nil)
	as.Nil(err)
	as.Equal(objb.ID, 5)
	as.Equal(objb.Name, "admin")
	as.Equal(objb.Password, "password")

	// 包含子元素
	objc := &c1{Sub: &b1{}}
	m = map[string]interface{}{
		"Password": "password",
		"Sub": map[string]interface{}{
			"ID":       6,
			"Name":     "test",
			"Password": "sub-password",
		},
	}
	err = Map2Obj(m, objc, nil)
	as.Nil(err)
	as.Equal(objc.Password, "password")
	as.Equal(objc.Sub.ID, 6)
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
	obja := &A1{6, "admin"}
	m, err := Obj2Map(obja, nil)
	as.Nil(err)
	as.Equal(m["ID"], 6)
	as.Equal(m["Name"], "admin")

	// 包含匿名字段
	objb := &b1{A1{6, "admin"}, "password"}
	m, err = Obj2Map(objb, nil)
	as.Nil(err)
	as.Equal(m["ID"], 6)
	as.Equal(m["Name"], "admin")
	as.Equal(m["Password"], "password")

	// 包含子元素
	objc := &c1{sub: &A1{6, "admin"}, Sub: &b1{A1{5, "test"}, "b-password"}, Password: "password"}
	m, err = Obj2Map(objc, nil)
	as.Nil(err)
	as.Equal(m["Password"], "password")
	sub := m["Sub"].(map[string]interface{})
	as.Equal(sub["Password"], "b-password")
	as.Equal(sub["ID"], 5)
	as.Equal(sub["Name"], "test")

	// 带转换函数
	m, err = Obj2Map(objc, ToUpperFieldConv)
	as.Nil(err)
	as.Equal(m["PASSWORD"], "password")
	sub = m["SUB"].(map[string]interface{})
	as.Equal(sub["ID"], 5)
}
