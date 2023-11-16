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
	"errors"
	"fmt"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
)

var executors = make(map[string]Geass)

// Execute 执行模块
func Execute(ctx contract.Context, name string, module any) error {
	executor := GetGeass(name)
	if !executor.OverallRender() {
		return executor.Execute(ctx, module)
	}
	if executor == nil {
		return errors.New(fmt.Sprintf("the module %s is not found!", name))
	}
	str, err := RenderObject4Yaml(ctx, module)
	if err != nil {
		return err
	}
	return executor.Execute(ctx, str)
}

// GetGeass 获取执行器
func GetGeass(name string) Geass {
	return executors[name]
}

// 注册执行器
func registerGeass(name string, exec Geass) {
	executors[name] = exec
}

// Geass 定义执行器
type Geass interface {
	Execute(ctx contract.Context, val any) error
	OverallRender() bool
}
