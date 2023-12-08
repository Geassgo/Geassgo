/*
----------------------------------------
@Create 2023/12/6-9:52
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe application
----------------------------------------
@Version 1.0 2023/12/6
@Memo create this file
*/

package template

type Application struct {
	Condition   `json:",inline" yaml:",inline"`
	Name        string        `json:"name" yaml:"name"`               // 应用名称
	Description string        `json:"description" yaml:"description"` // 描述
	Variables   []Variable    `json:"variables" yaml:"variables"`     // 组变量
	Single      bool          `json:"single" yaml:"single"`           // 单实例
	Type        string        `json:"type" yaml:"type"`               // 应用类型
	Properties  Properties    `json:"properties" yaml:"properties"`   // 应用属性 与 components 互斥
	Components  []Application `json:"components" yaml:"components"`   // 部署组件 与 properties 互斥 当且仅当 type 为nil 或 “” 时
}

type Condition struct {
	Tags      []string `json:"tags" yaml:"tags"`           // 标签
	Condition []any    `json:"condition" yaml:"condition"` // 条件筛选
}

type Properties any
