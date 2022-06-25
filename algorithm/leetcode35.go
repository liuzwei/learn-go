package algorithm

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//请必须使用时间复杂度为 O(log n) 的算法。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/search-insert-position
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func SearchInsert(nums []int, target int) int {
	// 二分查找，如果找到相同的位置则返回，
	left := 0
	right := len(nums) - 1

	for left <= right {
		// 中间位置
		center := (right-left)/2 + left
		if nums[center] == target {
			return center
		} else if target > nums[center] {
			left = center + 1
		} else {
			right = center - 1
		}
	}
	return left
}