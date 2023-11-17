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
	"github.com/lengpucheng/Geassgo/pkg/profess/mod"
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
	case mod.FilesDelete:
		return os.Remove(mFiles.Dest)
	case mod.FilesMkdir:
		return os.MkdirAll(mFiles.Dest, 0755)
	case mod.FilesFile:
		content, _ := RenderStr(ctx, mFiles.Content)
		return coderender.WriteFile(mFiles.Dest, []byte(content), os.ModePerm)
	default:
		return errors.New("the action is not support")
	}
}

func (e *geassFiles) OverallRender() bool {
	return false
}

func (e *geassFiles) OverloadRender() (bool, any) {
	return true, &mod.Files{}
}
