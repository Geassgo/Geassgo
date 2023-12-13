/*
----------------------------------------
@Create 2023/12/12-17:26
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geass_system
----------------------------------------
@Version 1.0 2023/12/12
@Memo create this file
*/

package geass

import "github.com/lengpucheng/Geassgo/pkg/profess/contract"

func init() {
	RegisterGeass(System, &geassSystem{})
}

const System = "system"

type geassSystem struct{}

func (s *geassSystem) Execute(ctx contract.Context, val any) error {
	if f, ok := val.(bool); f && ok {
		// 刷新system
		ctx.GetVariable().System = contract.GenerateSystemVariable(ctx)
	}
	return nil
}

func (s *geassSystem) OverallRender() bool {
	return false
}

func (s *geassSystem) OverloadRender() (bool, any) {
	return false, nil
}
