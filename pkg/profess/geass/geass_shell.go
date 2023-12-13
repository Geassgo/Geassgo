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
		// 当shell存在者执行 Bash
		if shell.Shell != "" {
			slog.Info("execute Shell:-> " + shell.Shell)
			std, ste, err = coderender.ExecShell(ctx, shell.Shell).Result()
			// 否则执行sh命令
		} else {
			slog.Info("execute Sh:-> " + shell.Command)
			std, ste, err = coderender.ExecCommandPrompt(ctx, shell.Command).Result()
		}

	case "windows":
		// 当shell存在则执行PowerShell
		if shell.WinShell != "" {
			slog.Info("execute PowerShell:-> " + shell.WinShell)
			std, ste, err = coderender.ExecShell(ctx, shell.WinShell).Result2Utf8()
			// 否则执行命令提示符 commandPrompt
		} else {
			slog.Info("execute CommandPrompt:-> " + shell.WinCommand)
			std, ste, err = coderender.ExecShell(ctx, shell.WinCommand).Result2Utf8()
		}
		// 为避免windows的GBK编码导致的乱码情况
		// 对输出的内容进行格式判断和转换
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
