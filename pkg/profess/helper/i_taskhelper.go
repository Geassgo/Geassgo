/*
----------------------------------------
@Create 2023/11/16
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe i_taskhelper 15:43
----------------------------------------
@Version 1.0 2023/11/16
@Memo create this file
*/

package helper

import "github.com/lengpucheng/Geassgo/pkg/profess/contract"

type TaskHelper interface {
	// GetContext  获取 contract context
	GetContext() *contract.Context
	// GetTask  获取原始的Task
	GetTask() *contract.Task
	// GetLog 获取日志文件
	GetLog() *string
	// Execute 执行
	Execute() error
}
