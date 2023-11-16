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
	"encoding/json"
	"github.com/lengpucheng/Geassgo/pkg/coderender"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"gopkg.in/yaml.v3"
)

// RenderStr 常规渲染
// ctx 上下文变量
// str 带渲染文本
func RenderStr(ctx contract.Context, str string) (string, error) {
	if ctx.GetItem() == nil {
		return coderender.Template(str, "{{", "}}", coderender.DefaultTemplateFunc(), ctx.GetVariable().ToMap())
	}
	return coderender.Template(str, "{{", "}}", coderender.DefaultTemplateFunc(), ctx.GetVariable().ToMap(), map[string]any{"item": ctx.GetItem(), "itemIndex": ctx.GetItemIndex()})
}

// RenderObject4Yaml template渲染对象使用 使用yaml作为中间媒介
// ctx 上下文变量
// obj 对象
func RenderObject4Yaml(ctx contract.Context, obj any) (string, error) {
	yml, err := yaml.Marshal(obj)
	if err != nil {
		return "", err
	}
	return RenderStr(ctx, string(yml))
}

// TransObject4Json 转换对象 使用json方案
// target 转换后的目标对象
// obj 转换前的对象
func TransObject4Json(target, object any) error {
	marshal, err := json.Marshal(object)
	if err != nil {
		return err
	}
	return json.Unmarshal(marshal, target)
}
