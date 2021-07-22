package binarysearch

// FindInArray 如果x存在y中, 返回x的下标, 如果不存在, 返回-1
func FindInArray(x int, nums []int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		middle := (high + low) / 2
		if nums[middle] == x {
			return middle
		} else if nums[middle] > x {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}
