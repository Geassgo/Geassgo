/*
----------------------------------------
@Create 2023/12/6-14:27
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe main
----------------------------------------
@Version 1.0 2023/12/6
@Memo create this file
*/
package main

import (
	"context"
	"fmt"
	"github.com/lengpucheng/Geassgo/pkg/coderender"
	"github.com/lengpucheng/Geassgo/pkg/profess/helper"
	"os"
	"time"
)

func main() {
	generateChart()
	now := time.Now()
	test()
	fmt.Println("用时： ", time.Now().UnixMilli()-now.UnixMilli(), "ms")
}

func test() {
	chart, err := helper.RunChart(context.Background(), "example/04_execute_charts/out/otnm-manager.tgz")
	fmt.Println(chart, err)
}

func generateChart() {
	_ = os.MkdirAll("example/04_execute_charts/out", 0755)
	// 带 main
	if err := coderender.Archive("example/04_execute_charts/otnm-manager-charts",
		"example/04_execute_charts/out/otnm-manager.tgz"); err != nil {
		panic(err)
	}
	// 不带main
	if err := coderender.Archive("example/04_execute_charts/otnm-manager-charts/roles/otnm-manager",
		"example/04_execute_charts/out/otnm-manager-nom.tgz"); err != nil {
		panic(err)
	}
}

//
//type ChartMetadata struct {
//	Name        string `json:"name" yaml:"name"`
//	Version     string `json:"version" yaml:"version"`
//	ApiVersion  string `json:"apiVersion" yaml:"apiVersion"`
//	AppVersion  string `json:"appVersion" yaml:"appVersion"`
//	Description string `json:"description" yaml:"description"`
//}
//
//type Chart struct {
//	Role
//	Name   string
//	Chart  ChartMetadata
//	Main   []byte
//	Values map[string]any
//	Roles  map[string]Role
//}
//
//type Role struct {
//	Defaults  map[string][]byte
//	Tasks     map[string][]byte
//	Templates map[string][]byte
//}
