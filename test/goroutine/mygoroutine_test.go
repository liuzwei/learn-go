package goroutine

import "testing"

func TestHello(t *testing.T) {
	Level1()
	t.Logf("This is test")
}

func TestMainRoutine(t *testing.T) {
	MainRoutine()
}
