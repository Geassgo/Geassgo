package coderender

import (
	"bytes"
	"context"
	"os/exec"
)

// EncodeAuto2Utf84Exec 自动将Exec执行结果转换为UTF8类型
// std,ste 对应执行的标准输出和输入 将分别转换为 UTF-8
// err 为可能的错误 本方法将不对错误进行处理 原样返回
func EncodeAuto2Utf84Exec(std, ste string, err error) (string, string, error) {
	return string(EncodeAuto2Utf8([]byte(std))), string(EncodeAuto2Utf8([]byte(ste))), err
}

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
