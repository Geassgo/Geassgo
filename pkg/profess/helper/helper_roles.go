/*
----------------------------------------
@Create 2023-11-19
@Author 冷朴承<lengpucheng@qq.com>
@Program Geassgo
@Describe helper_roles
----------------------------------------
@Version 1.0 2023/11/19-1:36
@Memo create this file
*/

package helper

import (
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
)

type helperRoles struct{}

// Execute
// ctx 的location已经在roles内的
// val
func (r *helperRoles) Execute(ctx contract.Context, val any) error {
	//TODO implement me
	panic("implement me")
}

func (r *helperRoles) OverallRender() bool {
	//TODO implement me
	panic("implement me")
}

func (r *helperRoles) OverloadRender() (bool, any) {
	//TODO implement me
	panic("implement me")
}
