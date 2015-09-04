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

	// 无法转换的
	s20 := "1a23"
	t20 := 444
	a.Error(Value(s20, reflect.ValueOf(&t20)))
}
