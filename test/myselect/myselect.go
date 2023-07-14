package myselect

import (
	"fmt"
	"sync"
	"time"
)

func SelectChannel() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case s := <-ch1:
				fmt.Printf("ch1 read value is %v \n", s)
			case i := <-ch2:
				fmt.Printf("ch2 read value is %v \n", i)
			default:
				fmt.Printf("sleep 500 millsecond \n")
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	var wg = sync.WaitGroup{}
	// send message to ch1
	wg.Add(1)
	go func() {
		for i := 0; i < 26; i++ {
			ch1 <- string(rune('A' + i))
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	// send message to ch2
	wg.Add(1)
	go func() {
		for i := 0; i < 26; i++ {
			ch2 <- i
			time.Sleep(time.Second * 2)
		}
		wg.Done()
	}()

	wg.Wait()
}
