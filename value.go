// SPDX-License-Identifier: MIT

package conv

import (
	"errors"
	"fmt"
	"reflect"
)

// Value 将 source 的值保存到成 target 中
//
// 如果 source 为 nil，则会将 target 的值设置为其默认的零值。
//
// 若类型不能直接转换，会尝试其它种方式转换，比如 strconv.ParseInt() 等。
func Value(source any, target reflect.Value) error {
	kind := target.Kind()

	for kind == reflect.Ptr {
		target = target.Elem()
		kind = target.Kind()
	}

	if !target.CanSet() {
		return fmt.Errorf("conv: 无法改变 target %s 的值", target.Kind())
	}

	if !target.IsValid() {
		return errors.New("conv: 无效的 target 值")
	}

	if source == nil {
		target.Set(reflect.Zero(target.Type()))
		return nil
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
	case reflect.Slice:
		s := reflect.ValueOf(source)
		sk := s.Kind()
		tek := target.Type().Elem().Kind()

		// []byte []rune 和 string 之间可以互相转换
		if sk == reflect.String && (tek == reflect.Uint8 || tek == reflect.Int32) {
			return valueDefault(source, target)
		}

		if sk != reflect.Array && sk != reflect.Slice {
			return typeError(source, "slice")
		}

		l := s.Len()
		tmp := reflect.MakeSlice(target.Type(), l, l)
		for i := 0; i < l; i++ {
			si := s.Index(i).Interface()
			if err := Value(si, tmp.Index(i)); err != nil {
				return err
			}
		}
		target.Set(tmp)
	case reflect.Array:
		s := reflect.ValueOf(source)
		sk := s.Kind()
		tek := target.Type().Elem().Kind()

		// []byte []rune 和 string 之间可以互相转换
		if sk == reflect.String && (tek == reflect.Uint8 || tek == reflect.Int32) {
			return valueDefault(source, target)
		}

		if sk != reflect.Array && sk != reflect.Slice {
			return typeError(source, "array")
		}

		l := s.Len()
		if l != target.Len() {
			return fmt.Errorf("conv: 两者长度不一样，无法转换 %d: %d", l, target.Len())
		}

		for i := 0; i < l; i++ {
			si := s.Index(i).Interface()
			if err := Value(si, target.Index(i)); err != nil {
				return err
			}
		}
	default:
		return valueDefault(source, target)
	}

	return nil
}

func valueDefault(source any, target reflect.Value) error {
	sourceValue := reflect.ValueOf(source)
	targetType := target.Type()
	if !sourceValue.Type().ConvertibleTo(targetType) {
		return typeError(source, targetType.String())
	}
	target.Set(sourceValue.Convert(targetType))

	return nil
}
