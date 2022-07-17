package swordoffer

// 给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算滑动窗口里所有数字的平均值。
//
//实现 MovingAverage 类：
//
//MovingAverage(int size) 用窗口大小 size 初始化对象。
//double next(int val) 成员函数 next 每次调用的时候都会往滑动窗口增加一个整数，请计算并返回数据流中最后 size 个值的移动平均值，即滑动窗口里所有数字的平均值。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/qIsx9U
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MovingAverage struct {
	queue    []int
	capacity int
	sum      int
}

/** Initialize your data structure here. */
func MovingAverageConstructor(size int) MovingAverage {
	return MovingAverage{
		capacity: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) == this.capacity {
		this.sum -= this.queue[0]
		this.queue = this.queue[1:]
	}
	this.queue = append(this.queue, val)
	this.sum += val
	return float64(this.sum) / float64(len(this.queue))
}
