/*
----------------------------------------
@Create 2023/12/14-11:38
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe debug
----------------------------------------
@Version 1.0 2023/12/14
@Memo create this file
*/

package mod

type Debug struct {
	Msg string `json:"msg" yaml:"msg"` // 消息输出
	Var string `json:"var" yaml:"var"` // 变量输出
}
