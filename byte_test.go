// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package conv

import (
	"testing"

	"github.com/issue9/assert"
)

func TestUnit(t *testing.T) {
	assert.Equal(t, b, 1)
	assert.Equal(t, kb, 1024)
	assert.Equal(t, mb, 1024*1024)
	assert.Equal(t, gb, 1024*1024*1024)
	assert.Equal(t, tb, 1024*1024*1024*1024)
}

func TestToByte(t *testing.T) {
	a := assert.New(t)

	fn := func(v interface{}, wont int64) {
		ret, err := ToByte(v)
		a.NotError(err)
		a.NotEqual(-1, ret)

		a.Equal(ret, wont, "[%v]!=[%v],实际值为[%v]", v, wont, ret)
	}

	fn("1024", 1024)
	fn("1K", 1024)
	fn("0.1K", 103)
	fn("1M", 1024*1024)
	fn("5M", 1024*1024*5)
	fn("5G", 5*1024*1024*1024)
}
