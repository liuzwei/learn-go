package util

// BubbleSort
func BubbleSort(nums []int) []int {
	if nums == nil || len(nums) == 1 {
		return nums
	}
	for l := len(nums) - 1; l > 0; l-- {
		// compare two element
		for i := 0; i < l; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums
}

// InsertionSort
func InsertionSort(nums []int) []int {
	for i := range nums {
		preIndex := i - 1
		currentValue := nums[i]

		for preIndex >= 0 && nums[preIndex] > currentValue {
			//交换
			nums[preIndex+1] = nums[preIndex]
			preIndex--
		}
		nums[preIndex+1] = currentValue
	}
	return nums
}

func HillSort(nums []int) []int {

	return nil
}

func FastSort(nums []int) []int {

	return nil
}
