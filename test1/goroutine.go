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
