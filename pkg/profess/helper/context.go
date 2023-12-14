/*
----------------------------------------
@Create 2023/11/17-17:04
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe context_helper
----------------------------------------
@Version 1.0 2023/11/17
@Memo create this file
*/

package helper

import (
	"context"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/geass"
)

type Context struct {
	contract.Context
	subContext []contract.Context
}

func NewContext(ctx context.Context, runtime contract.Runtime, selector contract.Selector, variable *contract.Variable) *Context {
	variable.Check()
	return &Context{
		Context:    geass.NewContext(ctx, runtime, selector, variable),
		subContext: make([]contract.Context, 0),
	}
}

func (c *Context) SubContext(runtime contract.Runtime) contract.Context {
	ctx := &Context{Context: c.Context.SubContext(runtime)}
	c.subContext = append(c.subContext, ctx)
	return ctx
}

func (c *Context) GetStdout() string {
	if c.Context.GetStdout() == "" {
		for _, sc := range c.subContext {
			if sc.GetStdout() != "" {
				c.Context.SetStdout(sc.GetStdout())
			}
		}
	}
	return c.Context.GetStdout()
}

func (c *Context) GetStderr() string {
	if c.Context.GetStderr() == "" {
		for _, sc := range c.subContext {
			if sc.GetStderr() != "" {
				c.Context.SetStderr(sc.GetStderr())
			}
		}
	}
	return c.Context.GetStderr()
}
