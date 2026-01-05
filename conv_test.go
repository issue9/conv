// SPDX-FileCopyrightText: 2014-2026 caixw
//
// SPDX-License-Identifier: MIT

package conv

import (
	"testing"

	"github.com/issue9/assert/v4"
)

func TestMustBool(t *testing.T) {
	a := assert.New(t, false)

	// 可解析
	a.True(MustBool("on", false))
	a.True(MustBool(true, false))
	a.True(MustBool("123", false))
	a.True(MustBool(123, false))
	a.True(MustBool(-123, false))
	a.True(MustBool(-1.23, false))

	a.False(MustBool("off", true))
	a.False(MustBool(false, true))

	a.True(MustBool("1"))
	a.True(MustBool("true"))
	a.False(MustBool("0"))
	a.False(MustBool("false"))

	// 不可解析
	a.False(MustBool("str", false))
	a.True(MustBool(";adf", true))
	a.Panic(func() {
		MustBool("str")
	})
}

func TestMustInt(t *testing.T) {
	a := assert.New(t, false)

	// 可解析
	a.Equal(MustInt("123", 456), 123)
	a.Equal(MustInt(true, 5), 1)
	a.Equal(MustInt(false, 5), 0)
	a.Equal(MustInt(123, 456), 123)
	a.Equal(MustInt(uint(8), 9), 8)
	a.Equal(MustInt(uint64(33), 99), 33)
	a.Equal(MustInt(int64(-123), 99), -123)
	a.Equal(MustInt(1.23, 99), 1)
	a.Equal(MustInt(-1.23, 99), -1)
	a.Equal(MustInt([]byte("123"), 44), 123)
	a.Equal(MustInt("1"), 1)

	// 不可解析
	a.Equal(MustInt(";sdf", 45), 45)
	a.Equal(MustInt("str", 45), 45)
	a.Panic(func() {
		MustInt(";str")
	})
}

func TestMustInt64(t *testing.T) {
	// 可解析
	a := assert.New(t, false)

	a.Equal(MustInt64("123", 456), int64(123))
	a.Equal(MustInt64(true, 5), int64(1))
	a.Equal(MustInt64(false, 5), int64(0))
	a.Equal(MustInt64(123, 456), int64(123))
	a.Equal(MustInt64(uint(8), 9), int64(8))
	a.Equal(MustInt64(uint64(33), 99), int64(33))
	a.Equal(MustInt64(int64(-123), 99), int64(-123))
	a.Equal(MustInt64(1.23, 99), int64(1))
	a.Equal(MustInt64(-1.23, 99), int64(-1))
	a.Equal(MustInt64([]byte("123"), 44), int64(123))
	a.Equal(MustInt64([]byte("1607957811"), 44), int64(1607957811))
	a.Equal(MustInt64("1"), 1)

	// 不可解析
	a.Equal(MustInt64(";sdf", 45), int64(45))
	a.Equal(MustInt64("str", 45), 45)
	a.Panic(func() {
		MustInt64("str")
	})
}

func TestMustUint32(t *testing.T) {
	a := assert.New(t, false)
	// 可解析
	a.Equal(MustUint32("123", 456), uint32(123))
	a.Equal(MustUint32(true, 5), uint32(1))
	a.Equal(MustUint32(false, 5), uint32(0))
	a.Equal(MustUint32(123, 456), uint32(123))
	a.Equal(MustUint32(uint(8), 9), uint32(8))
	a.Equal(MustUint32(uint64(33), 99), uint32(33))
	a.Equal(MustUint32(1.23, 99), uint32(1))
	a.Equal(MustUint32([]byte("123"), 44), uint32(123))
	a.Equal(MustUint32("1"), 1)

	// 不可解析
	a.Equal(MustUint32(int64(-123), 99), uint32(99))
	a.Equal(MustUint32(-1.23, 99), uint32(99))
	a.Equal(MustUint32(";sdf", 45), uint32(45))
	a.Equal(MustUint32("str", 45), uint32(45))
	a.Panic(func() {
		MustUint32("str")
	})
}

func TestMustUint64(t *testing.T) {
	a := assert.New(t, false)

	// 可解析
	a.Equal(MustUint64("123", 456), uint64(123))
	a.Equal(MustUint64(true, 5), uint64(1))
	a.Equal(MustUint64(false, 5), uint64(0))
	a.Equal(MustUint64(123, 456), uint64(123))
	a.Equal(MustUint64(uint(8), 9), uint64(8))
	a.Equal(MustUint64(uint64(33), 99), uint64(33))
	a.Equal(MustUint64(1.23, 99), uint64(1))
	a.Equal(MustUint64([]byte("123"), 44), uint64(123))
	a.Equal(MustUint64("1"), 1)

	// 不可解析
	a.Equal(MustUint64(int64(-123), 99), uint64(99))
	a.Equal(MustUint64(-1.23, 99), uint64(99))
	a.Equal(MustUint64(";sdf", 45), uint64(45))
	a.Equal(MustUint64("str", 45), uint64(45))
	a.Panic(func() {
		MustUint64("str")
	})
}

