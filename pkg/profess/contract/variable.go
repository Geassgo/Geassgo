/*
----------------------------------------
@Create 2023-11-08
@Author 冷朴承<lengpucheng@qq.com>
@Program geassgo
@Describe variable
----------------------------------------
@Version 1.0 2023/11/8-22:01
@Memo create this file
*/

package contract

type Variable struct {
	System   map[string]any `json:"system"`   // 系统变量
	Values   map[string]any `json:"values"`   // 普通变量
	Register map[string]any `json:"register"` // 注册变量
}

func (v *Variable) Check() {
	if v.System == nil {
		v.System = make(map[string]any)
	}
	if v.Register == nil {
		v.Register = make(map[string]any)
	}
	if v.Values == nil {
		v.Values = make(map[string]any)
	}
}

func (v *Variable) ToMap() map[string]any {
	return map[string]any{
		"system":   v.System,
		"values":   v.Values,
		"register": v.Register,
	}
}
