package test1

import (
	"fmt"
	"sync"
	"time"
)

func hello(i int) {
	defer wg.Done()
	second, _ := time.ParseDuration(fmt.Sprintf("%ds", i))
	time.Sleep(second)
	fmt.Println("Hello Goroutine! ", i)
}

var wg sync.WaitGroup

func Goroutine() {
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go hello(i)
	//}
	//wg.Wait()
	//fmt.Println("main goroutine")

	go func() {
		i := 0
		for {
			i++
			fmt.Println("Hello goroutine ", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Println("main goroutine ", i)
		time.Sleep(time.Second)
		if i == 5 {
			break
		}
	}
}

func SomeChan() {
	var strChan chan string = make(chan string)
	go receiveStr(strChan)
	for i := 0; i < 10; i++ {
		strChan <- fmt.Sprintf("hello chan %d", i)
		time.Sleep(time.Second)
	}
	fmt.Println("send string success!")
}

func receiveStr(c chan string) {

}

func setZeroes(matrix [][]int) {
	// 先找出为0的元素
	zeros := [][]int{}
	for i, ary := range matrix {
		for j, item := range ary {
			if item == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
	}
	// 再重置为0
	for _, ary := range zeros {
		// 行置为0
		for x, _ := range matrix[ary[0]] {
			matrix[ary[0]][x] = 0
		}
		// 列置为0
		for y, _ := range matrix[0] {
			matrix[y][ary[1]] = 0
		}
	}
}

func longestCommonPrefix(strs []string) string {
	minLength := len(strs[0])
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}
	if minLength == 0 {
		return ""
	}
	result := ""
	for i := 0; i < minLength; i++ {
		temp := strs[0][i]
		for _, str := range strs {
			if temp != str[i] {
				return result
			}
		}
		result = result + string(temp)
	}
	return result
}
