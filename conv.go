// SPDX-License-Identifier: MIT

package conv

import (
	"encoding"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// 抛出一个类型无法转换的错误
// val 当前值；t 目标类型。
func typeError(val any, t string) error {
	return fmt.Errorf("conv: %T:%v 无法转换成 %s 类型", val, val, t)
}

// 字符串转 bool 值，供 Bool() 函数调用。
// 添加了一些 strconv.ParseFloat 不支持但又比较常用的字符串转换
func str2Bool(str string) (bool, error) {
	if val, err := strconv.ParseBool(str); err == nil {
		return val, nil
	} else if val, err := strconv.ParseFloat(str, 32); err == nil {
		return val != 0, nil
	}

	switch strings.ToLower(strings.TrimSpace(str)) {
	case "on":
		return true, nil
	case "off":
		return false, nil
	default:
		return false, typeError(str, "bool")
	}
}

// Bool 将 val 转换成 bool 类型或是在无法转换的情况下返回 error
//
// 以下值被可以被正确转换：
//  123(true), 0(false),"-123"(true), "on"(true), "off"(false), "true"(true), "false"(false)
func Bool(val any) (bool, error) {
	switch ret := val.(type) {
	case bool:
		return ret, nil
	case int:
		return ret != 0, nil
	case int8:
		return ret != 0, nil
	case int32:
		return ret != 0, nil
	case int64:
		return ret != 0, nil
	case float32:
		return ret != 0, nil
	case float64:
		return ret != 0, nil
	case uint:
		return ret != 0, nil
	case uint8:
		return ret != 0, nil
	case uint32:
		return ret != 0, nil
	case uint64:
		return ret != 0, nil
	case []byte:
		return str2Bool(string(ret))
	case string:
		return str2Bool(ret)
	default:
		return false, typeError(val, "bool")
	}
}

// MustBool 将 val 转换成 bool 类型或是在无法转换的情况下返回 def 参数
func MustBool(val any, def ...bool) bool {
	if ret, err := Bool(val); err == nil {
		return ret
	}
	return def[0]
}

// IntOf 转换成指定类型的符号整数
func IntOf[T constraints.Signed](val any) (T, error) {
	ret, err := toInt64(val)
	if err != nil {
		return 1, err
	}
	return T(ret), nil
}

// UintOf 转换成指定类型的无符号整数
//
// 将一个有符号整数转换成无符号整数，负数将返回错误，正数和零正常转换
func UintOf[T constraints.Unsigned](val any) (T, error) {
	ret, err := toUint64(val)
	if err != nil {
		return 0, err
	}
	return T(ret), nil
}

// Uint64 将 val 转换成 uint64 类型或是在无法转换的情况下返回 error
func Uint64(val any) (uint64, error) { return UintOf[uint64](val) }

// MustUint64 将 val 转换成 uint64 类型或是在无法转换的情况下返回 def 参数
func MustUint64(val any, def ...uint64) uint64 { return MustUintOf(val, def...) }

// Uint 将 val 转换成 uint 类型或是在无法转换的情况下返回 error
func Uint(val any) (uint, error) { return UintOf[uint](val) }

// MustUint 将 val 转换成 uint 类型或是在无法转换的情况下返回 def 参数
func MustUint(val any, def ...uint) uint { return MustUintOf(val, def...) }

// Uint8 将 val 转换成 uint8 类型或是在无法转换的情况下返回 error
func Uint8(val any) (uint8, error) { return UintOf[uint8](val) }

// MustUint8 将 val 转换成 uint8 类型或是在无法转换的情况下返回 def 参数
func MustUint8(val any, def ...uint8) uint8 { return MustUintOf(val, def...) }

// Uint32 将 val 转换成 uint32 类型或是在无法转换的情况下返回 error
func Uint32(val any) (uint32, error) { return UintOf[uint32](val) }

// MustUint32 将 val 转换成 uint32 类型或是在无法转换的情况下返回 def 参数
func MustUint32(val any, def ...uint32) uint32 { return MustUintOf(val, def...) }

// Int64 将 val 转换成 int64 类型或是在无法转换的情况下返回 error
func Int64(val any) (int64, error) { return IntOf[int64](val) }

// MustInt64 将 val 转换成 int64 类型或是在无法转换的情况下返回 def 参数
func MustInt64(val any, def ...int64) int64 { return MustIntOf(val, def...) }

// Int 将 val 转换成 int 类型或是在无法转换的情况下返回 error
func Int(val any) (int, error) { return IntOf[int](val) }

// MustInt 将 val 转换成 int 类型或是在无法转换的情况下返回 def 参数
func MustInt(val any, def ...int) int { return MustIntOf(val, def...) }

// Int8 将 val 转换成 int8 类型或是在无法转换的情况下返回 error
func Int8(val any) (int8, error) { return IntOf[int8](val) }

// MustInt8 将 val 转换成 int8 类型或是在无法转换的情况下返回 def 参数
func MustInt8(val any, def ...int8) int8 { return MustIntOf(val, def...) }

// Int32 将 val 转换成 int32 类型或是在无法转换的情况下返回 error
func Int32(val any) (int32, error) { return IntOf[int32](val) }

// MustInt32 将 val 转换成 int32 类型或是在无法转换的情况下返回 def 参数
func MustInt32(val any, def ...int32) int32 { return MustIntOf(val, def...) }

// Float64 将 val 转换成 float64 类型或是在无法转换的情况下返回 error
func Float64(val any) (float64, error) {
	switch ret := val.(type) {
	case float64:
		return ret, nil
	case int:
		return float64(ret), nil
	case int8:
		return float64(ret), nil
	case int32:
		return float64(ret), nil
	case int64:
		return float64(ret), nil
	case uint:
		return float64(ret), nil
	case uint8:
		return float64(ret), nil
	case uint32:
		return float64(ret), nil
	case uint64:
		return float64(ret), nil
	case float32:
		return float64(ret), nil
	case bool:
		if ret {
			return 1.0, nil
		}
		return 0.0, nil
	case []byte:
		val, err := strconv.ParseFloat(string(ret), 64)
		if err == nil {
			return float64(val), nil
		}
		return -1, typeError(val, "float64")
	case string:
		val, err := strconv.ParseFloat(ret, 64)
		if err == nil {
			return float64(val), nil
		}
		return -1, typeError(val, "float64")
	default:
		return -1, typeError(ret, "float64")
	}
}

// MustFloat64 将 val 转换成 float64 类型或是在无法转换的情况下返回 def 参数
func MustFloat64(val any, def ...float64) float64 {
	if ret, err := Float64(val); err == nil {
		return ret
	}
	return def[0]
}

// Float32 将 val 转换成 float32 类型或是在无法转换的情况下返回 error
func Float32(val any) (float32, error) {
	ret, err := Float64(val)
	if err != nil {
		return -1.0, err
	}
	return float32(ret), nil
}

// MustFloat32 将 val 转换成 float32 类型或是在无法转换的情况下返回 def 参数
func MustFloat32(val any, def ...float32) float32 {
	if ret, err := Float64(val); err == nil {
		return float32(ret)
	}
	return def[0]
}

// String 将 val 转换成 string 类型或是在无法转换的情况下返回 error
//
// NOTE: fmt.Stringer, ret.Error 和 encoding.TextMarshaler 都将被正确转换成字符串。
func String(val any) (string, error) {
	switch ret := val.(type) {
	case string:
		return ret, nil
	case []byte:
		return string(ret), nil
	case []rune:
		return string(ret), nil
	case int64:
		return strconv.FormatInt(ret, 10), nil
	case int:
		return strconv.FormatInt(int64(ret), 10), nil
	case int8:
		return strconv.FormatInt(int64(ret), 10), nil
	case int32:
		return strconv.FormatInt(int64(ret), 10), nil
	case uint:
		return strconv.FormatInt(int64(ret), 10), nil
	case uint8:
		return strconv.FormatInt(int64(ret), 10), nil
	case uint32:
		return strconv.FormatInt(int64(ret), 10), nil
	case uint64:
		return strconv.FormatInt(int64(ret), 10), nil
	case float32:
		return strconv.FormatFloat(float64(ret), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(ret, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(ret), nil
	case fmt.Stringer:
		return ret.String(), nil
	case error:
		return ret.Error(), nil
	case encoding.TextMarshaler:
		v, err := ret.MarshalText()
		if err != nil {
			return "", err
		}
		return string(v), nil
	default:
		return "", typeError(ret, "string")
	}
}

// MustString 将 val 转换成 string 类型或是在无法转换的情况下返回 def 参数
func MustString(val any, def ...string) string {
	if ret, err := String(val); err == nil {
		return ret
	}
	return def[0]
}

// Bytes 将 val 转换成 []byte 类型或是在无法转换的情况下返回 error
func Bytes(val any) ([]byte, error) {
	switch ret := val.(type) {
	case []byte:
		return ret, nil
	case string:
		return []byte(ret), nil
	case int64:
		return []byte(strconv.FormatInt(ret, 10)), nil
	case int:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case int8:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case int32:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case uint:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case uint8:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case uint32:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case uint64:
		return []byte(strconv.FormatInt(int64(ret), 10)), nil
	case float32:
		return []byte(strconv.FormatFloat(float64(ret), 'f', 5, 32)), nil
	case float64:
		return []byte(strconv.FormatFloat(ret, 'f', 5, 64)), nil
	case bool:
		return []byte(strconv.FormatBool(ret)), nil
	default:
		return nil, typeError(ret, "[]byte")
	}
}

// MustBytes 将 val 转换成 []byte 类型或是在无法转换的情况下返回 def 参数
func MustBytes(val any, def ...[]byte) []byte {
	if ret, err := Bytes(val); err == nil {
		return ret
	}

	if len(def) == 0 {
		panic(typeError(val, "bytes"))
	}

	return def[0]
}

// Slice 将 val 转换成 slice 类型或是在无法转换的情况下返回 error
//
// []int, []interface{} 以及数组都可以转换。
// []byte("123") 返回 []interface{}{byte(49),byte(50),byte(51)}
// "123" 返回 []interface{}{rune(49),rune(50),rune(51)}
func Slice(val any) ([]any, error) { return SliceOf[any](val) }

// MustSlice 将 val 转换成 slice 类型或是在无法转换的情况下返回 def 参数
func MustSlice(val any, def ...[]any) []any { return MustSliceOf[any](val, def...) }

// SliceOf 将 val 转换成 []T
//
// 只要 val 是数组或是字符串，且其元素能转换成 T 类型即可。
func SliceOf[T any](val any) ([]T, error) {
	srcV := reflect.ValueOf(val)
	switch srcV.Kind() {
	case reflect.String:
		dest := make([]T, srcV.Len())
		destV := reflect.ValueOf(dest)
		destT := destV.Type().Elem()
		if reflect.TypeOf('a').ConvertibleTo(destT) {
			for i := 0; i < srcV.Len(); i++ {
				destV.Index(i).Set(srcV.Index(i).Convert(destT))
			}
		}

		return dest, nil
	case reflect.Slice, reflect.Array:
		dest := make([]T, srcV.Len())
		destV := reflect.ValueOf(dest)
		destT := destV.Type().Elem()

		for i := 0; i < srcV.Len(); i++ {
			v := srcV.Index(i)

			if v.Type().ConvertibleTo(destT) { // srcV 的项类型是确定的，比如 []any，可以包含任何类型。
				destV.Index(i).Set(v.Convert(destT))
			} else {
				if err := Value(v.Interface(), destV.Index(i)); err != nil {
					return nil, err
				}
			}
		}

		return dest, nil
	default:
		return nil, typeError(val, "slice")
	}
}

// MustSliceOf 将 val 转换成 slice 类型或是在无法转换的情况下返回 def 参数
func MustSliceOf[T any](val any, def ...[]T) []T {
	if ret, err := SliceOf[T](val); err == nil {
		return ret
	}

	if len(def) == 0 {
		panic(typeError(val, "slice"))
	}

	return def[0]
}

func toInt64(val any) (int64, error) {
	switch ret := val.(type) {
	case int64:
		return ret, nil
	case int:
		return int64(ret), nil
	case int8:
		return int64(ret), nil
	case int32:
		return int64(ret), nil
	case uint:
		return int64(ret), nil
	case uint8:
		return int64(ret), nil
	case uint32:
		return int64(ret), nil
	case uint64:
		return int64(ret), nil
	case float32:
		return int64(ret), nil
	case float64:
		return int64(ret), nil
	case bool:
		if ret {
			return 1, nil
		}
		return 0, nil
	case []byte:
		val, err := strconv.ParseFloat(string(ret), 64)
		if err == nil {
			return int64(val), nil
		}
		return -1, typeError(val, "int64")
	case string:
		val, err := strconv.ParseFloat(ret, 64)
		if err == nil {
			return int64(val), nil
		}
		return -1, typeError(val, "int64")
	default:
		return -1, typeError(ret, "int64")
	}
}

func toUint64(val any) (uint64, error) {
	switch ret := val.(type) {
	case uint64:
		return ret, nil
	case int:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case int8:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case int32:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case int64:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case uint:
		return uint64(ret), nil
	case uint8:
		return uint64(ret), nil
	case uint32:
		return uint64(ret), nil
	case float32:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case float64:
		if ret < 0 {
			return 0, typeError(ret, "uint64")
		}
		return uint64(ret), nil
	case bool:
		if ret {
			return 1, nil
		}
		return 0, nil
	case []byte:
		if val, err := strconv.ParseFloat(string(ret), 64); err == nil {
			return uint64(val), nil
		}
		return 0, typeError(val, "uint64")
	case string:
		if val, err := strconv.ParseFloat(ret, 64); err == nil {
			return uint64(val), nil
		}
		return 0, typeError(val, "uint64")
	default:
		return 0, typeError(ret, "uint64")
	}
}

// MustIntOf 将 val 转换成 T 类型或是在无法转换的情况下返回 def 参数
func MustIntOf[T constraints.Signed](val any, def ...T) T {
	if ret, err := IntOf[T](val); err == nil {
		return ret
	}
	return def[0]
}

// MustUintOf 将 val 转换成 T 类型或是在无法转换的情况下返回 def 参数
func MustUintOf[T constraints.Unsigned](val any, def ...T) T {
	if ret, err := UintOf[T](val); err == nil {
		return ret
	}
	return def[0]
}
