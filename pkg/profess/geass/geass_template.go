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
	"github.com/lengpucheng/Geassgo/pkg/profess/contract/mod"
	"os"
	"path/filepath"
)

func init() {
	RegisterGeass(Template, &geassTemplate{})
}

const Template = "template"

type geassTemplate struct{}

func (g *geassTemplate) Execute(ctx contract.Context, val any) error {
	tem := val.(*mod.Template)
	files, err := os.ReadFile(coderender.AbsPath(g.GetTemplateDirPath(ctx), tem.Src))
	if err != nil {
		return err
	}
	render, err := RenderStr(ctx, string(files))
	if err != nil {
		return err
	}
	return coderender.WriteFile(tem.Dest, []byte(render), os.ModePerm|os.ModeAppend)
}

func (g *geassTemplate) OverallRender() bool {
	return true
}

func (g *geassTemplate) OverloadRender() (bool, any) {
	return true, &mod.Template{}
}

func (g *geassTemplate) GetTemplateDirPath(ctx contract.Context) string {
	if ctx.GetRolePath() != "" {
		return filepath.Join(ctx.GetRolePath(), "templates")
	}
	return filepath.Dir(ctx.GetLocation())
}
