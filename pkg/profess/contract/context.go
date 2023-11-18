/*
----------------------------------------
@Create 2023/11/16
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe 契约上下文
----------------------------------------
@Version 1.0 2023/11/16
@Memo create this file
*/

package contract

import "context"

// Context 执行上下文对象
type Context interface {
	context.Context
	Runtime
	GetVariable() *Variable
	SubContext(runtime Runtime) Context
	SetStdout(stdout string)
	SetStderr(stderr string)
	GetStdout() string
	GetStderr() string
}

// Runtime geass 运行时环境
type Runtime interface {
	GetItem() any
	GetItemIndex() int
	GetLocation() string
}
