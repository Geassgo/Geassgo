/*
----------------------------------------
@Create 2023/11/17
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe context_geass 9:31
----------------------------------------
@Version 1.0 2023/11/17
@Memo create this file
*/

package geass

import (
	"context"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
)

type Context struct {
	context.Context
	contract.Runtime
	contract.Selector
	variable *contract.Variable
	stderr   string
	stdout   string
}

func Background() *Context {
	return NewContext(context.Background(), DefaultRuntime(), DefaultSelector(), new(contract.Variable).Check())
}

func NewContext(ctx context.Context, runtime contract.Runtime, selector contract.Selector, variable *contract.Variable) *Context {
	return &Context{
		Context:  ctx,
		Runtime:  runtime,
		Selector: selector,
		variable: variable,
	}
}

func (c *Context) GetVariable() *contract.Variable {
	return c.variable
}

func (c *Context) SetStdout(stdout string) {
	c.stdout = stdout
}

func (c *Context) SetStderr(stderr string) {
	c.stderr = stderr
}

func (c *Context) GetStdout() string {
	return c.stdout
}

func (c *Context) GetStderr() string {
	return c.stderr
}

func (c *Context) SubContext(runtime contract.Runtime) contract.Context {
	return &Context{
		Context:  c.Context,
		Selector: c.Selector,
		variable: c.variable,
		Runtime:  runtime,
	}
}
