package timer

import (
	"fmt"
	"time"
)

// UseTimer timer基本用法
func UseTimer() {
	timer := time.NewTimer(time.Second * 5)
	fmt.Printf("time now is t1=%v \n", time.Now())
	// 阻塞 timer的Duration
	t2 := <-timer.C
	fmt.Printf("time now is t2=%v \n", t2)
}

// CheckTimerResponseOnce 验证timer只能相应一次
func CheckTimerResponseOnce() {
	timer := time.NewTimer(time.Second)
	for {
		t := <-timer.C
		fmt.Printf("t is %v \n", t)
	}
}

// DelayTimer timer的延时功能
func DelayTimer() {
	fmt.Printf("开始：t=%v \n", time.Now())
	timer := time.NewTimer(time.Second * 3)
	fmt.Printf("3s时间到, t=%v \n", <-timer.C)
	i := 0
	for i < 5 {
		timer.Reset(time.Second * 3)
		fmt.Printf("3s时间到 again ! t=%v \n", <-timer.C)
		i++
	}
}

// StopTimer 定时器停止
func StopTimer() {
	timer := time.NewTimer(time.Second * 2)
	go func() {
		for {
			timer.Reset(time.Second * 2)
			fmt.Printf("2s时间到 again ! t=%v \n", <-timer.C)
		}
	}()
	time.Sleep(10 * time.Second)
	b := timer.Stop()
	if b {
		fmt.Printf("timer has stoped ! \n")
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("over ! \n")
}

func closeTimer(timer *time.Timer) bool {
	return timer.Stop()
}
