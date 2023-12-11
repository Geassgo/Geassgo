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
	"github.com/google/uuid"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"os"
)

func RunChart4data(ctx context.Context, chart []byte, values map[string]any) (contract.Context, error) {
	path := getDefaultPath("temp", uuid.New().String(), "chart.tgz")
	if err := os.WriteFile(path, chart, 0755); err != nil {
		return nil, err
	}
	return RunChart(ctx, path, values)
}

func RunClaim4data(ctx context.Context, taskData []byte, values map[string]any) (contract.Context, error) {
	path := getDefaultPath("temp", uuid.New().String(), "main.yaml")
	if err := os.WriteFile(path, taskData, 0755); err != nil {
		return nil, err
	}
	return RunClaim(ctx, path, "", values)
}
