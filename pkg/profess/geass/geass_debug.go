/*
----------------------------------------
@Create 2023/12/14-11:41
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geass_debug
----------------------------------------
@Version 1.0 2023/12/14
@Memo create this file
*/

package geass

import (
	"fmt"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract/mod"
)

func init() {
	RegisterGeass(Debug, &geassDebug{})
}

const Debug = "debug"

type geassDebug struct{}

func (g *geassDebug) Execute(ctx contract.Context, val any) error {
	debug := val.(*mod.Debug)
	str, err := RenderStr(ctx, debug.Msg)
	if err != nil {
		return err
	}
	fmt.Println(str)
	return nil
}

func (g *geassDebug) OverallRender() bool {
	return false
}

func (g *geassDebug) OverloadRender() (bool, any) {
	return true, &mod.Debug{}
}
