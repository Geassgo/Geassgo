/*
----------------------------------------
@Create 2023-11-19
@Author 冷朴承<lengpucheng@qq.com>
@Program Geassgo
@Describe helper_roles
----------------------------------------
@Version 1.0 2023/11/19-1:36
@Memo create this file
*/

package helper

import (
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/geass"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func init() {
	geass.RegisterGeass(Roles, &helperRoles{})
}

const Roles = "_role_"

type helperRoles struct{}

type Role struct {
	Default map[string]string
}

// Execute
// ctx 的location已经在roles内的
// val
func (r *helperRoles) Execute(ctx contract.Context, val any) error {
	// 加载默认变量
	variable, err := r.loadDefault(ctx)
	if err != nil {
		return err
	}
	context := NewContext(ctx, ctx, variable)
	if c, ok := ctx.(*Context); ok {
		c.subContext = append(c.subContext, context)
	}
	return LoadAndExecute4File(context, "tasks/main.yaml")
}

func (r *helperRoles) OverallRender() bool {
	return false
}

func (r *helperRoles) OverloadRender() (bool, any) {
	return false, nil
}

// 加载默认配置
func (r *helperRoles) loadDefault(ctx contract.Context) (*contract.Variable, error) {
	variable := ctx.GetVariable().DeepCopy()
	val := map[string]any{}
	defValPath := filepath.Join(ctx.GetRolePath(), "defaults/main.yaml")
	if _, err := os.Stat(defValPath); os.IsNotExist(err) {
		return variable, nil
	}
	file, err := os.ReadFile(defValPath)
	if err != nil {
		return variable, err
	}
	err = yaml.Unmarshal(file, &val)
	// 合并val
	for k, v := range val {
		if _, ok := variable.Values[k]; !ok {
			variable.Values[k] = v
		}
	}
	// TODO 这里日后需要多级渲染
	return variable, nil
}
