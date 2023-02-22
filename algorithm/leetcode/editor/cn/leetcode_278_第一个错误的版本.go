package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l := 0
	r := n
	for l <= r {
		m := l + (r-l)/2
		if isBadVersion(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

//leetcode submit region end(Prohibit modification and deletion)

// test self
func isBadVersion(m int) bool {

	return false
}
