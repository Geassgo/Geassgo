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

type Claim struct {
	Task        `json:",inline" yaml:",inline"`
	When        string   // 条件判断
	Register    string   // 变量注册
	IgnoreError bool     // 是否忽略错误
	WithItem    []string // 迭代器
	Tasks       []Task   // 任务组
	Tags        []string // 标签组
	Roles       []string // 导入角色
	Include     string   // 导入外部Claim
}
