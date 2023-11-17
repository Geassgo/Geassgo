/*
----------------------------------------
@Create 2023/11/16
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe 契约清单
----------------------------------------
@Version 1.0 2023/11/16
@Memo create this file
*/

package contract

import "github.com/lengpucheng/Geassgo/pkg/coderender"

type Claim struct {
	Mod         `json:",inline" yaml:",inline"`
	Task        `json:",inline" yaml:",inline"`
	When        string   `json:"when" yaml:"when"`                 // 条件判断
	Register    string   `json:"register" yaml:"register"`         // 变量注册
	IgnoreError bool     `json:"ignore_error" yaml:"ignore_error"` // 是否忽略错误
	WithItem    []string `json:"with_item" yaml:"with_item"`       // 迭代器
	Tasks       []Claim  `json:"tasks" yaml:"tasks"`               // 任务组
	Tags        []string `json:"tags" yaml:"tags"`                 // 标签组
	Roles       []string `json:"roles" yaml:"roles"`               // 导入角色
	Include     string   `json:"include" yaml:"include"`           // 导入外部Claim
}

// IsWhen 渲染并判断是否成立
func (t *Claim) IsWhen(variable *Variable) bool {
	if t.When == "" {
		return true
	}
	s, _ := coderender.Template(t.When, "{{", "}}", coderender.DefaultTemplateFunc(), variable.ToMap())
	return s == "true"
}
