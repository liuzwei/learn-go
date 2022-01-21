package main

import (
	"fmt"
	"learn-go/test1"
	_ "learn-go/test1"
)

// 主入口函数
func main() {
	fmt.Println("2.主函数")
	// # command-line-arguments
	//.\main.go:11:2: undefined: test1
	test1.SayHello()
}

// 初始化函数
func init() {
	fmt.Println("1.初始化函数")
}
