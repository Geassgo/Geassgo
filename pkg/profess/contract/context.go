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

type Context interface {
	context.Context
	GetVariable() *Variable
	GetItem() any
	GetItemIndex() int
	GenLocation() string
	SubContext(item any, index int) Context
	SetStdout(stdout string)
	SetStderr(stderr string)
	GetStdout() string
	GetStderr() string
}
