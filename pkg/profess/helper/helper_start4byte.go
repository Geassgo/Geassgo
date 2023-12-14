/*
----------------------------------------
@Create 2023/12/11-11:23
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe helper_start_byte
----------------------------------------
@Version 1.0 2023/12/11
@Memo create this file
*/

package helper

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"os"
)

func init() {
	_ = os.MkdirAll(getDefaultPath("temp"), 0755)
}

// RunChart4data 启动一个chart任务 使用valuesPath的变量 values可以为空
// ctx 上下文对象
// chart chart包
// values 覆盖变量
func RunChart4data(ctx context.Context, selector contract.Selector, chart []byte, values map[string]any) (contract.Context, error) {
	path := getDefaultPath("temp", fmt.Sprintf("%s-chart.tgz", uuid.New().String()))
	defer os.Remove(path)
	if err := os.WriteFile(path, chart, 0755); err != nil {
		return nil, err
	}
	return RunChart(ctx, selector, path, values)
}

// RunClaim4data 启动一个claim任务 使用valuesPath的变量 values可以为空
// ctx 上下文对象
// data claim内容（yaml格式）
// values 覆盖变量
func RunClaim4data(ctx context.Context, selector contract.Selector, data []byte, values map[string]any) (contract.Context, error) {
	path := getDefaultPath("temp", fmt.Sprintf("%s-main.yaml", uuid.New().String()))
	defer os.Remove(path)
	if err := os.WriteFile(path, data, 0755); err != nil {
		return nil, err
	}
	return RunClaim(ctx, selector, path, "", values)
}
