/*
----------------------------------------
@Create 2023/11/17
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geass_files 9:24
----------------------------------------
@Version 1.0 2023/11/17
@Memo create this file
*/

package geass

import (
	"errors"
	"github.com/lengpucheng/Geassgo/pkg/coderender"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract/mod"
	"os"
)

func init() {
	RegisterGeass(Files, &geassFiles{})
}

const Files = "file"

type geassFiles struct{}

func (e *geassFiles) Execute(ctx contract.Context, val any) error {
	mFiles := val.(*mod.Files)
	if mFiles.Dest != "" {
		mFiles.Dest, _ = RenderStr(ctx, mFiles.Dest)
	}
	switch mFiles.Action {
	case mod.FilesDelete: // 删除文件
		return os.Remove(mFiles.Dest)
	case mod.FilesMkdir: // 创建文件夹
		return os.MkdirAll(mFiles.Dest, 0755)
	case mod.FilesFile: // 写入文件
		fallthrough
	case mod.FilesFileCover: // 覆盖写入
		content, _ := RenderStr(ctx, mFiles.Content)
		return coderender.WriteFile(mFiles.Dest, []byte(content), os.ModePerm)
	case mod.FilesFileAdd: // 追加写入
		content, _ := RenderStr(ctx, mFiles.Content)
		return coderender.WriteFileAdd(mFiles.Dest, []byte(content), os.ModePerm)
	case mod.FilesMove: // 移动
		return coderender.MoveFiles(mFiles.Src, mFiles.Dest)
	case mod.FilesCopy: // 复制
		return coderender.CopyFiles(mFiles.Src, mFiles.Dest)
	case mod.FilesLink: // 连接
		return coderender.MkLink(ctx, mFiles.Src, mFiles.Dest)
	default: // 如果出现其他不存在的类型 则提升不支持
		return errors.New("the action of file is not support")
	}
}

func (e *geassFiles) OverallRender() bool {
	return false
}

func (e *geassFiles) OverloadRender() (bool, any) {
	return true, &mod.Files{}
}
