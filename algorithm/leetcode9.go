package algorithm

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//例如，121 是回文，而 123 不是。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/palindrome-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// IsPalindrome
func IsPalindrome(x int) bool {
	if (x < 0 || x%10 == 0) && x != 0 {
		return false
	}
	recoverNumber := 0
	for x > recoverNumber {
		recoverNumber = recoverNumber*10 + x%10
		x /= 10
	}

	return x == recoverNumber || recoverNumber/10 == x
}
