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
	for {
		str := <-c
		fmt.Println("接收字符串", str)
	}

}
