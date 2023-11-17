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

package contract

import "context"

type executeContext struct {
	context.Context
	location  string
	variable  *Variable
	item      any
	itemIndex int
	stderr    string
	stdout    string
}

func DefaultContext() Context {
	return NewContext(context.Background(), new(Variable).Check())
}

func NewContext(ctx context.Context, variable *Variable) Context {
	return &executeContext{
		Context:   ctx,
		variable:  variable,
		item:      nil,
		itemIndex: -1,
	}
}

func (c *executeContext) GetVariable() *Variable {
	return c.variable
}

func (c *executeContext) GetItem() any {
	return c.item
}

func (c *executeContext) GetItemIndex() int {
	return c.itemIndex
}

func (c *executeContext) SetStdout(stdout string) {
	c.stdout = stdout
}

func (c *executeContext) SetStderr(stderr string) {
	c.stderr = stderr
}

func (c *executeContext) GenerateSubContext(item any, index int) Context {
	return &executeContext{
		variable:  c.variable,
		item:      c.item,
		itemIndex: c.itemIndex,
	}
}

func (c *executeContext) SetItem(item any, index int) {
	c.item = item
	c.itemIndex = index
}
