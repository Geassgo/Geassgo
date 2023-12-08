/*
----------------------------------------
@Create 2023/12/6-9:36
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe template
----------------------------------------
@Version 1.0 2023/12/6
@Memo create this file
*/

package template

type Template struct {
	Metadata     Metadata            `json:"metadata" yaml:"metadata"`         // 元数据
	Env          Env                 `json:"env" yaml:"env"`                   // 环境变量
	Workflow     Workflow            `json:"workflow" yaml:"workflow"`         // 工作流
	Variables    []Variable          `json:"variable" yaml:"variable"`         // 全局变量
	Applications []Application       `json:"applications" yaml:"applications"` // 应用
	Actions      []Application       `json:"action" yaml:"action"`             // 动作组
	HostSelect   map[string][]string `json:"hostSelect" yaml:"hostSelect"`     // 节点筛选
	HostReflect  map[string]string   `json:"hostReflect" yaml:"hostReflect"`   // 节点映射
}

type Metadata struct {
	Name        string `json:"name" yaml:"name"`               // 模板名称
	Type        string `json:"type" yaml:"type"`               // 模板类型
	Version     string `json:"version" yaml:"version"`         // 模板版本
	Description string `json:"Description" yaml:"Description"` // 模板描述
	CreateTime  string `json:"createTime" yaml:"createTime"`   // 创建时间
}

type Workflow struct {
	Preview []string       `json:"preview" yaml:"preview"` // 预先部署
	Before  []string       `json:"before" yaml:"before"`   // 前置部署
	After   []string       `json:"after" yaml:"after"`     // 后置部署
	Action  ActionWorkflow `json:"action" yaml:"action"`   // 动作定义
	Timeout int            `json:"timeout" yaml:"timeout"` // 超时时间
	Serial  bool           `json:"serial" yaml:"serial"`   // 是否串行
}

type ActionWorkflow struct {
	Install   []string `json:"install" yaml:"install"`
	Uninstall []string `json:"uninstall" yaml:"uninstall"`
	Upgrade   []string `json:"upgrade" yaml:"upgrade"`
	Rollback  []string `json:"rollback" yaml:"rollback"`
}

type Env map[string]any

type Variable struct {
	Name        string `json:"name" yaml:"name"`               // 变量名称
	Type        string `json:"type" yaml:"type"`               // 变量类型
	Value       any    `json:"value" yaml:"value"`             // 变量内容
	Required    bool   `json:"required" yaml:"required"`       // 必须
	Editable    bool   `json:"editable" yaml:"editable"`       // 可编辑
	Description string `json:"description" yaml:"description"` // 描述
}
