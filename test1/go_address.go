package test1

import (
	"fmt"
	"unsafe"
)

func PrintIntAddress() {
	var num1 = 100
	var num2 = 100
	var num3 = num1

	fmt.Println(num1, &num1)
	fmt.Println(num2, &num2)
	fmt.Println(num3, &num3)

	var str1 = "go1"
	var str2 = "go2"
	var str3 = str1

	fmt.Println(str1, &str1)
	fmt.Println(str2, &str2)
	fmt.Println(str3, &str3)

	var num5 int8 = 100

	// 打印占用内存大小
	fmt.Println("num1内存占用大小", unsafe.Sizeof(num1))
	fmt.Println("num5内存占用大小", unsafe.Sizeof(num5))
}

func PrintStringAddress() {
	var str1 = "go1"
	var str2 = "go2"
	var str3 = str1

	fmt.Println(str1, &str1)
	fmt.Println(str2, &str2)
	fmt.Println(str3, &str3)
}
