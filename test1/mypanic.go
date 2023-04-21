package test1

import "fmt"

func panicUse1() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("has a panic: %v\n", r)
		}
	}()

	panic("I'm not sure this is a panic!")
}

func sendMessageToCloseChannel() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("has a panic: %v\n", r)
		}
	}()

	ch := make(chan int)
	close(ch)

	ch <- 1
}

func panicOfDefer() {
	defer func() {
		// 只能捕获到“second panic”，捕获不到 First panic
		if r := recover(); r != nil {
			fmt.Printf("defer1 has a panic: %v\n", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("defer2 has a panic: %v\n", r)
			panic("second panic")
		}
	}()

	panic("First panic")
}

func protectCode(x, y int) int {
	var z int

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("has a panic: %v\n", r)
				z = 10
			}
		}()
		panic("has error")
		z = x + y
	}()

	return z
}
