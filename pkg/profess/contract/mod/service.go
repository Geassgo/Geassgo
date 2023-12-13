/*
----------------------------------------
@Create 2023/12/12-14:26
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe service
----------------------------------------
@Version 1.0 2023/12/12
@Memo create this file
*/

package mod

type Service struct {
	Name   string       `json:"name" yaml:"name"`     // 服务名称 linux下后缀可以带上.service
	State  ServiceState `json:"state" yaml:"state"`   // 期望状态
	Enable *bool        `json:"enable" yaml:"enable"` // 开机自启
	Reload bool         `json:"reload" yaml:"reload"` // 执行操作前都reload
}

type ServiceState string

const (
	ServiceRestart ServiceState = "restart" // 重启
	ServiceStart   ServiceState = "start"   // 启动
	ServiceStop    ServiceState = "stop"    // 停止
)
