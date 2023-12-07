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

// RunTask 启动一个task任务 使用valuesPath的变量 values可以为空
func RunTask(ctx context.Context, taskPath, valuePath string) (contract.Context, error) {
	values := loadValuesFromFile(valuePath)
	tCtx := NewContext(ctx, geass.DefaultRuntime(), values)
	return tCtx, LoadAndExecute4File(tCtx, taskPath)
}

// RunChart 执行Chart包
func RunChart(ctx context.Context, chartPath string) (contract.Context, error) {
	var dir = "/etc/geassgo/chart"
	if runtime.GOOS == "windows" {
		dir = "C:\\ProgramData\\geassgo\\chart"
	}

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
			return runChartDir(ctx, filepath.Join(dir, d.Name()))
		}
	}
	return nil, nil
}

func runChartDir(ctx context.Context, chartDir string) (contract.Context, error) {
	if !coderender.IsNotExist(filepath.Join(chartDir, "Chart.yaml")) {
		// TODO 如果Chart.yaml 文件存在
	}
	if coderender.IsNotExist(filepath.Join(chartDir, "values.yaml")) {
		return nil, errors.New("the chart is illegal,not have values.yaml file")
	}
	if !coderender.IsNotExist(filepath.Join(chartDir, "main.yaml")) {
		return RunTask(ctx, filepath.Join(chartDir, "main.yaml"), filepath.Join(chartDir, "values.yaml"))
	} else if coderender.IsNotExist(filepath.Join(chartDir, "tasks/main.yaml")) {
		return nil, errors.New("the chart is illegal,not have main.yaml or tasks/main.yaml")
	}
	// 没有roles直接平铺的情况
	variable := loadValuesFromFile(filepath.Join(chartDir, "values.yaml"))
	rolePath := chartDir + "/"
	tCtx := NewContext(ctx, geass.NewRuntime(rolePath, rolePath, -1, nil), variable)
	return tCtx, geass.Execute(tCtx, Roles, nil)
}

// 从values文件夹中加载变量
func loadValuesFromFile(valuePath string) *contract.Variable {
	var values = &contract.Variable{Values: make(map[string]any)}
	if valuePath != "" {
		valuesFile, err := os.ReadFile(valuePath)
		if err != nil {
			if err = yaml.Unmarshal(valuesFile, values); err != nil {
				panic(err)
			}
		}
	}
	return values
}
