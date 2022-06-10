// SPDX-License-Identifier: MIT

package conv

import (
	"errors"
	"fmt"
	"reflect"
)

// FieldConvert 字段转换
//
// 用于 map 转换到一个对象实例或是从一个对象实例转换到 map 时，字段名称的转换。
type FieldConvert func(src string) (dest string)

// FieldConvert 的默认实现
func defaultFieldConvert(src string) string {
	return src
}

// 将 obj 对象转换成 map[string]interface{} 格式的数据
func obj2Map(obj any, maps map[string]any, conv FieldConvert) error {
	objVal := reflect.ValueOf(obj)
	for objVal.Kind() == reflect.Ptr { // 如果是指针，则获取指向的对象
		objVal = objVal.Elem()
	}

	if objVal.Kind() != reflect.Struct {
		return typeError(obj, "map[string]interface{}")
	}

	objType := objVal.Type()
	num := objType.NumField()
	for i := 0; i < num; i++ {
		fieldType := objType.Field(i)
		fieldVal := objVal.Field(i)
		if !fieldVal.CanInterface() {
			continue
		}

		var err error
		switch {
		case fieldType.Anonymous: // 匿名字段
			err = obj2Map(fieldVal.Interface(), maps, conv)
		case fieldType.Type.Kind() == reflect.Ptr: // 如果是指针，就获取指针指向的元素
			fieldVal = fieldVal.Elem()
			fallthrough
		case fieldType.Type.Kind() == reflect.Struct: // 嵌套类型
			m := make(map[string]any)
			err = obj2Map(fieldVal.Interface(), m, conv)
			maps[conv(fieldType.Name)] = m
		default:
			maps[conv(fieldType.Name)] = fieldVal.Interface()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// Obj2Map 将 obj 转换成 map
//
// NOTE: 只能转换可导出的数据。
func Obj2Map(obj any, conv FieldConvert) (map[string]any, error) {
	ret := make(map[string]any)

	if conv == nil {
		conv = defaultFieldConvert
	}
	return ret, obj2Map(obj, ret, conv)
}

// Map2Obj 将 map 中的数据转换成一个结构中的数据
func Map2Obj(src any, dest any, conv FieldConvert) error {
	srcVal, destVal, conv, err := map2ObjCheck(src, dest, conv)
	if err != nil {
		return err
	}

	keys := srcVal.MapKeys()
	l := len(keys)
	for i := 0; i < l; i++ {
		k := keys[i]

		if k.Kind() != reflect.String {
			return errors.New("conv: src 必须为 map[string]interface{} 类型")
		}

		srcItemVal := srcVal.MapIndex(k)
		if !srcItemVal.CanInterface() {
			continue
		}

		fieldValue := destVal.FieldByName(conv(k.String()))
		if !fieldValue.CanSet() {
			continue
		}

		fieldType := fieldValue.Type()
		srcItemType := srcItemVal.Type()
		if fieldType.Kind() == srcItemType.Kind() { // 类型相同
			fieldValue.Set(srcItemVal)
			continue
		}

		// 如果 src 中元素的类型为 Interface，则再对该值使用 reflect.ValueOf
		// 就能正常地使用类型断言和其它判断了。
		if srcItemType.Kind() == reflect.Interface {
			srcItemVal = reflect.ValueOf(srcItemVal.Interface())
			srcItemType = srcItemVal.Type()
		}

		if !srcItemVal.CanInterface() {
			continue
		}

		if srcItemType.Kind() == reflect.Map { // 含有子元素
			err := Map2Obj(srcItemVal.Interface(), fieldValue.Interface(), conv)
			if err != nil {
				return err
			}
			continue
		}

		if srcItemType.ConvertibleTo(fieldType) { // 类型之间可转换
			fieldValue.Set(srcItemVal.Convert(fieldType))
		}
	}

	return nil
}

// 对 map2Obj 各个参数的检测，并返回正确的值或是错误信息。
func map2ObjCheck(src any, dest any, conv FieldConvert) (srcVal reflect.Value, destVal reflect.Value, fun FieldConvert, err error) {
	destVal = reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr {
		err = fmt.Errorf("conv: dest 必须为一个 struct 对象的指针，实际类型为[%v]", destVal.Type())
		return
	}

	destVal = destVal.Elem()
	if destVal.Kind() != reflect.Struct {
		err = fmt.Errorf("conv: dest 必须为一个 struct 对象的指针，实际类型为[%v]", destVal.Type())
		return
	}

	srcVal = reflect.ValueOf(src)

	if srcVal.Kind() == reflect.Ptr { // src有可能是个map指针，需要转换成map对象
		srcVal = srcVal.Elem()
	}

	if srcVal.Kind() != reflect.Map {
		err = fmt.Errorf("conv: src 必须为 map 类型或是 map 指针，实际类型为[%v]", srcVal.Type())
		return
	}

	if conv == nil {
		fun = defaultFieldConvert
	} else {
		fun = conv
	}
	return
}
