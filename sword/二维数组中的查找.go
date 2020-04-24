package sword

//给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。给定一个数，判断这个数是否在该二维数组中
//// 错误的代码
//func Find(target int, matrix [][]int) bool {
//	if len(matrix) == 0 || len(matrix[0]) == 0 || matrix[0][0] > target || matrix[len(matrix)-1][len(matrix[0])-1] < target{
//		return false
//	}
//	for i:=0; i<len(matrix)-1; i++ {
//		if matrix[i][0] == target {
//			return true
//		} else if matrix[i][0] < target && matrix[i+1][0] > target {
//			for j:=0; j<len(matrix[i]); j++ {
//				if target == matrix[i][j] {
//					return true
//				}
//			}
//		} else {
//			for j:=0; j<len(matrix[0]); j++ {
//				if target == matrix[len(matrix)-1][j]{
//					return true
//				}
//			}
//		}
//	}
//	return false
//}

func Find(target int, matrix [][]int) bool {
	// 从matrix[0]的最右边开始遍历
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for i:=len(matrix)-1; i>-1; i-- {
		if matrix[i][0] < target {
			for j:=0; j<len(matrix[0]); j++ {
				if matrix[i][j] == target{
					return true
				}
			}
		}
	}
	return false
}
