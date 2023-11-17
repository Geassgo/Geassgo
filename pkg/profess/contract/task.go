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
	Mod   `json:",inline" yaml:",inline"`
	Shell `json:",inline" yaml:",inline"`
	Name  string `json:"name" yaml:"name"`
}

type Shell struct {
	Shell    string `json:"shell" yaml:"shell"`
	WinShell string `json:"win_shell" yaml:"win_shell"`
}

type Mod map[string]any
