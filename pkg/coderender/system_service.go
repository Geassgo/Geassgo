/*
----------------------------------------
@Create 2023/12/13-9:44
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe 系统服务启动停止
----------------------------------------
@Version 1.0 2023/12/13
@Memo create this file
*/

package coderender

import (
	"context"
	"errors"
	"fmt"
	"github.com/lengpucheng/Geassgo/pkg/geasserr"
	"runtime"
)

// ServiceReload 重新加载服务列表
func ServiceReload(ctx context.Context) error {
	var ste string
	var err error
	switch runtime.GOOS {
	case "windows":
		// 当前场景下 powerShell 不需要加载服务
		return err
	case "linux":
		_, ste, err = ExecShell(ctx, "systemctl daemon-reload").Result()
	default:
		return geasserr.NotSupportSystem.New()
	}
	if err != nil {
		return err
	}
	if ste != "" {
		return errors.New(fmt.Sprintf("reload service fail,error is -> %s", ste))
	}
	return nil
}

// ServiceEnable 启用服务列表
// windows 下需要管理员权限
func ServiceEnable(ctx context.Context, name string) error {
	var ste string
	var err error
	switch runtime.GOOS {
	case "windows":
		_, ste, err = ExecShell(ctx, "Set-Service -Name ", name, " -StartupType Automatic").Result2Utf8()
	case "linux":
		_, ste, err = ExecShell(ctx, "systemctl enable", name).Result()
	default:
		return geasserr.NotSupportSystem.New()
	}
	if err != nil {
		return err
	}
	if ste != "" {
		return errors.New(fmt.Sprintf("enable service %s fail,error is -> %s", name, ste))
	}
	return nil
}

// ServiceDisable 禁用服务
// windows 下需要管理员权限
func ServiceDisable(ctx context.Context, name string) error {
	var ste string
	var err error
	switch runtime.GOOS {
	case "windows":
		_, ste, err = ExecShell(ctx, "Set-Service -Name ", name, " -StartupType Manual").Result2Utf8()
	case "linux":
		_, ste, err = ExecShell(ctx, "systemctl disable", name).Result()
	default:
		return geasserr.NotSupportSystem.New()
	}
	if err != nil {
		return err
	}
	if ste != "" {
		return errors.New(fmt.Sprintf("disable service %s fail,error is -> %s", name, ste))
	}
	return nil
}

// ServiceStop 停止服务
func ServiceStop(ctx context.Context, name string) error {
	var ste string
	var err error
	switch runtime.GOOS {
	case "windows":
		_, ste, err = ExecShell(ctx, "Stop-Service -Name", name).Result2Utf8()
	case "linux":
		_, ste, err = ExecShell(ctx, "systemctl stop", name).Result()
	default:
		return geasserr.NotSupportSystem.New()
	}
	if err != nil {
		return err
	}
	if ste != "" {
		return errors.New(fmt.Sprintf("stop service %s fail,error is -> %s", name, ste))
	}
	return nil
}

// ServiceStart 启动服务
func ServiceStart(ctx context.Context, name string) error {
	var ste string
	var err error
	switch runtime.GOOS {
	case "windows":
		_, ste, err = ExecShell(ctx, "Start-Service -Name", name).Result2Utf8()
	case "linux":
		_, ste, err = ExecShell(ctx, "systemctl start", name).Result()
	default:
		return geasserr.NotSupportSystem.New()
	}
	if err != nil {
		return err
	}
	if ste != "" {
		return errors.New(fmt.Sprintf("start service %s fail,error is -> %s", name, ste))
	}
	return nil
}

// ServiceRestart 重启服务
func ServiceRestart(ctx context.Context, name string) error {
	var ste string
	var err error
	switch runtime.GOOS {
	case "windows":
		_, ste, err = ExecShell(ctx, "Restart-Service -Name", name).Result2Utf8()
	case "linux":
		_, ste, err = ExecShell(ctx, "systemctl restart", name).Result()
	default:
		return geasserr.NotSupportSystem.New()
	}
	if err != nil {
		return err
	}
	if ste != "" {
		return errors.New(fmt.Sprintf("restart service %s fail,error is -> %s", name, ste))
	}
	return nil
}
