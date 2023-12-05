/*
----------------------------------------
@Create 2023/11/17
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe files 9:22
----------------------------------------
@Version 1.0 2023/11/17
@Memo create this file
*/

package mod

type Files struct {
	Src     string      `json:"src"`
	Dest    string      `json:"dest"`
	Force   bool        `json:"force"`
	Action  FilesAction `json:"action"`
	Content string      `json:"content"`
	Recurse bool        `json:"recurse"` // 递归
}

type FilesAction string

const (
	FilesDelete    FilesAction = "del"        // 删除
	FilesMkdir     FilesAction = "dir"        // 创建文件夹
	FilesFile      FilesAction = "file"       // 创建文件
	FilesFileAddon FilesAction = "file_add"   // 追加
	FilesFileCover FilesAction = "file_cover" // 覆盖
	FilesCopy      FilesAction = "copy"       // 拷贝
	FilesMove      FilesAction = "move"       // 移动
	FilesLink      FilesAction = "link"       // 连接
)
