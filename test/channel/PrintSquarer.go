package channel

import "fmt"

var numberChan = make(chan int)
var squarerChan = make(chan int)

// generateNumber generate numbers from size
func generateNumber(size int) {
	for i := 0; i < size; i++ {
		// 发送产生的数字到chan中
		numberChan <- i
	}
	close(numberChan)
}

func squarerNumber() {
	for n := range numberChan {
		// 计算平方，得到结果后发送到squarerChan
		squarerChan <- n * n
	}
	close(squarerChan)
}

func logSquarer() {
	for x := range squarerChan {
		fmt.Println(x)
	}
}
