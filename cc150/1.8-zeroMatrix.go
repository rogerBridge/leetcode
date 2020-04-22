package cc150

// 编写一种算法,若 M × N 矩阵中某个元素为 0,则将其所在的行与列清零。
// 先查找到值为0的元素的下标, 然后遍历, 将i=0, j=0的元素赋值为0

func zeroMatrix(s [][]int) [][]int {
	// 首先, 创造一个slicd, 存储值为0的元素的下标
	z := make([][]int, 0)
	for i:=0; i<len(s); i++ {
		for j:=0; j<len(s[0]); j++ {
			z = append(z, []int{i, j})
		}
	}
	// 然后, 遍历z, 将i, j所在的行, 列的所有值赋值为0
	for _, v := range z {
		x := v[0]
		for j:=0; j<len(s[x]); j++ {
			s[x][j] = 0
		}
		y := v[1]
		for i:=0; i<len(s); i++ {
			s[i][y] = 0
		}
	}
	return s
}