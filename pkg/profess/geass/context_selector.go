/*
----------------------------------------
@Create 2023/12/14-9:39
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe context_filter
----------------------------------------
@Version 1.0 2023/12/14
@Memo create this file
*/

package geass

import "github.com/lengpucheng/Geassgo/pkg/profess/contract"

// Selector 任务执行过滤器
type Selector struct {
	// 标签 若执行的task存在tags 者必须完全一致时才会执行该任务
	// 特别的 将定义如下特殊关键字 使用关键字时不可以和其他tag合用
	// all 	无论如何执行全部的任务（待tag和不带tag）
	// only 无论如何仅执行全部带tag的任务
	Tags     []string `json:"tags" yaml:"tags"`
	SkipTags []string `json:"skipTags" yaml:"skipTags"` // 跳过标签
	Action   string   `json:"action" yaml:"action"`     // 动作 可为空
}

// Default4Nil 判断当前select是否为nil 否则返回Default
func Default4Nil(selector contract.Selector) contract.Selector {
	if selector != nil {
		return selector
	}
	return DefaultSelector()
}

// DefaultSelector 获取一个默认的select
func DefaultSelector() *Selector {
	return &Selector{
		Tags:     nil,
		SkipTags: nil,
		Action:   "",
	}
}

// NewSelector 初始化应该新的过滤器
func NewSelector(action string, tags, skipTags []string) *Selector {
	return &Selector{
		Tags:     tags,
		SkipTags: skipTags,
		Action:   action,
	}
}

func (c *Selector) GetTags() []string {
	return c.Tags
}

func (c *Selector) GetSkipTags() []string {
	return c.SkipTags
}

func (c *Selector) GetAction() string {
	return c.Action
}
