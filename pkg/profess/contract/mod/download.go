/*
----------------------------------------
@Create 2023/12/13-14:35
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe download
----------------------------------------
@Version 1.0 2023/12/13
@Memo create this file
*/

package mod

// Download 从远程场景下拷贝下载
// 当为本地执行时候 即为从本地获取 src 路径拷贝至 dest 目录
// 为远程主机时 则为从远程主机的 dest 路径拷贝至 本地的dest路径
type Download struct {
	Src  string `json:"src" yaml:"src"`   // 待下载文件的地址
	Dest string `json:"dest" yaml:"dest"` // 保存的本机地址
}
