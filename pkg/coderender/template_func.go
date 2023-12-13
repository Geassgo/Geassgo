/*
----------------------------------------
@Create 2023/11/15
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe 渲染函数
----------------------------------------
@Version 1.0 2023/11/15
@Memo create this file
*/

package coderender

import (
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"reflect"
	"strings"
	"text/template"
)

// FuncMap 核心的模板函数 包含 sprig和default 同helm一致
func FuncMap() template.FuncMap {
	// 获取 sprig 的 文本渲染函数
	funcMap := sprig.TxtFuncMap()

	// 添加默认的内置函数
	templateFunc := DefaultTemplateFunc()

	// merge func
	for k, v := range templateFunc {
		funcMap[k] = v
	}
	return funcMap
}

// DefaultTemplateFunc 默认的模板语法 自定义
// 使用时应该使用 FuncMap
func DefaultTemplateFunc() template.FuncMap {
	return map[string]any{
		"str":           Str,
		"belong":        Belong,
		"contain":       Contain,
		"exist":         Exist,
		"subset":        Subset,
		"eqs":           Equals,
		"ands":          Ands,
		"ors":           Ors,
		"obj":           Object,
		"json":          Object,
		"toToml":        toTOML,
		"toYaml":        toYAML,
		"fromYaml":      fromYAML,
		"fromYamlArray": fromYAMLArray,
		"toJson":        toJSON,
		"fromJson":      fromJSON,
		"fromJsonArray": fromJSONArray,
	}
}

// Str 强制转换为带双引号字符串(在外围加上双引号)
func Str(key interface{}) string {
	var str string
	switch key.(type) {
	case string:
		str = key.(string)
	default:
		str = fmt.Sprintf("%v", key)
	}
	return fmt.Sprintf(`"%s"`, str)
}

// Ors ===> if t1||t2||t4.... t可以为bool或string
func Ors(ts ...interface{}) bool {
	for _, t := range ts {
		if tv := fmt.Sprintf("%v", t); strings.ToLower(tv) == "true" {
			return true
		}
	}
	return false
}

// Ands ===> if t1&t2&t3&t4.... t可以为bool或string
func Ands(ts ...interface{}) bool {
	for _, t := range ts {
		if tv := fmt.Sprintf("%v", t); strings.ToLower(tv) == "false" {
			return false
		}
	}
	return ts != nil && len(ts) > 0
}

// Belong 属于
// str为ts 中的元素
func Belong(str interface{}, ts ...interface{}) bool {
	return Subset([]interface{}{str}, ts...)
}

// Contain 包含
// 当ts为 arr 的子集时候即为true
func Contain(arr []interface{}, ts ...interface{}) bool {
	return ArrayIn(arr, false, ts...)
}

// Subset 属于
// 当arr的全部元素都属于ts时候为true （子集）
func Subset(arr []interface{}, ts ...interface{}) bool {
	if arr == nil || ts == nil {
		return ts == nil
	}
	set := make(map[interface{}]struct{}, len(arr))
	for _, key := range ts {
		set[fmt.Sprintf("%v", key)] = struct{}{}
	}
	for _, key := range arr {
		if _, ok := set[fmt.Sprintf("%v", key)]; !ok {
			return false
		}
	}
	return true
}

// Exist 存在
// 当arr中存在一个元素属于ts时候返回true (交集)
func Exist(arr []interface{}, ts ...interface{}) bool {
	return ArrayIn(arr, true, ts...)
}

// ArrayIn 判断元素是否在数组中, belong 表示是否是属于(只有一个满足即可),false表示包含
// arr或ts为nil时候 返回 ts==nil (空集包含空集,空集是然后集合的子集)
func ArrayIn(arr []interface{}, belong bool, ts ...interface{}) bool {
	if arr == nil || ts == nil {
		return ts == nil
	}
	set := make(map[interface{}]struct{}, len(arr))
	for _, key := range arr {
		set[fmt.Sprintf("%v", key)] = struct{}{}
	}
	for _, t := range ts {
		_, ok := set[fmt.Sprintf("%v", t)]
		if belong && ok {
			return true
		} else if !(ok || belong) {
			return false
		}
	}
	return false
}

// Equals 判断两个是否相同
func Equals(a, b interface{}) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

// Object 转换为json字符串显示(不带引号)
func Object(obj interface{}) string {
	var res string
	vt := reflect.ValueOf(obj)
	kt := reflect.TypeOf(obj)
	num := 0
	switch kt.Kind() {
	case reflect.Slice:
		res += "["
		for i := 0; i < vt.Len(); i++ {
			res += fmt.Sprintf("%s,", Object(vt.Index(i).Interface()))
			num++
		}
		if num > 0 {
			res = res[:len(res)-1]
		}
		res += "]"
	case reflect.Map:
		res += "{"
		for _, k := range vt.MapKeys() {
			res += fmt.Sprintf(`"%v": %s,`, k.Interface(), Object(vt.MapIndex(k).Interface()))
			num++
		}
		if num > 0 {
			res = res[:len(res)-1]
		}
		res += "}"
	case reflect.Struct:
		res += "{"
		for i := 0; i < vt.NumField(); i++ {
			var name string
			if name = kt.Field(i).Tag.Get("json"); name != "" {
			} else if name = kt.Field(i).Tag.Get("yaml"); name != "" {
			} else if name = kt.Field(i).Name; name != "" {
			} else {
				continue
			}
			res += fmt.Sprintf(`"%v": %s,`, name, Object(vt.Field(i).Interface()))
			num++
		}
		if num > 0 {
			res = res[:len(res)-1]
		}
		res += "}"
	default:
		res = fmt.Sprintf(`"%v"`, obj)
	}
	return res
}
