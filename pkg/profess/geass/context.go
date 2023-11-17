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
	location  string
	variable  *contract.Variable
	item      any
	itemIndex int
	stderr    string
	stdout    string
}

func Background() *Context {
	return NewContext(context.Background(), new(contract.Variable).Check())
}

func NewContext(ctx context.Context, variable *contract.Variable) *Context {
	return &Context{
		Context:   ctx,
		variable:  variable,
		item:      nil,
		itemIndex: -1,
	}
}

func (c *Context) GetVariable() *contract.Variable {
	return c.variable
}

func (c *Context) GetItem() any {
	return c.item
}

func (c *Context) GetItemIndex() int {
	return c.itemIndex
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

func (c *Context) GenLocation() string {
	return c.location
}

func (c *Context) SubContext(item any, index int) contract.Context {
	return &Context{
		Context:   c.Context,
		variable:  c.variable,
		item:      item,
		itemIndex: index,
	}
}
