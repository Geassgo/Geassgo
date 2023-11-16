package coderender

import (
	"bytes"
	"context"
)

// ExecShell 执行shell
func ExecShell(ctx context.Context, shell string, args ...any) (string, string, error) {
	command := initShellCommand(ctx, shell, args...)
	var std, ste = new(bytes.Buffer), new(bytes.Buffer)
	command.Stdout, command.Stderr = std, ste
	if err := command.Start(); err != nil {
		return std.String(), ste.String(), err
	}
	if err := command.Wait(); err != nil {
		return std.String(), ste.String(), err
	}
	return cutTailLineBreak(std.String()), cutTailLineBreak(ste.String()), nil
}

// 去掉末尾的换行符
func cutTailLineBreak(str string) string {
	if len(str) >= 2 && str[len(str)-2:] == "\r\n" {
		return str[:len(str)-2]
	}
	return str
}
