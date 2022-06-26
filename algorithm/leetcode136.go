package algorithm

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/single-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// SingleNumber 因为只有一个不同的数，那么剩余的数肯定可以两两配对，
// 异或运算：有两个值a,b（值只能是0或1）  如果a，b值相同,异或结果为0，如果a,b值不同，异或结果为1
// 根据异或运算的特点，所有值异或完，结果即为只出现一次的元素
func SingleNumber(nums []int) int {
	a := "hello"
	b := "world"
	a, b = b, a

	result := 0
	for _, item := range nums {
		result ^= item
	}
	return result
}
