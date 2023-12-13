/*
----------------------------------------
@Create 2023/11/16
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe 契约任务
----------------------------------------
@Version 1.0 2023/11/16
@Memo create this file
*/

package contract

type Task struct {
	Mod `json:",inline" yaml:",inline"`
	// 执行shell命令
	// linux上执行的bash 或 sh
	// windows上执行的powerShell 或 cmd
	// 部分命令可能需要管理员权限才可以执行 因此在windows运行时请务必以管理员方式执行
	Shell `json:",inline" yaml:",inline"`
	Name  string `json:"name" yaml:"name"`
}

type Shell struct {
	// 使用 bash 执行
	Shell string `json:"shell" yaml:"shell"`
	// 使用 PowerShell 执行
	WinShell string `json:"win_shell" yaml:"win_shell"`
	// 使用 sh 执行
	Command string `json:"command" yaml:"Command"`
	// 使用 commandPrompt 执行
	WinCommand string `json:"win_command" yaml:"win_command"`
}

type Mod map[string]any
