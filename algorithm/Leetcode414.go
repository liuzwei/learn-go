package algorithm

import "math"

func ThirdMax(nums []int) int {
	const MIN_VALUE = -1<<31 - 1
	var max1, max2, max3 = MIN_VALUE, MIN_VALUE, MIN_VALUE
	for _, item := range nums {
		if item > max1 {
			max1, max2, max3 = item, max1, max2
		} else if item > max2 && item < max1 {
			max2, max3 = item, max2
		} else if item > max3 && item < max2 {
			max3 = item
		}
	}
	if max3 == MIN_VALUE {
		return max1
	}
	return max3
}

func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > a {
			a, b, c = num, a, b
		} else if a > num && num > b {
			b, c = num, b
		} else if b > num && num > c {
			c = num
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c
}
