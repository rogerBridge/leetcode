package double_point

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//
//
//
//说明:
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	r := make([]int, 0, m+n)
	i, j := 0, 0
	for ; i < m && j < n; {
		if nums1[i] < nums2[j] {
			r = append(r, nums1[i])
			i++
		} else {
			r = append(r, nums2[j])
			j++
		}
	}
	if i<m {
		r = append(r, nums1[i:]...)
	}
	if j<m {
		r = append(r, nums2[j:]...)
	}
	return r
}