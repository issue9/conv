// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package conv

import (
	"errors"
	"fmt"
	"reflect"
)

// 将source的值保存到成target中。
//
// 若类型不能直接转换，会尝试其它种方式转换，比如strconv.ParseInt()等。
func Value(source interface{}, target reflect.Value) error {
	kind := target.Kind()

	for kind == reflect.Ptr {
		target = target.Elem()
		kind = target.Kind()
	}

	if !target.CanSet() {
		return fmt.Errorf("conv.Value:无法改变target的值[%v]", target.Kind())
	}

	if !target.IsValid() {
		return errors.New("conv.Value:无效的target值")
	}

	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val, err := Uint64(source); err != nil {
			return err
		} else {
			target.SetUint(val)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val, err := Int64(source); err != nil {
			return err
		} else {
			target.SetInt(val)
		}
	case reflect.Float32, reflect.Float64:
		if val, err := Float64(source); err != nil {
			return err
		} else {
			target.SetFloat(val)
		}
	case reflect.Bool:
		if val, err := Bool(source); err != nil {
			return err
		} else {
			target.SetBool(val)
		}
	case reflect.String:
		if val, err := String(source); err != nil {
			return err
		} else {
			target.SetString(val)
		}
	default:
		sourceValue := reflect.ValueOf(source)
		targetType := target.Type()
		if !sourceValue.Type().ConvertibleTo(targetType) {
			return fmt.Errorf("conv.Value:当前类型[%v]无法转换成目标类型:[%v]", sourceValue.Type(), targetType)
		}
		target.Set(sourceValue.Convert(targetType))
	}

	return nil
}
