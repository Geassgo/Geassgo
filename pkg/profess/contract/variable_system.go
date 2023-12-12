/*
----------------------------------------
@Create 2023/12/12-16:37
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe variable_system
----------------------------------------
@Version 1.0 2023/12/12
@Memo create this file
*/

package contract

import (
	"context"
	"fmt"
	"github.com/lengpucheng/Geassgo/pkg/coderender"
	"log/slog"
	"net"
	"os"
	"runtime"
)

type System struct {
	Ipv4s       []string           `json:"ipv4s"`       // ipv4 栈 不包括回环网卡
	Ipv6s       []string           `json:"ipv6s"`       // ipv6 栈 不包括回环网卡
	Inters      []string           `json:"inters"`      // 适配器栈
	Adapters    []Adapter          `json:"adapters"`    // 适配器
	AdaptersMap map[string]Adapter `json:"adaptersMap"` // 名称适配器映射
	CpuNum      int                `json:"CpuNum"`      // cpu 核心数
	Hostname    string             `json:"hostname"`    // 主机名称
	Os          string             `json:"os"`          // 系统类型
	Arch        string             `json:"arch"`        // 系统架构
	Release     string             `json:"release"`     // 系统发行名称 linux -> /etc/os-release windows-> ver
	Version     string             `json:"version"`     // 系统发行版本
	Kernel      string             `json:"kernel"`      // 内核版本
	Name        string             `json:"systemName"`  // 操作系统名称
}

// Adapter 网络适配器
type Adapter struct {
	Name     string    `json:"name"`     // 适配器名称
	Loopback bool      `json:"loopback"` // 回环网卡
	Up       bool      `json:"up"`       // 是否启用
	Ipv4     []Address `json:"ipv4"`     // ipv4 地址
	Ipv6     []Address `json:"ipv6"`     // ipv6 地址
}

type Address struct {
	Ip   string `json:"ip"`   // ip
	Mask string `json:"mask"` // 掩码
}

func GenerateSystemVariable(ctx context.Context) System {
	system := System{
		Ipv4s:       []string{},
		Ipv6s:       []string{},
		Inters:      []string{},
		Adapters:    []Adapter{},
		AdaptersMap: map[string]Adapter{},
		CpuNum:      runtime.NumCPU(),
		Hostname:    func() string { hostname, _ := os.Hostname(); return hostname }(),
		Os:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		Release:     coderender.GetOsRelease(ctx),
		Version:     coderender.GetOSVersion(ctx),
		Kernel:      coderender.GetOsKernel(ctx),
	}
	// 获取网络相关信息
	if err := _variableSystemGenerateAdapters(&system); err != nil {
		slog.Error("get variable system adapter fail!", "error", err.Error())
	}
	system.Name = fmt.Sprintf("%s %s", system.Release, system.Version)
	return system
}

func _variableSystemGenerateAdapters(system *System) error {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}
	for _, inter := range interfaces {
		var adapter = Adapter{
			Name: inter.Name,
			Up:   inter.Flags&net.FlagUp != 0,
			Ipv4: []Address{},
			Ipv6: []Address{},
		}
		addrs, err := inter.Addrs()
		// 获取Address 错误则跳过
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ips, ok := addr.(*net.IPNet); ok {
				// 若干4字节转换为nil 则表示无IPV4地址 可能为IPV6
				ip := ips.IP.String()
				if ips.IP.To4() == nil {

					adapter.Ipv6 = append(adapter.Ipv6, Address{Ip: ip, Mask: net.IP(ips.Mask).String()})
					// 实际可用ip添加到ipv6 栈中
					if adapter.Up && inter.Flags&net.FlagLoopback == 0 {
						system.Ipv6s = append(system.Ipv6s, ip)
					}
					// 否则4字节不为空则一定为 IPV4
				} else {
					adapter.Ipv4 = append(adapter.Ipv4, Address{Ip: ip, Mask: net.IP(ips.Mask).String()})
					// 实际可用ip添加到ipv4 栈中
					if adapter.Up && inter.Flags&net.FlagLoopback == 0 {
						system.Ipv4s = append(system.Ipv4s, ip)
					}
				}
			}
		}
		// 添加到system的适配器列表中
		system.Adapters = append(system.Adapters, adapter)
		// 添加到网络映射中
		system.Inters = append(system.Inters, adapter.Name)
		system.AdaptersMap[adapter.Name] = adapter
	}
	return nil
}
