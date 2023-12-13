/*
----------------------------------------
@Create 2023/12/13-15:12
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geass_fetch
----------------------------------------
@Version 1.0 2023/12/13
@Memo create this file
*/

package geass

import (
	"github.com/lengpucheng/Geassgo/pkg/coderender"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract/mod"
)

func init() {
	RegisterGeass(Upload, &geassUpload{})
}

const Upload = "upload"

type geassUpload struct{}

func (g *geassUpload) Execute(ctx contract.Context, val any) error {
	download := val.(*mod.Upload)
	// 当非远程环境时 启用本地拷贝
	if !ctx.GetVariable().System.Remote.Enable {
		return coderender.CopyFiles(download.Src, download.Dest)
	}
	// 远程环境时启用Http调用远程接口通知从本地获取文件 远程主机并调用 http/s 下载并保存
	// 或调用远程主机的上传接口直接上传文件
	// TODO 这里补充远程上传  当前暂不实现远程上传
	return nil
}

func (g *geassUpload) OverallRender() bool {
	return true
}

func (g *geassUpload) OverloadRender() (bool, any) {
	return true, &mod.Upload{}
}
