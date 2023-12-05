/*
----------------------------------------
@Create 2023/11/17-17:45
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
	"fmt"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/geass"
	"github.com/lengpucheng/Geassgo/pkg/profess/helper"
	"time"
)

func test() {
	var values = &contract.Variable{Values: make(map[string]any)}
	ctx := helper.NewContext(context.Background(), geass.DefaultRuntime(), values)
	if err := helper.LoadAndExecute4File(ctx, "example/03_execute_roles/main.yaml"); err != nil {
		panic(err)
	}
	fmt.Println(ctx)
}

func main() {
	now := time.Now()
	test()
	fmt.Println("用时： ", time.Now().UnixMilli()-now.UnixMilli(), "ms")
}
