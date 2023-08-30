package channel

import "testing"

func TestGenerateNumber(t *testing.T) {
	go generateNumber(1000)
	go squarerNumber()
	logSquarer()
}
