/*
----------------------------------------
@Create 2023/11/16
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe impl_taskhelper 15:46
----------------------------------------
@Version 1.0 2023/11/16
@Memo create this file
*/

package helper

import (
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/geass"
	"runtime"
)

type taskHelper struct {
	*contract.Task
}

func (t *taskHelper) GetContext() *contract.Context {
	//TODO implement me
	panic("implement me")
}

func (t *taskHelper) GetTask() *contract.Task {
	//TODO implement me
	panic("implement me")
}

func (t *taskHelper) GetLog() *string {
	//TODO implement me
	panic("implement me")
}

func (t *taskHelper) Execute(ctx contract.Context) error {
	// mod 执行
	if t.Mod != nil {
		for k, v := range t.Mod {
			return geass.Execute(ctx, k, v)
		}
	}
	// shell 执行
	switch runtime.GOOS {
	case "linux":
		return geass.Execute(ctx, "shell", t.Shell)
	case "windows":
		return geass.Execute(ctx, "shell", t.Shell)
	}
	return nil
}
