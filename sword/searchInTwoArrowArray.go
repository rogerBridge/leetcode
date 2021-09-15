package sword

func FindInMatrix1(target int, matrix [][]int) bool {
	for i:=0; i<len(matrix); i++ {
		if matrix[i][len(matrix[0])-1] >= target {
			for j:=0; j<len(matrix[0])-1; j++ {
				if target == matrix[i][j] {
					return true
				}
			}
		}
	}
	return false
}