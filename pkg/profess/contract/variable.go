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

// Variable 变量值
type Variable struct {
	System   *System        `json:"system"`   // 系统变量
	Values   map[string]any `json:"values"`   // 普通变量
	Register map[string]any `json:"register"` // 注册变量
}

// Check 检查variable 存在为nil的补偿为非空空内容
func (v *Variable) Check() *Variable {
	if v.System == nil {
		v.System = &System{}
	}
	if v.Register == nil {
		v.Register = make(map[string]any)
	}
	if v.Values == nil {
		v.Values = make(map[string]any)
	}
	return v
}

// ToMap 转换为map类型
func (v *Variable) ToMap() map[string]any {
	return map[string]any{
		"System":   v.System,
		"Values":   v.Values,
		"Register": v.Register,
		"system":   v.System,
		"values":   v.Values,
		"register": v.Register,
	}
}

// DeepCopy 深拷贝
func (v *Variable) DeepCopy() *Variable {
	return &Variable{
		System:   v.System,
		Values:   deepCopyMap(v.Values),
		Register: deepCopyMap(v.Register),
	}
}

// 拷贝map
func deepCopyMap(src map[string]any) map[string]any {
	if src == nil {
		return nil
	}
	dest := make(map[string]any)
	for k, v := range src {
		dest[k] = v
	}
	return dest
}
