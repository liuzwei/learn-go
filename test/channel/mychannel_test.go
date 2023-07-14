package channel

import "testing"

func TestSendMessage(t *testing.T) {
	SendMessage("Hello channel!")
	//SendMessage("Hello Golang!")
	//SendMessage("Hello boy!")
}

func TestJudgeChannel(t *testing.T) {
	JudgeChannel()
}
