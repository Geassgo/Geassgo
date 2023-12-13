package coderender

import (
	"bytes"
	"context"
	"os/exec"
)

// ExecShell 执行shell
// windows 	下为 powerShell
// linux 	下为 bashShell
func ExecShell(ctx context.Context, shell string, args ...any) ExecResult {
	return ExecSystemCommand(initShellCommand(ctx, shell, args...))
}

// ExecCommandPrompt 运行命令提示符
// windows	下为 CommandPrompt (cmd)
// linux	下为 sh
func ExecCommandPrompt(ctx context.Context, command string, args ...any) ExecResult {
	return ExecSystemCommand(initCommandPrompt(ctx, command, args...))
}

// ExecSystemCommand 执行go cmd
func ExecSystemCommand(command *exec.Cmd) ExecResult {
	var std, ste = new(bytes.Buffer), new(bytes.Buffer)
	command.Stdout, command.Stderr = std, ste
	if err := command.Start(); err != nil {
		return newExecResult(std, ste, err)
	}
	if err := command.Wait(); err != nil {
		return newExecResult(std, ste, err)
	}
	return newExecResult(std, ste, nil)
}

// 快捷构造ExecResult
func newExecResult(std, ste *bytes.Buffer, err error) ExecResult {
	return &execResult{std: std, ste: ste, err: err}
}

// ExecResult 执行结果
type ExecResult interface {
	// Result 获取结果
	// 三个值分贝对应 stdout stderr err
	// stdout 标准输出,通常情况下没有报错的时候的输出
	// stderr 标准传入,当执行发生错误的时候会输出内容
	// err 执行的错误,当执行发生错误以及返回非0时将纯在错误
	Result() (string, string, error)
	// Result2Utf8 自动将Exec执行结果转换为UTF8类型
	// 返回类型和 Result相同,但是stdout和stderr会自动判断类型并将GBK转换为UTF8类型
	// error 为原始错误,不会对该错误进行处理,将和Result的返回一致
	Result2Utf8() (string, string, error)
}

// ExecResult 执行结果
type execResult struct {
	std *bytes.Buffer // 标准输出
	ste *bytes.Buffer // 标准输入
	err error         // 执行错误
}

// Result 获取结果
// 三个值分贝对应 stdout stderr err
// stdout 标准输出,通常情况下没有报错的时候的输出
// stderr 标准传入,当执行发生错误的时候会输出内容
// err 执行的错误,当执行发生错误以及返回非0时将纯在错误
func (e *execResult) Result() (string, string, error) {
	return cutTailLineBreak(e.std.String()), cutTailLineBreak(e.std.String()), e.err
}

// Result2Utf8 自动将Exec执行结果转换为UTF8类型
// 返回类型和 Result相同,但是stdout和stderr会自动判断类型并将GBK转换为UTF8类型
// error 为原始错误,不会对该错误进行处理,将和Result的返回一致
func (e *execResult) Result2Utf8() (string, string, error) {
	return cutTailLineBreak(string(EncodeAuto2Utf8(e.std.Bytes()))), cutTailLineBreak(string(EncodeAuto2Utf8(e.ste.Bytes()))), e.err
}
