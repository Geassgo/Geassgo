/*
----------------------------------------
@Create 2023/12/12-14:25
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe geass_service
----------------------------------------
@Version 1.0 2023/12/12
@Memo create this file
*/

package geass

import (
	"fmt"
	"github.com/lengpucheng/Geassgo/pkg/coderender"
	"github.com/lengpucheng/Geassgo/pkg/geasserr"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract/mod"
)

func init() {
	RegisterGeass(Service, &geassService{})
}

const Service = "service"

type geassService struct{}

func (s *geassService) Execute(ctx contract.Context, val any) error {
	service := val.(*mod.Service)
	var err error
	// 重新加载服务配置
	if service.Reload {
		if err = coderender.ServiceReload(ctx); err != nil {
			return fmt.Errorf("reload service fail....%s", err.Error())
		}
	}
	switch service.State {
	case mod.ServiceRestart:
		err = coderender.ServiceRestart(ctx, service.Name)
	case mod.ServiceStart:
		err = coderender.ServiceStart(ctx, service.Name)
	case mod.ServiceStop:
		err = coderender.ServiceStop(ctx, service.Name)
	default:
		return geasserr.NotSupportSystem.New()
	}
	if err != nil {
		return err
	}
	// 设置服务自启动服务配置
	if service.Enable != nil {
		if *service.Enable {
			err = coderender.ServiceEnable(ctx, service.Name)
		} else {
			err = coderender.ServiceDisable(ctx, service.Name)
		}
	}
	return err
}

func (s *geassService) OverallRender() bool {
	return true
}

func (s *geassService) OverloadRender() (bool, any) {
	return true, &mod.Service{}
}
