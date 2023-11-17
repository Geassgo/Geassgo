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
	"log/slog"
	"runtime"
)

func init() {
	RegisterGeass(Shell, &geassShell{})
}

const Shell = "shell"

type geassShell struct{}

func (g *geassShell) Execute(ctx contract.Context, val any) error {
	shell, ok := val.(*contract.Shell)
	if !ok {
		return geasserr.ModuleValueNotSupport.New(val)
	}
	var std, ste string
	var err error
	switch runtime.GOOS {
	case "linux":
		slog.Info("execute Shell:->" + shell.Shell)
		std, ste, err = coderender.ExecShell(ctx, shell.Shell)
	case "windows":
		slog.Info("execute Shell:-> " + shell.WinShell)
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

func (g *geassShell) OverloadRender() (bool, any) {
	return true, &contract.Shell{}
}
