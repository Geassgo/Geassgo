/*
----------------------------------------
@Create 2023/11/17
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geass_template 9:23
----------------------------------------
@Version 1.0 2023/11/17
@Memo create this file
*/

package geass

import (
	"github.com/lengpucheng/Geassgo/pkg/coderender"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/mod"
	"os"
	"path/filepath"
)

func init() {
	RegisterGeass(Template, &executorTemplate{})
}

const Template = "template"

type executorTemplate struct{}

func (e *executorTemplate) Execute(ctx contract.Context, val any) error {
	tem := val.(*mod.Template)
	files, err := os.ReadFile(coderender.AbsPath(filepath.Join(ctx.GetLocation(), "templates/"), tem.Src))
	if err != nil {
		return err
	}
	render, err := RenderStr(ctx, string(files))
	return coderender.WriteFile(tem.Dest, []byte(render), os.ModePerm|os.ModeAppend)
}

func (e *executorTemplate) OverallRender() bool {
	return true
}

func (e *executorTemplate) OverloadRender() (bool, any) {
	return true, &mod.Template{}
}
