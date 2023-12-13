/*
----------------------------------------
@Create 2023/12/12-21:45
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe system_os
----------------------------------------
@Version 1.0 2023/12/12
@Memo create this file
*/

package coderender

import (
	"context"
	"runtime"
	"strings"
)

// GetOsRelease 获取系统的发行版本
// 获取失败时候显示Unknown
func GetOsRelease(ctx context.Context) string {
	switch runtime.GOOS {
	case "windows":
		// windows 直接返回
		return "Microsoft Windows"
	case "linux":
		std, _, err := ExecShell(ctx, `cat /etc/os-release |grep ID= | grep -v VERSION |awk -F "[\"]" '{print $2}'`).Result()
		if err == nil {
			return std
		}
	}
	return "unknown"
}

// GetOSVersion 获取操作系统版本
// 获取失败时候显示Unknown
func GetOSVersion(ctx context.Context) string {
	switch runtime.GOOS {
	case "linux":
		return _osVersion4Linux(ctx)
	case "windows":
		return _osVersion4Windows(ctx)
	}
	return "unknown"
}

// GetOsKernel 获取操作系统内核版本
func GetOsKernel(ctx context.Context) string {
	switch runtime.GOOS {
	case "windows":
		return _osVersion4Windows(ctx)
	case "linux":
		std, _, err := ExecShell(ctx, "uname -r").Result()
		if err == nil {
			return std
		}
	}
	return "unknown"
}

// 获取linux 的操作系统发行版本
func _osVersion4Linux(ctx context.Context) string {
	std, _, err := ExecShell(ctx, `cat /etc/os-release |grep VERSION_ID |awk -F "[\"]" '{print $2}'`).Result()
	if err != nil {
		return "unknown"
	}
	return std
}

// 获取linux 的操作系统发行版本
func _osVersion4Windows(ctx context.Context) string {
	// ver 将返回 ‘Microsoft Windows [version a.bb.ccc.dd]’
	std, _, err := ExecCommandPrompt(ctx, "ver").Result2Utf8()
	if err != nil {
		return "unknown"
	} else if len(std) <= len("\r\nMicrosoft Windows ") {
		return "unknown"
	}
	// 取中括号内的空格后的版本
	std = std[len("\r\nMicrosoft Windows "):]
	version := strings.Split(std, " ")[1]
	if len(version) > 1 {
		version = version[:len(version)-1]
	}
	return version
}
