// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package conv

import (
	"fmt"
	"strconv"
	"strings"
)

// 抛出一个类型无法转换的错误
// val当前值；t目标类型。
func typeError(val interface{}, t string) error {
	return fmt.Errorf("[%T:%v]无法转换成[%v]类型", val, val, t)
}

// 字符串转Bool值，供Bool()函数调用。
// 添加了一些strconv.ParseFloat不支持但又比较常用的字符串转换
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

// 将val转换成bool类型或是在无法转换的情况下返回error。
// 以下值被可以被正确转换：
//  123(true), 0(false),"-123"(true), "on"(true), "off"(false), "true"(true), "false"(false)
func Bool(val interface{}) (bool, error) {
	switch ret := val.(type) {
	case bool:
		return ret, nil
	//case int, int8, int32, int64, float32, float64, uint, uint8, uint32, uint64:
	//	return ret != 0, nil
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

// 将val转换成bool类型或是在无法转换的情况下返回def参数。
func MustBool(val interface{}, def bool) bool {
	if ret, err := Bool(val); err != nil {
		return def
	} else {
		return ret
	}
}

// 将val转换成uint64类型或是在无法转换的情况下返回error。
// 将一个有符号整数转换成无符号整数，负数将返回错误，正数和零正常转换
func Uint64(val interface{}) (uint64, error) {
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
		if val, err := strconv.ParseFloat(string(ret), 32); err == nil {
			return uint64(val), nil
		}
		return 0, typeError(val, "uint64")
	case string:
		if val, err := strconv.ParseFloat(ret, 32); err == nil {
			return uint64(val), nil
		}
		return 0, typeError(val, "uint64")
	default:
		return 0, typeError(ret, "uint64")
	}
}

// 将val转换成uint64类型或是在无法转换的情况下返回def参数。
func MustUint64(val interface{}, def uint64) uint64 {
	if ret, err := Uint64(val); err != nil {
		return def
	} else {
		return ret
	}
}

// 将val转换成uint类型或是在无法转换的情况下返回error。
func Uint(val interface{}) (uint, error) {
	if ret, err := Uint64(val); err != nil {
		return 0, err
	} else {
		return uint(ret), nil
	}
}

// 将val转换成uint类型或是在无法转换的情况下返回def参数。
func MustUint(val interface{}, def uint) uint {
	if ret, err := Uint64(val); err != nil {
		return def
	} else {
		return uint(ret)
	}
}

// 将val转换成uint8类型或是在无法转换的情况下返回error。
func Uint8(val interface{}) (uint8, error) {
	if ret, err := Uint64(val); err != nil {
		return 0, err
	} else {
		return uint8(ret), nil
	}
}

// 将val转换成uint8类型或是在无法转换的情况下返回def参数。
func MustUint8(val interface{}, def uint8) uint8 {
	if ret, err := Uint64(val); err != nil {
		return def
	} else {
		return uint8(ret)
	}
}

// 将val转换成uint32类型或是在无法转换的情况下返回error。
func Uint32(val interface{}) (uint32, error) {
	if ret, err := Uint64(val); err != nil {
		return 0, err
	} else {
		return uint32(ret), nil
	}
}

// 将val转换成uint32类型或是在无法转换的情况下返回def参数。
func MustUint32(val interface{}, def uint32) uint32 {
	if ret, err := Uint64(val); err != nil {
		return def
	} else {
		return uint32(ret)
	}
}

// 将val转换成int64类型或是在无法转换的情况下返回error。
func Int64(val interface{}) (int64, error) {
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
		} else {
			return 0, nil
		}
	case []byte:
		if val, err := strconv.ParseFloat(string(ret), 32); err == nil {
			return int64(val), nil
		} else {
			return -1, typeError(val, "int64")
		}
	case string:
		if val, err := strconv.ParseFloat(ret, 32); err == nil {
			return int64(val), nil
		} else {
			return -1, typeError(val, "int64")
		}
	default:
		return -1, typeError(ret, "int64")
	}
}

// 将val转换成int64类型或是在无法转换的情况下返回def参数。
func MustInt64(val interface{}, def int64) int64 {
	if ret, err := Int64(val); err != nil {
		return def
	} else {
		return ret
	}
}

// 将val转换成int类型或是在无法转换的情况下返回error。
func Int(val interface{}) (int, error) {
	if ret, err := Int64(val); err != nil {
		return -1, err
	} else {
		return int(ret), err
	}
}

