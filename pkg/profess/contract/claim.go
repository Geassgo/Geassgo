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

// Loop 循环探测
// 将该任务最多retries次 每次间隔 delay 秒
// 当 until的条件满足（返回true) 时 结束
type Loop struct {
	Until   string `json:"until" yaml:"until"`     // 条件 为true即为成功
	Retries int    `json:"retries" yaml:"retries"` // 重试次数
	Delay   int    `json:"delay" yaml:"delay"`     // 每次延时
}

type Claim struct {
	Mod         `json:",inline" yaml:",inline"`
	Task        `json:",inline" yaml:",inline"`
	When        string   `json:"when" yaml:"when"`                 // 条件判断
	Register    string   `json:"register" yaml:"register"`         // 变量注册
	IgnoreError bool     `json:"ignore_error" yaml:"ignore_error"` // 是否忽略错误
	WithItem    []string `json:"with_item" yaml:"with_item"`       // 迭代器
	Tasks       []Claim  `json:"tasks" yaml:"tasks"`               // 任务组
	Roles       []string `json:"roles" yaml:"roles"`               // 导入角色
	Include     string   `json:"include" yaml:"include"`           // 导入外部Claim
	// 循环执行的次数
	Loop *Loop `json:"loop" yaml:"loop"` // 循环探测
	// 标签过滤
	// 若 执行时 为传入 标签/跳过标签 则该配置项目不生效
	// 若 执行时 传入 tags 则将只会执行设置了tag 且 匹配的
	// 若 执行时 传入 skip-tags 则将仅跳过匹配的tag
	Tags []string `json:"tags" yaml:"tags"` // 标签组
	// 动作过滤
	// 和标签过滤类似当相反
	// 未设置Action的条目将总是被执行
	// 若设置Action的条目则必须匹配后才能执行
	Action string `json:"action" yaml:"action"` // 动作
}

// IsWhen 渲染并判断是否成立
func (t *Claim) IsWhen(variable *Variable) bool {
	if t.When == "" {
		return true
	}
	s, _ := coderender.Template(t.When, "{{", "}}", coderender.FuncMap(), variable.ToMap())
	return s == "true"
}

// IsSelect 是否应该被选择执行
// 当tag和action都被选中时即为选中
func (t *Claim) IsSelect(selector Selector) bool {
	return t.IsTag(selector) && t.IsAction(selector)
}

// IsTag 是否被tag筛选
// 返回true 表示应该被执行
// 当 selector 的 tag   和 skip-tag为空时候 总是被选择
// 当 selector 的 tag  不为空时 仅当tag匹配时的会被执行
// 当 selector 的 skip 不会空时 仅当匹配的skip的tag不会被执行
// 当 tag 和 skip 同时存在时候 优先级 skip > tag
func (t *Claim) IsTag(selector Selector) bool {
	// 转换格式
	var claimTags = coderender.Slice2Any(t.Tags)
	// 当 tags 为空 取决于是否被跳过
	// 没有tag 或者 未被匹配 则被选中
	if len(selector.GetTags()) < 1 {
		return len(claimTags) < 1 || !coderender.ArrayIn(coderender.Slice2Any(selector.GetSkipTags()), false, claimTags...)
	}

	// 当 tag 不为空 tag必须存在
	if len(claimTags) < 1 {
		return false
	}

	// 当 tag 不为空且 skip 为空 取决于 是否需要匹配以及是否被匹配
	if len(selector.GetSkipTags()) < 1 {
		return coderender.ArrayIn(coderender.Slice2Any(selector.GetTags()), false, claimTags...)
	}

	// 当都不为空时
	// 必须 tag 匹配 且 不被跳过
	return coderender.ArrayIn(coderender.Slice2Any(selector.GetTags()), false, claimTags...) &&
		!coderender.ArrayIn(coderender.Slice2Any(selector.GetTags()), false, claimTags...)
}

// IsAction 是否应该被执行动作
// 若当前执行的动作匹配则返回true
// 特别的 当未定义 claim的动作时候 将返回 true
func (t *Claim) IsAction(selector Selector) bool {
	return t.Action == "" || t.Action == selector.GetAction()
}
