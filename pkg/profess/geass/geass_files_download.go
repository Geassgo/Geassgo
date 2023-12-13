/*
----------------------------------------
@Create 2023/12/13-15:04
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geass_download
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
	RegisterGeass(Download, &geassDownload{})
}

const Download = "download"

type geassDownload struct{}

func (g *geassDownload) Execute(ctx contract.Context, val any) error {
	download := val.(*mod.Download)
	// 当非远程环境时 启用本地拷贝
	if !ctx.GetVariable().System.Remote.Enable {
		return coderender.CopyFiles(download.Src, download.Dest)
	}
	// 远程环境时启用Http调用远程接口获取下载地址 并调用 http/s 下载并保存
	// TODO 这里补充远程下载  当前暂不实现远程下载
	return nil
}

func (g *geassDownload) OverallRender() bool {
	return true
}

func (g *geassDownload) OverloadRender() (bool, any) {
	return true, &mod.Download{}
}
