// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package conv

import (
	"reflect"
	"testing"

	"github.com/issue9/assert"
)

func TestValue(t *testing.T) {
	a := assert.New(t)

	s1 := int(5)
	t1 := int8(7)
	a.NotError(Value(s1, reflect.ValueOf(&t1)))
	a.Equal(s1, t1)

	s2 := "abc"
	t2 := "def"
	a.NotError(Value(s2, reflect.ValueOf(&t2)))
	a.Equal(s2, t2)

	s3 := "abc"
	t3 := []byte("def")
	a.NotError(Value(s3, reflect.ValueOf(&t3)))
	a.Equal(s3, string(t3))

	s4 := []byte("abc")
	t4 := "def"
	a.NotError(Value(s4, reflect.ValueOf(&t4)))
	a.Equal(s4, []byte(t4))

	s5 := "on"
	t5 := false
	a.NotError(Value(s5, reflect.ValueOf(&t5)))
	a.Equal(true, t5)

	s6 := "123"
	t6 := 456
	a.NotError(Value(s6, reflect.ValueOf(&t6)))
	a.Equal(123, t6)

	s7 := float64(0)
	t7 := int64(1)
	a.NotError(Value(s7, reflect.ValueOf(&t7)))
	a.Equal(int64(0), t7)

	s8 := byte('1')
	t8 := byte('2')
	a.NotError(Value(s8, reflect.ValueOf(&t8)))
	a.Equal(byte('1'), t8)

	s9 := byte('1')
	t9 := int('2')
	a.NotError(Value(s9, reflect.ValueOf(&t9)))
	a.Equal(49, t9)

	var s10 interface{}
	s10 = '1'
	t10 := int('2')
	a.NotError(Value(s10, reflect.ValueOf(&t10)))
	a.Equal(49, t10)

	t11 := int('2')
	a.NotError(Value(nil, reflect.ValueOf(&t11)))
	a.Equal(0, t11)

	// 无法转换的
	s20 := "1a23"
	t20 := 444
	a.Error(Value(s20, reflect.ValueOf(&t20)))
}
