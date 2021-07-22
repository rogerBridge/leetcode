package sword

// 在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func FindInMatrix(target int, matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for i:=0; i<len(matrix); i++ {
		// 每一行的最后一位(最大的那个数值) 如果大于等于我们给定的数值, 那么要查找的数值就有可能在那一行
		if matrix[i][len(matrix[i])-1] >= target {
			for j :=0; j<len(matrix[i]); j++{
				if matrix[i][j] == target {
					return true
				}
			}
		}
	}
	return false
}
