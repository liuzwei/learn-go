package test1

import "fmt"

func myFunc(x, y int) (int, error) {
	return x + y, nil
}

// define a constant
const c1 = 10

type User struct {
	age     int
	name    string
	address string
}

func printConstant() {
	var c2 int = 100
	// 常量取地址错误，字面量值没有地址
	//fmt.Printf("Constant c1 is %v", &c1)
	fmt.Printf("Constant c2 is %v", &c2)
}
