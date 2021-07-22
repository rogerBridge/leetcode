package sword

import "log"

//func Fibonacci(n int) int {
//	// 求解: Fibonacci数列的第几项
//	if n == 0 {
//		return 0
//	}
//	if n == 1 {
//		return 1
//	}
//	return Fibonacci(n-2) + Fibonacci(n-1)
//}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	resultList := []int{0, 1}
	for i:=2; i<=n; i++ {
		resultList = append(resultList, resultList[i-2] + resultList[i-1])
	}
	log.Println(resultList)
	return resultList[len(resultList)-1]
}
