// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package conv

import (
	"errors"
	"fmt"
	"reflect"
)

// Value 将 source 的值保存到成 target 中。
//
// 若类型不能直接转换，会尝试其它种方式转换，比如 strconv.ParseInt() 等。
func Value(source interface{}, target reflect.Value) error {
	kind := target.Kind()

	for kind == reflect.Ptr {
		target = target.Elem()
		kind = target.Kind()
	}

	if !target.CanSet() {
		return fmt.Errorf("无法改变target的值[%v]", target.Kind())
	}

	if !target.IsValid() {
		return errors.New("无效的target值")
	}

	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := Uint64(source)
		if err != nil {
			return err
		}
		target.SetUint(val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := Int64(source)
		if err != nil {
			return err
		}
		target.SetInt(val)
	case reflect.Float32, reflect.Float64:
		val, err := Float64(source)
		if err != nil {
			return err
		}
		target.SetFloat(val)
	case reflect.Bool:
		val, err := Bool(source)
		if err != nil {
			return err
		}
		target.SetBool(val)
	case reflect.String:
		val, err := String(source)
		if err != nil {
			return err
		}
		target.SetString(val)
	default:
		sourceValue := reflect.ValueOf(source)
		targetType := target.Type()
		if !sourceValue.Type().ConvertibleTo(targetType) {
			return fmt.Errorf("当前类型[%v]无法转换成目标类型:[%v]", sourceValue.Type(), targetType)
		}
		target.Set(sourceValue.Convert(targetType))
	}

	return nil
}
