/*
----------------------------------------
@Create 2023/11/15
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe 模板渲染
----------------------------------------
@Version 1.0 2023/11/15
@Memo create this file
*/

package coderender

import (
	"bytes"
	"text/template"
)

// Template 渲染
// str 待渲染文本
// left right 渲染符号
// funcMap 渲染函数列表
// params 渲染参数列表
// defs 覆盖参数列表,优先级 defs > params
func Template(str, left, right string, funcMap map[string]any, params map[string]any, defs ...map[string]any) (string, error) {
	data, err := TemplateBytes([]byte(str), left, right, funcMap, params, defs...)
	return string(data), err
}

// TemplateBytes 字节渲染
// str 待渲染文本
// left right 渲染符号
// funcMap 渲染函数列表
// params 渲染参数列表
// defs 覆盖参数列表,优先级 defs > params
func TemplateBytes(data []byte, left, right string, funcMap map[string]any, params map[string]any, defs ...map[string]any) ([]byte, error) {
	if funcMap == nil {
		funcMap = make(map[string]any)
	}
	if params == nil {
		params = make(map[string]any)
	}

	// 合并参数
	for i := range defs {
		if defs[i] == nil {
			continue
		}
		for k, v := range defs[i] {
			params[k] = v
		}
	}

	// 渲染
	parse, err := template.New("").Delims(left, right).Funcs(funcMap).Parse(string(data))
	if err != nil {
		return nil, err
	}
	buff := new(bytes.Buffer)
	//buff.Bytes()
	e := parse.Execute(buff, params)
	return buff.Bytes(), e
}
