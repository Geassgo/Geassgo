/*
----------------------------------------
@Create 2023/12/6-15:19
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe helper_start
----------------------------------------
@Version 1.0 2023/12/6
@Memo create this file
*/

package helper

import (
	"context"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/geass"
	"gopkg.in/yaml.v3"
	"os"
)

// RunTask 启动一个task任务 使用valuesPath的变量 values可以为空
func RunTask(ctx context.Context, taskPath, valuePath string) (contract.Context, error) {
	var values = &contract.Variable{Values: make(map[string]any)}
	if valuePath != "" {
		valuesFile, err := os.ReadFile(valuePath)
		if err != nil {
			if err = yaml.Unmarshal(valuesFile, values); err != nil {
				panic(err)
			}
		}
	}
	tCtx := NewContext(ctx, geass.DefaultRuntime(), values)
	return tCtx, LoadAndExecute4File(tCtx, taskPath)
}

// RunChart 执行Chart包
func RunChart(ctx context.Context, chartPath string) (contract.Context, error) {
	return nil, nil
}
