package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var sg sync.WaitGroup

//MainRoutine 主携程推出后，其他任务不再执行
func MainRoutine() {
	go func() {
		i := 0
		for {
			fmt.Printf("MainRoutine go routine i=====%v \n", i)
			time.Sleep(time.Second)
			i++
		}
	}()

	j := 1
	for {
		fmt.Printf("MainRoutine go routine j=====%v \n", j)
		j++
		if j == 3 {
			break
		}
		time.Sleep(time.Second)
	}
}
func Level1() {
	go Level2()
	fmt.Printf("This is Level1 method ! \n")
	for i := 1; i < 10; i++ {
		sg.Add(1)
		go Level3(i)
	}
	sg.Wait()
}

func Level2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("This is Level2 method ! %v times\n", i)
		time.Sleep(time.Second)
	}
}

func Level3(c int) {
	fmt.Printf("This is Level3 method ! count=%v \n", c)
	time.Sleep(time.Second * 2)
	sg.Done()
}
