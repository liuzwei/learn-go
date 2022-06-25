package main

import (
	"fmt"

	"learn-go/algorithm"
	"learn-go/test1"
)

// 声明一个加法函数
var add = func(a int, b int) int {
	return a + b
}

func Add(a int, b int) int {
	return a + b
}

func main() {
	nums := []int{3}
	r := algorithm.RemoveElement(nums, 3)
	fmt.Println(r)
}
func main3() {
	//test1.PrintIntAddress()
	//test1.PrintStringAddress()
	ary := []int{1, 2, -2147483648}
	algorithm.ThirdMax(ary)

	c := add(1, 2)
	fmt.Println(c)
	d := Add(1, 2)
	fmt.Println(d)

	var shape test1.Shape = test1.Rectangle{A: 1.0, B: 2.0}
	rect, ok := shape.(test1.Rectangle)
	if ok {
		fmt.Println("shape perimeter ", rect.Perimeter(), "area is ", rect.Area())
	} else {
		fmt.Println("shape.(test1.Rectangle) not ok!")
	}
	var shape2 test1.Shape = test1.Triangle{A: 3, B: 4, C: 5}
	fmt.Println("shape2 perimeter ", shape2.Perimeter(), "area is ", shape2.Area())

	table := map[int]int{}
	table[1] = 0
}

// 主入口函数
func main2() {
	fmt.Println("2.主函数")
	// # command-line-arguments
	//.\main.go:11:2: undefined: test1
	//test1.SayHello()
	//var s int = int(math.Max(10, 30))

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

	// 数组指针
	ptrArray()
	// 切片拷贝
	copySlice()

}

// 切片拷贝
func copySlice() {
	var a = []int{1, 2, 4}
	var s = make([]int, 6)
	copy(s, a)
	fmt.Printf("a array %v, type:%T, cap:%d, len:%d \n", a, a, cap(a), len(a))
	fmt.Printf("s slice %v, type:%T, cap:%d, len:%d \n", s, s, cap(s), len(s))
}

// 测试数组是值传递还是引用传递
func ptrArray() {
	a := [...]int{1, 2, 4}
	fmt.Printf("数组a，指针地址【%p】,数组信息：%v \n", &a, a)

	//printArray(a)
	printArrayQuote(&a)
}
func printArrayQuote(b *[3]int) {
	fmt.Printf("数组b，指针地址【%p】,数组信息：%v \n", b, *b)
	c := b
	fmt.Printf("数组c，指针地址【%p】,数组信息：%v \n", c, *c)
	var slice []int
	slice = append(slice, 1, 2, 4)
}

func printArray(a [3]int) {
	fmt.Printf("数组b，指针地址【%p】,数组信息：%v \n", &a, a)
	c := a
	fmt.Printf("数组c，指针地址【%p】,数组信息：%v \n", &c, c)
}

// 切片相关测试
func sliceTest() {
	// 声明一个10个长度的数组，并初始化
	var originArray [10]int
	for i := 0; i < len(originArray); i++ {
		originArray[i] = i + 1
	}

	// 获得一个切片
	slice1 := originArray[5:]
	slice2 := slice1[:]
	fmt.Printf("slice1 is %v \n", slice1)
	fmt.Printf("slice2 is %v \n", slice2)
	// 修改originArray数组，看slice1是否变化
	originArray[5] = 0
	fmt.Printf("change ? slice1 is %v \n", slice1)
	fmt.Printf("change ? slice2 is %v \n", slice2)
	// slice1追加一个元素,append操作会导致底层数组变化，发生数组的拷贝
	slice3 := append(slice1, 100)
	fmt.Printf("slice3 is %v \n", slice3)
	fmt.Printf("apeend ? slice1 is %v \n", slice1)
	originArray[5] = 6
	fmt.Printf("change ? slice3 is %v \n", slice3)
	fmt.Printf("change ? slice1 is %v \n", slice1)

	var slice4 []int = make([]int, 10)
	fmt.Printf("slice4 is %v \n", slice4)

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
