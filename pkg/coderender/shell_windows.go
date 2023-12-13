//go:build windows

package coderender

import (
	"context"
	"fmt"
	"os/exec"
)

// 初始化shell命令
func initShellCommand(ctx context.Context, command string, args ...any) *exec.Cmd {
	if args != nil && len(args) > 0 {
		for _, arg := range args {
			command = fmt.Sprintf("%s %s", command, arg)
		}
	}
	return exec.CommandContext(ctx, "PowerShell", command)
}

// 初始化命令提示符
func initCommandPrompt(ctx context.Context, command string, args ...any) *exec.Cmd {
	if args != nil && len(args) > 0 {
		for _, arg := range args {
			command = fmt.Sprintf("%s %v", command, arg)
		}
	}
	return exec.CommandContext(ctx, "cmd", "/C", command)
}

// 去掉末尾的换行符
func cutTailLineBreak(str string) string {
	if len(str) >= 2 && str[len(str)-2:] == "\r\n" {
		return str[:len(str)-2]
	}
	return str
}
