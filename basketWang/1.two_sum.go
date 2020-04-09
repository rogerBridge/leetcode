package main

// [1, 2, 9, 7, 5]
// 两数之和
func TwoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		wantNum := target - nums[i]
		if _, ok := dict[wantNum]; !ok {
			dict[nums[i]] = i
			continue
		}
		return []int{i, dict[wantNum]}
	}
	return []int{}
}
