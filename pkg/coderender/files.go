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
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
	return filepath.Join(location, path)
}

// IsNotExist 判断文件夹是否不存在
func IsNotExist(dir string) bool {
	_, err := os.Stat(dir)
	return err != nil && os.IsNotExist(err)
}

// WriteFile 创建写入文件或者覆盖写入文件
// dest 目标位置 当上层目录不存在时会自动创建
// data 需要写入的内容, 使用string时 需要使用 []byte(str) 进行包装
// perm 写入文件的权限 默认应该为 os.ModePerm 777最高权限
func WriteFile(dest string, data []byte, perm os.FileMode) error {
	return writeFile(dest, data, os.O_WRONLY|os.O_CREATE, perm)
}

// WriteFileAdd 在文件末尾进行追加,若不存在将创建
// dest 目标位置 当上层目录不存在时会自动创建
// data 需要写入的内容, 使用string时 需要使用 []byte(str) 进行包装
// perm 写入文件的权限 默认应该为 os.ModePerm 777最高权限
func WriteFileAdd(dest string, data []byte, perm os.FileMode) error {
	return writeFile(dest, data, os.O_APPEND|os.O_WRONLY|os.O_CREATE, perm)
}

// MoveFiles 移动文件/文件夹
// src 源文件或文件夹 若原路径以 /或 \ 结尾 则为拷贝旗下全部内容
// dest 目标文件或文件夹 当dest存在且为文件夹时 将会把 src 移动到 dest 目录下
func MoveFiles(src, dest string) error {
	// 若不为文件夹下全部内容则判断文件夹或文件是否存在

	if _, err := os.Stat(src); err != nil {
		slog.Error("Failed to obtain move Source status")
		return err
	}

	// 以分隔符结尾 则为复制其下全部内容
	if src[len(src)-1:] == "/" || src[len(src)-1:] == "\\" {
		items, err := os.ReadDir(src)
		if err != nil {
			slog.Error("Failed to open move Source dir")
			return err
		}
		// 遍历文件夹中每一个文件并并进行拷贝
		for _, item := range items {
			if err = MoveFiles(filepath.Join(src, item.Name()), filepath.Join(dest, item.Name())); err != nil {
				return err
			}
		}
	}

	// 对于非拷贝文件夹下内容则对单个文件夹或文件夹进行拷贝
	info, err := os.Stat(dest)
	if err == nil && info.IsDir() {
		// 目标位置是文件夹且存在
		// 将源文件移动到目标文件夹内部
		subPath := filepath.Join(dest, strings.TrimPrefix(src, filepath.Dir(src)))
		//// 如果下的子目录也存在且开启了强制覆盖则对同级目录下的文件进行单独移动
		//subInfo, err := os.Stat(subPath)
		//if err == nil && subInfo.IsDir() && srcDir {
		//	return MoveFiles(src+"/", subPath)
		//}
		return os.Rename(src, subPath)
	}
	// 目标部署文件夹或者目标不存在者将目标作为移动后的文件名称
	return os.Rename(src, dest)
}

// CopyFiles 拷贝文件
// src 待拷贝的文件或文件夹 文件夹将递归拷贝
// dest 目标位置 目标位置应该为文件夹
func CopyFiles(src, dest string) error {
	stat, err := os.Stat(src)
	if err != nil {
		slog.Error("Failed to obtain Copy Source status")
		return err
	}
	// 如果为单个文件就单独拷贝
	if !stat.IsDir() {
		return copyFile(src, filepath.Join(dest, strings.TrimPrefix(src, filepath.Dir(src))))
	}
	// 如果为文件夹就递归拷贝
	return filepath.Walk(src, func(filePath string, info fs.FileInfo, err error) error {
		// 当为文件夹时则创建文件
		nDest := filepath.Join(dest, strings.TrimPrefix(filePath, src))
		if info.IsDir() {
			if err = os.MkdirAll(nDest, 0755); err != nil {
				slog.Error("create dir fail in copy files", "filePath", nDest)
				return err
			}
			return nil
		}
		// 普通文件就拷贝
		return copyFile(filePath, nDest)
	})
}

// MkLink 创建软连接
// ctx 上下文对象
// src 被连接的地址 （要连接的目标路径）
// dest 软连接生成的路径
func MkLink(ctx context.Context, src, dest string) error {
	stat, e := os.Stat(src)
	if e != nil {
		slog.Error("Creating soft connection failed, original path does not exist", "srcPath", src)
		return e
	}
	var err error
	var stderr string
	switch runtime.GOOS {
	case "windows":
		if stat.IsDir() {
			_, stderr, err = ExecShell(ctx, "mklink /J ", dest, src)
		} else {
			_, stderr, err = ExecShell(ctx, "mklink", dest, src)
		}
	case "linux":
		_, stderr, err = ExecShell(ctx, "ln -s", src, dest)
	default:
		return errors.New(fmt.Sprintf("no support this os %s", runtime.GOOS))
	}
	if err == nil && stderr != "" {
		err = errors.New(fmt.Sprintf("Failed to create soft connection, error reason is %s", stderr))
	}
	return err
}

// 复制文件夹
// src 源地址
// dest 目标地址 要包含到具体的文件名称
func copyFile(src string, dest string) error {
	oSrc, err := os.Open(src)
	defer oSrc.Close()
	if err != nil {
		slog.Error("open file of src fail in copy", "fileName", src)
		return err
	}
	oDest, err := os.Create(dest)
	defer oDest.Close()
	if err != nil {
		slog.Error("open file of dest fail in copy", "fileName", dest)
		return err
	}
	_, err = io.Copy(oDest, oSrc)
	return err
}

// writeFile 对文件进行操作
// dest 目标位置
// data 需要写入的内容, 使用string时 需要使用 []byte(str) 进行包装
// flag 写入的方式 若为覆盖和创建为   os.O_WRONLY|os.O_CREATE  若为追加则应该为 os.O_APPEND|os.O_WRONLY|os.O_CREATE
// perm 写入文件的权限 默认应该为 os.ModePerm 777最高权限
func writeFile(dest string, data []byte, flag int, perm os.FileMode) error {
	// 上级目录不存在则创建目录
	dir := filepath.Dir(dest)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, 0755)
	}
	// 打开并创建文件
	var err error
	var f *os.File
	f, err = os.OpenFile(dest, flag, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}
