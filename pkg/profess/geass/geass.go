/*
----------------------------------------
@Create 2023/11/16
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geass 15:54
----------------------------------------
@Version 1.0 2023/11/16
@Memo create this file
*/

package geass

import (
	"github.com/lengpucheng/Geassgo/pkg/geasserr"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
)

// 执行容器
var geasses = make(map[string]Geass)

// Execute 执行那里
// ctx 执行的上下文对象
// name 需要执行的模块
// val 执行模块的参数
func Execute(ctx contract.Context, name string, val any) error {
	if ctx == nil {
		ctx = Background()
	}
	executor := GetGeass(name)
	if executor == nil {
		return geasserr.ModuleValueNotSupport.New(name)
	}
	// 整体渲染
	if executor.OverallRender() {
		str, err := RenderObject4YamlStr(ctx, val)
		if err != nil {
			return err
		}
		val = str
	}
	// 重载渲染
	if ok, instance := executor.OverloadRender(); ok {
		if err := TransObject4Yaml(instance, val); err != nil {
			return err
		}
		val = instance
	}

	return executor.Execute(ctx, val)
}

// GetGeass 获取执行器
func GetGeass(name string) Geass {
	return geasses[name]
}

// RegisterGeass 注册执行器
func RegisterGeass(name string, exec Geass) {
	geasses[name] = exec
}

// Geass 定义能力
type Geass interface {
	// Execute 执行Geass
	// ctx 执行的上下文对象
	// val 执行的参数 当OverloadRender == true时 val将为OverloadRender 返回的any类型
	Execute(ctx contract.Context, val any) error
	// OverallRender 整体覆盖渲染
	// 当该接口返回 true 将渲染val 中的全部 插值参数
	OverallRender() bool
	// OverloadRender 重载渲染类型
	// 当该接口返回 true 时 execute中执行的val将为返回的any同类型
	// any 应该为指针类型变量  否则将导致错误
	OverloadRender() (bool, any)
}
