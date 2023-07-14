package channel

import (
	"fmt"
	"time"
)

func receiveMessage(ch chan string) {
	for {
		fmt.Printf("receive message %v \n", <-ch)
	}
}

func SendMessage(msg string) {
	var ch = make(chan string)
	go receiveMessage(ch)

	// 向channel中发送消息
	for i := 0; i < 5; i++ {
		ch <- msg
		time.Sleep(time.Second)
	}
}

func JudgeChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			// 向ch1中发送100个数
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			// i的平方发送到ch2
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 可以两个携程同时读取channel中的数据
	go func() {
		for i := range ch2 {
			fmt.Println("go func ", i)
		}
	}()

	go func() {
		for {
			j, ok := <-ch2
			if ok {
				fmt.Printf("main get %v \n", j)
			} else {
				fmt.Printf("main get None! \n")
				break
			}
		}
	}()

}
