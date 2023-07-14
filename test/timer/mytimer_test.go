package timer

import "testing"

func TestUseTimer(t *testing.T) {
	UseTimer()
}

func TestCheckTimerResponseOnce(t *testing.T) {
	go CheckTimerResponseOnce()
}

func TestDelayTimer(t *testing.T) {
	DelayTimer()
}

func TestStopTimer(t *testing.T) {
	StopTimer()
}