// 将val转换成int类型或是在无法转换的情况下返回def参数。
func MustInt(val interface{}, def int) int {
	if ret, err := Int64(val); err != nil {
		return def
	} else {
		return int(ret)
	}
}

// 将val转换成int8类型或是在无法转换的情况下返回error。
func Int8(val interface{}) (int8, error) {
	if ret, err := Int64(val); err != nil {
		return -1, err
	} else {
		return int8(ret), err
	}
}

// 将val转换成int8类型或是在无法转换的情况下返回def参数。
func MustInt8(val interface{}, def int8) int8 {
	if ret, err := Int64(val); err != nil {
		return def
	} else {
		return int8(ret)
	}
}

// 将val转换成int32类型或是在无法转换的情况下返回error。
func Int32(val interface{}) (int32, error) {
	if ret, err := Int64(val); err != nil {
		return -1, err
	} else {
		return int32(ret), err
	}
}

// 将val转换成int32类型或是在无法转换的情况下返回def参数。
func MustInt32(val interface{}, def int32) int32 {
	if ret, err := Int64(val); err != nil {
		return def
	} else {
		return int32(ret)
	}
}

// 将val转换成float64类型或是在无法转换的情况下返回error。
func Float64(val interface{}) (float64, error) {
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
		} else {
			return 0.0, nil
		}
	case []byte:
		if val, err := strconv.ParseFloat(string(ret), 64); err == nil {
			return float64(val), nil
		} else {
			return -1, typeError(val, "float64")
		}
	case string:
		if val, err := strconv.ParseFloat(ret, 64); err == nil {
			return float64(val), nil
		} else {
			return -1, typeError(val, "float64")
		}
	default:
		return -1, typeError(ret, "float64")
	}
}

// 将val转换成float64类型或是在无法转换的情况下返回def参数。
func MustFloat64(val interface{}, def float64) float64 {
	if ret, err := Float64(val); err != nil {
		return def
	} else {
		return ret
	}
}

// 将val转换成float32类型或是在无法转换的情况下返回error。
func Float32(val interface{}) (float32, error) {
	if ret, err := Float64(val); err != nil {
		return -1.0, err
	} else {
		return float32(ret), nil
	}
}

// 将val转换成float32类型或是在无法转换的情况下返回def参数。
func MustFloat32(val interface{}, def float32) float32 {
	if ret, err := Float64(val); err != nil {
		return def
	} else {
		return float32(ret)
	}
}

// 将val转换成string类型或是在无法转换的情况下返回error。
func String(val interface{}) (string, error) {
	switch ret := val.(type) {
	case string:
		return ret, nil
	case []byte:
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
	default:
		return "", typeError(ret, "string")
	}
}

// 将val转换成string类型或是在无法转换的情况下返回def参数。
func MustString(val interface{}, def string) string {
	if ret, err := String(val); err != nil {
		return def
	} else {
		return ret
	}
}

// 将val转换成[]byte类型或是在无法转换的情况下返回error。
func Bytes(val interface{}) ([]byte, error) {
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

// 将val转换成[]byte类型或是在无法转换的情况下返回def参数。
func MustBytes(val interface{}, def []byte) []byte {
	if ret, err := Bytes(val); err != nil {
		return def
	} else {
		return ret
	}
}

// 将val转换成slice类型或是在无法转换的情况下返回error。
// []int, []interface{}以及数组都可以转换。
// []byte("123")返回[]interface{}{byte(49),byte(50),byte(51)}
// "123"返回[]interface{}{rune(49),rune(50),rune(51)}
func Slice(val interface{}) ([]interface{}, error) {
	switch data := val.(type) {
	case []interface{}:
		return data, nil
	case []int:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	case []int8:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	case []int32:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	case []int64:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	case []uint:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	case []uint8:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	case []uint32:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	case []uint64:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	case []float32:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	case []string:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	case string:
		ret := make([]interface{}, len(data))
		for k, v := range data {
			ret[k] = v
		}
		return ret, nil
	default:
		return nil, typeError(data, "slice")
	}
}

// 将val转换成slice类型或是在无法转换的情况下返回def参数。
func MustSlice(val interface{}, def []interface{}) []interface{} {
	if ret, err := Slice(val); err != nil {
		return def
	} else {
		return ret
	}
}
