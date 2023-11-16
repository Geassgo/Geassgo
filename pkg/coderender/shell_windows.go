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
	return exec.CommandContext(ctx, "cmd", "/C", command)
}
