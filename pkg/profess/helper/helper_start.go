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
	"errors"
	"github.com/google/uuid"
	"github.com/lengpucheng/Geassgo/pkg/coderender"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/geass"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

// RunClaim 启动一个task任务 使用valuesPath的变量 values可以为空
// ctx 上下文对象
// claimPath claim文件路径
// valuePath values文件路径
// coverValues values文件路径
// coverValues 覆盖变量 map[string]any
func RunClaim(ctx context.Context, claimPath, valuePath string, coverValues ...map[string]any) (contract.Context, error) {
	values := loadValuesFromFile(valuePath, coverValues...)
	tCtx := NewContext(ctx, geass.DefaultRuntime(), values)
	return tCtx, LoadAndExecute4File(tCtx, claimPath)
}

// RunChart 执行Chart包
// ctx 上下文对象
// chartPath chart包路径路径
// valuePath values文件路径
// coverValues 覆盖变量 map[string]any
func RunChart(ctx context.Context, chartPath string, coverValues ...map[string]any) (contract.Context, error) {
	dir := getDefaultPath("chart")
	id := uuid.New().String()
	dir = filepath.Join(dir, id)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	if err := coderender.UnArchive(chartPath, dir); err != nil {
		return nil, err
	}
	// 获取入口
	readDir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, d := range readDir {
		if d.IsDir() {
			slog.Info("Execute Charts of", "name", d.Name())
			return runChartDir(ctx, filepath.Join(dir, d.Name()), coverValues...)
		}
	}
	return nil, nil
}

// RunChart 执行Chart包
// ctx 上下文对象
// chartPath chart包路径路径
// valuePath values文件路径
// coverValues 覆盖变量 map[string]any
func runChartDir(ctx context.Context, chartPath string, coverValues ...map[string]any) (contract.Context, error) {
	if !coderender.IsNotExist(filepath.Join(chartPath, "Chart.yaml")) {
		// TODO 如果Chart.yaml 文件存在
	}
	if coderender.IsNotExist(filepath.Join(chartPath, "values.yaml")) {
		return nil, errors.New("the chart is illegal,not have values.yaml file")
	}
	if !coderender.IsNotExist(filepath.Join(chartPath, "main.yaml")) {
		return RunClaim(ctx, filepath.Join(chartPath, "main.yaml"), filepath.Join(chartPath, "values.yaml"))
	} else if coderender.IsNotExist(filepath.Join(chartPath, "tasks/main.yaml")) {
		return nil, errors.New("the chart is illegal,not have main.yaml or tasks/main.yaml")
	}
	// 没有roles直接平铺的情况
	variable := loadValuesFromFile(filepath.Join(chartPath, "values.yaml"), coverValues...)
	rolePath := chartPath + "/"
	tCtx := NewContext(ctx, geass.NewRuntime(rolePath, rolePath, -1, nil), variable)
	return tCtx, geass.Execute(tCtx, Roles, nil)
}

// 从values文件夹中加载values变量
// valuesPath 加载变量路径
// coverValues 覆盖变量,优先级将高于加载的变量,当存在时仅会取第一个
func loadValuesFromFile(valuePath string, coverValues ...map[string]any) *contract.Variable {
	var values = &contract.Variable{Values: make(map[string]any)}
	if valuePath != "" {
		valuesFile, err := os.ReadFile(valuePath)
		if err != nil {
			if err = yaml.Unmarshal(valuesFile, values.Values); err != nil {
				panic(err)
			}
		}
	}
	// 使用覆盖变量进行覆盖
	if len(coverValues) > 0 && coverValues[0] != nil {
		for k, v := range coverValues[0] {
			values.Values[k] = v
		}
	}
	return values
}

// 获取默认的存储地址
func getDefaultPath(subPath ...string) string {
	switch runtime.GOOS {
	case "windows":
		subPath = append([]string{"C:\\ProgramData\\geassgo\\"}, subPath...)
	case "linux":
		fallthrough
	default:
		subPath = append([]string{"/etc/geassgo/"}, subPath...)
	}
	return filepath.Join(subPath...)
}
