/*
----------------------------------------
@Create 2023/12/7-16:19
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe files_tgz
----------------------------------------
@Version 1.0 2023/12/7
@Memo create this file
*/

package coderender

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// UnArchive 解压tgz文件
// src tgz\tar.gz 文件路径
// dest 解压到的目标位置
func UnArchive(src, dest string) error {
	// 打开.tar.gz文件
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建gzip.Reader以读取.tar.gz文件
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	// 创建tar.Reader以读取gzip流
	tarReader := tar.NewReader(gzipReader)

	// 遍历tar文件中的每个文件/目录
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // 已经读取完所有文件/目录
		}
		if err != nil {
			return err
		}

		path := filepath.Join(dest, header.Name)
		// 检查文件类型（目录或文件）并相应地处理
		switch header.Typeflag {
		case tar.TypeDir:
			// 创建目录（如果目录不存在）
			err = os.MkdirAll(path, 0755)
			if err != nil {
				return err
			}
		case tar.TypeReg: // 普通文件或数据文件（非目录）
			// 创建文件（如果文件不存在）并写入数据流（从tar文件中读取的数据）
			dir := filepath.Dir(path)
			if _, err = os.Stat(dir); os.IsNotExist(err) {
				if err = os.MkdirAll(dir, 0755); err != nil {
					return err
				}
			}
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(file, tarReader)
			if err != nil {
				return err
			}
		default: // 其他类型（符号链接、门等）忽略不处理（可选）
			fmt.Printf("忽略文件类型：%c - %s\n", header.Typeflag, header.Name)
		}
	}
	fmt.Println("解压完成！")
	return nil
}

// Archive 打包并存档到指定目录
// src 为需要打包的文件夹 末尾带/表示目录下文件
// dst 为输出的tar包地址
func Archive(src, dst string) (err error) {
	// 创建文件
	fw, err := os.Create(dst)
	if err != nil {
		return
	}
	defer fw.Close()

	// 将 tar 包使用 gzip 压缩，其实添加压缩功能很简单，
	// 只需要在 fw 和 tw 之前加上一层压缩就行了，和 Linux 的管道的感觉类似
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 创建 Tar.Writer 结构
	tw := tar.NewWriter(gw)
	// 如果需要启用 gzip 将上面代码注释，换成下面的

	defer tw.Close()

	dir := filepath.Dir(src)

	// 下面就该开始处理数据了，这里的思路就是递归处理目录及目录下的所有文件和目录
	// 这里可以自己写个递归来处理，不过 Golang 提供了 filepath.Walk 函数，可以很方便的做这个事情
	// 直接将这个函数的处理结果返回就行，需要传给它一个源文件或目录，它就可以自己去处理
	// 我们就只需要去实现我们自己的 打包逻辑即可，不需要再去路径相关的事情
	return filepath.Walk(src, func(fileName string, fi os.FileInfo, err error) error {
		// 因为这个闭包会返回个 error ，所以先要处理一下这个
		if err != nil {
			return err
		}

		// 这里就不需要我们自己再 os.Stat 了，它已经做好了，我们直接使用 fi 即可
		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}
		// 这里需要处理下 hdr 中的 Name，因为默认文件的名字是不带路径的，
		// 打包之后所有文件就会堆在一起，这样就破坏了原本的目录结果
		// 例如： 将原本 hdr.Name 的 syslog 替换程 log/syslog
		// 这个其实也很简单，回调函数的 fileName 字段给我们返回来的就是完整路径的 log/syslog
		// strings.TrimPrefix 将 fileName 的最左侧的 / 去掉，
		// 熟悉 Linux 的都知道为什么要去掉这个
		hdr.Name = strings.TrimPrefix(fileName, dir+string(filepath.Separator))

		// 写入文件信息
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		// 判断下文件是否是标准文件，如果不是就不处理了，
		// 如： 目录，这里就只记录了文件信息，不会执行下面的 copy
		if !fi.Mode().IsRegular() {
			return nil
		}

		// 打开文件
		fr, err := os.Open(fileName)
		defer fr.Close()
		if err != nil {
			return err
		}

		// copy 文件数据到 tw
		n, err := io.Copy(tw, fr)
		if err != nil {
			return err
		}

		// 记录下过程，这个可以不记录，这个看需要，这样可以看到打包的过程
		log.Printf("Archive %s ,write size %d byte\n", fileName, n)

		return nil
	})
}
