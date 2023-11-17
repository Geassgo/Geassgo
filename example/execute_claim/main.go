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
	"github.com/lengpucheng/Geassgo/pkg/profess/helper"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

func test() {
	file, err := os.ReadFile("example/execute_claim/values.yaml")
	if err != nil {
		panic(err)
	}
	var values = &contract.Variable{Values: make(map[string]any)}
	if err = yaml.Unmarshal(file, &values.Values); err != nil {
		panic(err)
	}
	ctx := helper.NewContext(context.Background(), values)
	if err = helper.LoadAndExecute4File(ctx, "example/execute_claim/claim.yaml"); err != nil {
		panic(err)
	}
	fmt.Println(ctx)
}

func main() {
	now := time.Now()
	test()
	fmt.Println("用时： ", time.Now().UnixMilli()-now.UnixMilli(), "ms")
}
