/*
----------------------------------------
@Create 2023/12/13-14:38
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe upload
----------------------------------------
@Version 1.0 2023/12/13
@Memo create this file
*/

package mod

// Upload 仅在远程控制下生效,若为本地直接执行(非远控场景）,将只是执行copy 命令
type Upload struct {
	Src  string `json:"src" yaml:"src"`   // 待上传的文件路径
	Dest string `json:"dest" yaml:"dest"` // 保存到远程节点的路径
}
