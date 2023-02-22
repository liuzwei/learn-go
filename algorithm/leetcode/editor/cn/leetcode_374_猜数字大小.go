package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if guess(mid) <= 0 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func guess(num int) int {
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
