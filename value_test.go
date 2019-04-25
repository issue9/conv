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

	s3 := "34"
	t3 := []byte("def")
	a.NotError(Value(s3, reflect.ValueOf(&t3)))
	a.Equal(s3, string(t3))
	s3_0 := []string{"1", "2"}
	a.NotError(Value(s3_0, reflect.ValueOf(&t3)))
	a.Equal([]byte{1, 2}, string(t3))

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

	t12 := "22"
	a.NotError(Value(nil, reflect.ValueOf(&t12)))
	a.Equal("", t12)

	s13 := []string{"4", "5", "6"}
	t13 := []int{1, 2}
	a.NotError(Value(s13, reflect.ValueOf(&t13)))
	a.Equal([]int{4, 5, 6}, t13)

	// array 转换
	s14 := []string{"4", "5"}
	t14 := [2]int{1, 2}
	a.NotError(Value(s14, reflect.ValueOf(&t14)))
	a.Equal([]int{4, 5}, t14)

	// array 长度不一样，无法转换
	s15 := []string{"4", "5", "6"}
	t15 := [2]int{1, 2}
	a.Error(Value(s15, reflect.ValueOf(&t15)))
	a.Equal([]int{1, 2}, t15)

	// slice 长度不同，可以转换
	s16 := []string{"4", "5"}
	t16 := []int{1, 2, 3}
	a.NotError(Value(s16, reflect.ValueOf(&t16)))
	a.Equal([]int{4, 5}, t16)

	s17 := []string{"4", "5"}
	t17 := []byte{1, 2, 3}
	a.NotError(Value(s17, reflect.ValueOf(&t17)))
	a.Equal([]byte{4, 5}, t17)

	// 无法转换的
	s20 := "1a23"
	t20 := 444
	a.Error(Value(s20, reflect.ValueOf(&t20)))
}
