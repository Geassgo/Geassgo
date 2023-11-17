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

var geasses = make(map[string]Geass)

// Execute 执行模块
// val 不可以是指针
func Execute(ctx contract.Context, name string, val any) error {
	if ctx == nil {
		ctx = contract.DefaultContext()
	}
	executor := GetGeass(name)
	if executor == nil {
		return geasserr.ModuleValueNotSupport.New(name)
	}
	// 整体渲染
	if executor.OverallRender() {
		if err := RenderObject4Yaml(ctx, &val); err != nil {
			return err
		}
	}
	// 重载渲染
	if ok, instance := executor.OverloadRender(); ok {
		if err := TransObject4Json(instance, val); err != nil {
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

// 注册执行器
func registerGeass(name string, exec Geass) {
	geasses[name] = exec
}

// Geass 定义执行器
type Geass interface {
	Execute(ctx contract.Context, val any) error
	OverallRender() bool
	OverloadRender() (bool, any)
}
