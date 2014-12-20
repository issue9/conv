// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package conv

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 字节单位
const (
	b int64 = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
	// e,z,y  // 这些貌似要超过int64的大小了
)

var byteMap = map[string]int64{
	"b":    b,
	"byte": b,

	"k":  kb,
	"kb": kb,

	"m":  mb,
	"mb": mb,

	"g":  gb,
	"gb": gb,

	"t":  tb,
	"tb": tb,

	"p":  pb,
	"pb": pb,
}

// 0.1M
// 1K
func toByte(v string) (int64, error) {
	index := -1
	for i, val := range v {
		if (val < '0' || val > '9') && val != '.' {
			index = i
			break
		}
	}

	if index <= 0 { // 全是数字，交由conv.Int去转换
		return Int64(v)
	}

	num, err := strconv.ParseFloat(string(v[:index]), 64)
	if err != nil {
		return -1, err
	}

	unitStr := strings.ToLower(v[index:])
	unit, found := byteMap[unitStr]
	if !found {
		return -1, fmt.Errorf("未知的字节单位[%v]", unitStr)
	}

	ret := float64(unit) * num
	if ret < 1 {
		return -1, fmt.Errorf("无意义的大小[%v]", ret)
	}

	// 0.1K == 102.4byte
	// 0.4byte没有实际意义，需要向上取整成103byte
	return int64(math.Ceil(ret)), nil
}

// ToByte 将各种字符串类型转换成字节的单位的数值。
// 以下类型将会被正常转换：
//  "1Kb"  ==> 1024
//  "0.1k" ==> 103  // 0.1k=102.4byte, 向上取整成103
func ToByte(v interface{}) (int64, error) {
	switch vv := v.(type) {
	case []byte:
		return toByte(string(vv))
	case string:
		return toByte(vv)
	case []rune:
		return toByte(string(vv))
	case int64, int32, int16, int8, int, uint64, uint32, uint16, uint8, uint, float32, float64:
		return Int64(v)
	default:
		return -1, fmt.Errorf("[%v]无法转换成数值", v)
	}
}
