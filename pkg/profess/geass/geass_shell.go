/*
----------------------------------------
@Create 2023/11/16
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geass_shell 16:52
----------------------------------------
@Version 1.0 2023/11/16
@Memo create this file
*/

package geass

import (
	"github.com/lengpucheng/Geassgo/pkg/coderender"
	"github.com/lengpucheng/Geassgo/pkg/geasserr"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"runtime"
)

func init() {
	registerGeass(Shell, &geassShell{})
}

const Shell = "shell"

type geassShell struct{}

func (g *geassShell) Execute(ctx contract.Context, val any) error {
	shell, ok := val.(contract.Shell)
	if !ok {
		return geasserr.ModuleValueNotSupport.New()
	}
	var std, ste string
	var err error
	switch runtime.GOOS {
	case "linux":
		std, ste, err = coderender.ExecShell(ctx, shell.Shell)
	case "windows":
		std, ste, err = coderender.ExecShell(ctx, shell.WinShell)
		std = string(coderender.EncodeAuto2Utf8([]byte(std)))
		ste = string(coderender.EncodeAuto2Utf8([]byte(ste)))
	}
	ctx.SetStdout(std)
	ctx.SetStderr(ste)
	return err
}

func (g *geassShell) OverallRender() bool {
	return true
}

func (g *geassShell) executeShell(ctx contract.Context, shell string) (string, string, error) {

	return coderender.ExecShell(ctx, shell)
}