func TestMustFloat32(t *testing.T) {
	a := assert.New(t, false)
	// 可解析
	a.Equal(MustFloat32("123", 456), float32(123))
	a.Equal(MustFloat32(true, 5), float32(1))
	a.Equal(MustFloat32(false, 5), float32(0))
	a.Equal(MustFloat32(123, 456), float32(123))
	a.Equal(MustFloat32(int64(-123), 99), float32(-123.0))
	a.Equal(MustFloat32(uint(8), 9), float32(8))
	a.Equal(MustFloat32(uint64(33), 99), float32(33))
	a.Equal(MustFloat32([]byte("123"), 44), float32(123))
	a.Equal(MustFloat32("1"), 1)

	// 不可解析
	a.Equal(MustFloat32(";sdf", 45), 45)
	a.Equal(MustFloat32("str", 45), 45)
	a.Panic(func() {
		MustFloat32("str")
	})
}

func TestMustString(t *testing.T) {
	a := assert.New(t, false)
	// 可解析
	a.Equal(MustString(123, "222"), "123")
	a.Equal(MustString(-11, "22"), "-11")
	a.Equal(MustString(-11.111, "22"), "-11.111")
	a.Equal(MustString(true, "22"), "true")
	a.Equal(MustString(1), "1")

	// 不可解析
	a.Equal(MustString([]int{1}, "22"), "22")
	a.Panic(func() {
		MustString([]int{1})
	})
}

func TestMustBytes(t *testing.T) {
	a := assert.New(t, false)
	// 可解析
	a.Equal(MustBytes("123", []byte("456")), []byte{49, 50, 51})
	a.Equal(MustBytes(123, []byte("456")), []byte{49, 50, 51})
	a.Equal(MustBytes(uint(123), []byte("456")), []byte{49, 50, 51})
	a.Equal(MustBytes(123), []byte("123"))

	// 不可解析
	a.Equal(MustBytes([]int{1}, []byte("123")), []byte{49, 50, 51})
	a.PanicString(func() {
		MustBytes([]int{1})
	}, "conv: []int:[1] 无法转换成 bytes 类型")
}

func TestMustSlice(t *testing.T) {
	def := []any{4, 5, 5}
	a := assert.New(t, false)
	// 可解析
	a.Equal(MustSlice([]int{1, 2, 3}, def), []any{int(1), int(2), int(3)})
	a.Equal(MustSlice([]uint64{1, 2, 3}, def), []any{uint64(1), uint64(2), uint64(3)})
	a.Equal(MustSlice([]any{"1", 2, 3.0}, def), []any{"1", 2, 3.0})
	a.Equal(MustSlice([]string{"1", "2", "3"}, def), []any{"1", "2", "3"})

	a.Equal(MustSlice([]byte("123"), def), []any{byte(49), byte(50), byte(51)})
	a.Equal(MustSlice("123", def), []any{rune(49), rune(50), rune(51)})
	a.Equal(MustSlice([]int{1, 2, 3}), []any{int(1), int(2), int(3)})
	a.Equal(MustSlice([]uint64{1, 2, 3}), []any{uint64(1), uint64(2), uint64(3)})

	// 不可解析
	a.Equal(MustSlice(7, def), []any{4, 5, 5})
	a.PanicString(func() {
		MustSlice(7)
	}, "conv: int:7 无法转换成 slice 类型")
}

func TestBool(t *testing.T) {
	a := assert.New(t, false)
	fn := func(val any, result bool) {
		ret, err := Bool(val)
		a.Equal(ret, result)
		a.Nil(err)
	}

	fn(5, true)
	fn(int64(-11), true)
	fn(int32(5), true)
	fn(uint(3), true)
	fn(uint64(0), false)
	fn(1.32, true)
	fn(0.00, false)
	fn("off", false)
	fn("true", true)
	fn(-1.3, true)
}

func TestInt(t *testing.T) {
	a := assert.New(t, false)

	fn := func(val any, result int) {
		ret, err := Int(val)
		a.Nil(err)
		a.Equal(ret, result)

	}

	fn(5, 5)
	fn(int64(-11), -11)
	fn(int32(5), 5)
	fn(uint(3), 3)
	fn(uint64(0), 0)
	fn(1.32, 1)
	fn(0.00, 0)
	fn(-1.3, -1)
	fn("-1.1", -1)
	fn("123", 123)
}

