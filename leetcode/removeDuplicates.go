package leetcode

func RemoveDuplicates(nums []int) int {
	slow :=  0
	for fast := 1; fast<len(nums); fast++ {
		if nums[fast] > nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow+1
}