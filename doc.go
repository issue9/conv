// SPDX-License-Identifier: MIT

// Package conv 提供了基础的类型转换功能
//
// 会尽最大可能地将当前的值转换成指定类型的值。
//
//	conv.MustInt("123", 0)  // 返回 123 的数值
//	conv.MustString(123, "")// 返回字符串 123
//	conv.Int("123", 0)      // 返回 123 数值和 nil 的 error 接口
//	v := 5
//	conv.Value("3", reflect.ValueOf(v)) // 将 3 转换成数值,并写入 v 中。
package conv