func TestInt64(t *testing.T) {
	a := assert.New(t, false)

	fn := func(val any, result int64) {
		ret, err := Int64(val)
		a.Nil(err)
		a.Equal(ret, result)
	}

	fn(5, 5)
	fn(int(-11), -11)
	fn(int32(5), 5)
	fn(uint(3), 3)
	fn(uint64(0), 0)
	fn(1.32, 1)
	fn(0.00, 0)
	fn(-1.3, -1)
	fn("-1.1", -1)
	fn("9223372036854775807", 9223372036854775807)
	fn("-9223372036854775808", -9223372036854775808)
	fn("123", 123)
}

func TestUint(t *testing.T) {
	a := assert.New(t, false)

	fn := func(val any, result uint) {
		ret, err := Uint(val)
		a.Nil(err)
		a.Equal(ret, result)
	}

	fn(5, 5)
	fn(int32(5), 5)
	fn(uint64(0), 0)
	fn(1.32, 1)
	fn(0.00, 0)
	fn("1.1", 1)
	fn("18446744073709551615", 18446744073709551615) // 64 位最大值
	fn("123", 123)
}

func TestUint64(t *testing.T) {
	a := assert.New(t, false)

	fn := func(val any, result uint64) {
		ret, err := Uint64(val)
		a.Nil(err)
		a.Equal(ret, result)
	}

	fn(5, 5)
	fn(int32(5), 5)
	fn(uint(0), 0)
	fn(1.32, 1)
	fn(0.00, 0)
	fn("1.1", 1)
	fn("123", 123)
}

func TestFloat32(t *testing.T) {
	a := assert.New(t, false)

	fn := func(val any, result float32) {
		ret, err := Float32(val)
		a.Nil(err)
		a.Equal(ret, result)
	}

	fn(5, 5)
	fn(int32(5), 5)
	fn(uint(0), 0)
	fn(1.32, 1.32)
	fn("0.00", 0)
	fn("1.1", 1.1)
	fn("123", 123)
}

func TestString(t *testing.T) {
	a := assert.New(t, false)

	fn := func(val any, result string) {
		ret, err := String(val)
		a.Nil(err)
		a.Equal(ret, result)
	}

	fn(5, "5")
	fn(int32(5), "5")
	fn(uint(0), "0")
	fn(1.32, "1.32")
	fn("0.00", "0.00")
	fn(-1.1, "-1.1")
	fn(1.0, "1")
}

func TestBytes(t *testing.T) {
	a := assert.New(t, false)

	fn := func(val any, result []byte) {
		ret, err := Bytes(val)
		a.Nil(err)
		a.Equal(ret, result)
	}

	fn(11, []byte{49, 49})
	fn("1.11", []byte{49, 46, 49, 49})
	fn(-1.11, []byte{45, 49, 46, 49, 49, 48, 48, 48})
	fn(0, []byte{48})
}

func TestSlice(t *testing.T) {
	a := assert.New(t, false)

	fn := func(val any, result []any) {
		ret, err := Slice(val)
		a.Nil(err)
		a.Equal(ret, result)
	}

	fn("123", []any{int32(49), int32(50), int32(51)})
	fn([]int{1, 2, 3}, []any{int(1), int(2), int(3)})
	fn([]string{"1", "ss"}, []any{"1", "ss"})
}

func TestSliceOf(t *testing.T) {
	a := assert.New(t, false)

	// int ==>int
	ret1, err := SliceOf[int]([]int{1, 2, 3})
	a.NotError(err).Equal(ret1, []int{1, 2, 3})

	// int ==> int8
	ret2, err := SliceOf[int8]([]int{1, 2, 3})
	a.NotError(err).Equal(ret2, []int8{1, 2, 3})

	// int ==> any
	ret3, err := SliceOf[any]([]int{1, 2, 3})
	a.NotError(err).Equal(ret3, []any{1, 2, 3})

	// int ==> int8，溢出。
	ret4, err := SliceOf[int8]([]int{1000, 2, 3})
	a.NotError(err).Equal(ret4, []int8{-24, 2, 3})

	// []string ==> int
	ret5, err := SliceOf[int]([]string{"1", "2", "3"})
	a.NotError(err).Equal(ret5, []int{1, 2, 3})

	// any ==> int
	ret6, err := SliceOf[int]([]any{1, 2, "3"})
	a.NotError(err).Equal(ret6, []int{1, 2, 3})

	// string ==> int
	ret7, err := SliceOf[int]("123")
	a.NotError(err).Equal(ret7, []int{'1', '2', '3'})

	// string ==> byte
	ret8, err := SliceOf[byte]("123")
	a.NotError(err).Equal(ret8, []byte{'1', '2', '3'})
}
