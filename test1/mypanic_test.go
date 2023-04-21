package test1

import (
	"fmt"
	"testing"
)

func TestPanicUse1(t *testing.T) {
	panicUse1()
}

func TestSendMessageToCloseChannel(t *testing.T) {
	sendMessageToCloseChannel()
}

func Test_panicOfDefer(t *testing.T) {
	panicOfDefer()
}

func Test_protectCode(t *testing.T) {
	fmt.Printf("x + y=%v\n", protectCode(1, 2))
}
