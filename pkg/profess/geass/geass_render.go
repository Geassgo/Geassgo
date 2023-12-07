/*
----------------------------------------
@Create 2023/11/16
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geasser_render 15:55
----------------------------------------
@Version 1.0 2023/11/16
@Memo create this file
*/

package geass

import (
	"github.com/lengpucheng/Geassgo/pkg/coderender"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"gopkg.in/yaml.v3"
)

// RenderStr 常规渲染
// ctx 上下文变量
// str 带渲染文本
func RenderStr(ctx contract.Context, str string) (string, error) {
	if ctx.GetItem() == nil {
		return coderender.Template(str, "{{", "}}", coderender.FuncMap(), ctx.GetVariable().ToMap())
	}
	return coderender.Template(str, "{{", "}}", coderender.FuncMap(), ctx.GetVariable().ToMap(), map[string]any{"item": ctx.GetItem(), "itemIndex": ctx.GetItemIndex()})
}

// RenderObject4Yaml template渲染对象使用 使用yaml作为中间媒介
// ctx 上下文变量
// obj 对象 渲染后将最终保存在此
func RenderObject4Yaml(ctx contract.Context, obj any) error {
	str, err := RenderObject4YamlStr(ctx, obj)
	if err != nil {
		return err
	}
	return yaml.Unmarshal([]byte(str), obj)
}

// RenderObject4YamlStr template渲染对象使用 使用yaml作为中间媒介
// ctx 上下文变量
// obj 对象
func RenderObject4YamlStr(ctx contract.Context, obj any) (string, error) {
	yml, err := yaml.Marshal(obj)
	if err != nil {
		return "", err
	}
	return RenderStr(ctx, string(yml))
}

// TransObject4Yaml 转换对象 使用yaml方案
// target 转换后的目标对象
// obj 转换前的对象
func TransObject4Yaml(target, object any) error {
	var marshal []byte
	if str, ok := object.(string); ok {
		marshal = []byte(str)
	} else if data, ok := object.([]byte); ok {
		marshal = data
	} else {
		data, err := yaml.Marshal(object)
		if err != nil {
			return err
		}
		marshal = data
	}
	return yaml.Unmarshal(marshal, target)
}
