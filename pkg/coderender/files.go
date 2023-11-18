/*
----------------------------------------
@Create 2023/11/17
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe utils_wirtfile 9:26
----------------------------------------
@Version 1.0 2023/11/17
@Memo create this file
*/

package coderender

import (
	"os"
	"path/filepath"
)

// AbsPath 返回绝对路径或拼接路径
// location 当前位置 若location非以/或\结尾将去掉最后一个 /或\后的内容
// path 待判断路径
// 若 path 为绝对路径则返回本身
// 若 path 非绝对路径 则拼接 location的文件夹路径 和 path
func AbsPath(location, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(filepath.Dir(location), path)
}

// WriteFile
// 把data写入到 dest对应的目录中
// 若 dest 目录不存在则创建
func WriteFile(dest string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(dest)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, 0755)
	}
	return os.WriteFile(dest, data, perm)
}
