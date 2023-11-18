/*
----------------------------------------
@Create 2023-11-19
@Author 冷朴承<lengpucheng@qq.com>
@Program Geassgo
@Describe runtime
----------------------------------------
@Version 1.0 2023/11/19-0:45
@Memo create this file
*/

package geass

import "os"

type Runtime struct {
	location  string // 运行时位置
	item      any    // 运行时迭代器
	itemIndex int    // 迭代器位置
}

// DefaultRuntime 获取一个默认的空runtime
func DefaultRuntime() *Runtime {
	return NewRuntime("", -1, nil)
}

// NewRuntime 初始化一个geass执行时环境
func NewRuntime(location string, index int, item any) *Runtime {
	if size := len(location); size > 0 && location[size-1] != '/' && location[size-1] != '\\' {
		location = location + string(os.PathSeparator)
	}
	return &Runtime{
		location:  location,
		item:      item,
		itemIndex: index,
	}
}

func (r *Runtime) GetItem() any {
	return r.item
}

func (r *Runtime) GetItemIndex() int {
	return r.itemIndex
}

func (r *Runtime) GetLocation() string {
	return r.location
}
