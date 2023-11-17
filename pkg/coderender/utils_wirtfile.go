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
