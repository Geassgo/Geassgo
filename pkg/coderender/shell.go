package coderender

import (
	"bytes"
	"context"
	"os/exec"
)

// ExecShell 执行shell
// windows 	下为 powerShell
// linux 	下为 bashShell
func ExecShell(ctx context.Context, shell string, args ...any) (string, string, error) {
	return ExecSystemCommand(initShellCommand(ctx, shell, args...))
}

// ExecCommandPrompt 运行命令提示符
// windows	下为 CommandPrompt (cmd)
// linux	下为 sh
func ExecCommandPrompt(ctx context.Context, command string, args ...any) (string, string, error) {
	return ExecSystemCommand(initCommandPrompt(ctx, command, args...))
}

// ExecSystemCommand 执行go cmd
func ExecSystemCommand(command *exec.Cmd) (string, string, error) {
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
