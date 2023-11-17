package main

import (
	"fmt"
	"github.com/lengpucheng/Geassgo/pkg/profess/contract"
	"github.com/lengpucheng/Geassgo/pkg/profess/geass"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

func test() {
	file, err := os.ReadFile("C:\\dev\\go\\src\\Geassgo\\example\\execute_task\\task.yaml")
	if err != nil {
		panic(err)
	}
	var task = contract.Task{}
	if err = yaml.Unmarshal(file, &task); err != nil {
		panic(err)
	}
	if err = geass.Execute(nil, geass.Task, task); err != nil {
		panic(err)
	}
}

func main() {
	now := time.Now()
	test()
	fmt.Println("用时： ", time.Now().UnixMilli()-now.UnixMilli(), "ms")
}
