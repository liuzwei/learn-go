package main

import (
	"fmt"
	"learn-go/test1"
	//_ "learn-go/test1"
)

// 主入口函数
func main() {
	fmt.Println("2.主函数")
	// # command-line-arguments
	//.\main.go:11:2: undefined: test1
	//test1.SayHello()

	// 声明变量
	var name string

	fmt.Printf("name default value is [%s]", name)

	// rune测试
	runeTest()

	fmt.Println()
	// 打印格式
	printSomething()
	//数组相关
	testArray()
	//切片相关
	sliceTest()
}

// 切片相关测试
func sliceTest() {

}

// 测试数组
func testArray() {
	// 声明，即初始化，默认值为0
	var ary [8]int
	for i := 0; i < len(ary); i++ {
		fmt.Printf("%d is %v \n", i, ary[i])
	}
}

// 打印格式相关
func printSomething() {
	// %v
	var name string = "中国"
	fmt.Printf("%v\n", name)
	fmt.Printf("%+v\n", name)
	// 打印一个学生信息
	student := test1.Student{Name: "小米", Age: 8}
	fmt.Printf("student is %v\n", student)
	fmt.Printf("student is %+v\n", student)
	fmt.Printf("student is %#v\n", student)

	// 打印整形
	num := 1000
	fmt.Printf("num is [%d]\n", num)
	fmt.Printf("num is [%8d]\n", num)
	// 八进制
	fmt.Printf("num is [%o]\n", num)
	// %q
	fmt.Printf("num is %d,[%q]\n", 0x41, 0x41)
}

// 字符关键字 rune 测试
func runeTest() {
	var str string = "Abc啊喔鹅"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v(%c)", str[i], str[i])
	}
	fmt.Println()
	for _, s := range str {
		fmt.Printf("%v(%c,%X, %T, %b)", s, s, s, s, s)
	}
	fmt.Println()
}

// 初始化函数
func init() {
	fmt.Println("1.初始化函数")
}
