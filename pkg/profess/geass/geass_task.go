/*
----------------------------------------
@Create 2023/11/17
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geass_task 9:48
----------------------------------------
@Version 1.0 2023/11/17
@Memo create this file
*/

package geass

import (
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
)

func init() {
	registerGeass(Task, &geassTask{})
}

const Task = "task"

type geassTask struct{}

func (g *geassTask) Execute(ctx contract.Context, val any) error {
	task := val.(contract.Task)
	// mod 执行
	if task.Mod != nil {
		for k, v := range task.Mod {
			return Execute(ctx, k, v)
		}
	}
	return Execute(ctx, Shell, task.Shell)
}

func (g *geassTask) OverallRender() bool {
	return false
}

func (g *geassTask) OverloadRender() (bool, any) {
	return false, nil
}
