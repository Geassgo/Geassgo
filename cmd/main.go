/*
----------------------------------------
@Create 2023/11/17-14:35
@Author lpc<lengpucheng@qq.com>
@Program Geassgo
@Describe main
----------------------------------------
@Version 1.0 2023/11/17
@Memo create this file
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/lengpucheng/Geassgo/pkg/profess/helper"
	"time"
)

var valuePath string
var taskPath string

func init() {
	flag.StringVar(&valuePath, "v", "", "values file")
	flag.StringVar(&taskPath, "t", "", "task file")
}

func main() {
	flag.Parse()
	if taskPath == "" {
		flag.Usage()
		return
	}
	now := time.Now()
	_, err := helper.RunTask(context.Background(), taskPath, valuePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("用时： ", time.Now().UnixMilli()-now.UnixMilli(), "ms")
}
