package main

import (
	"math"
)

func Reverse(x int) int {
	singleNumList := make([]int, 0)
	for {
		if x > 0 {
			singleNumList = append(singleNumList, x%10)
			x = x / 10
		} else {
			break
		}
	}
	//fmt.Println(singleNumList)
	sum := 0.0
	for idx := 0; idx < len(singleNumList); idx++ {
		sum += float64(singleNumList[idx]) * math.Pow10(len(singleNumList)-1-idx)
	}
	//fmt.Println(int(sum))
	if int(sum) < int(-math.Pow(2, 31)) || int(sum) > int(math.Pow(2,31))-1 {
		return 0
	} else {
		return int(sum)
	}
}

func reverse(x int) int {
		if x < 0 {
			return -1 * Reverse(-x)
		} else if x > 0 {
			return Reverse(x)
		} else {
			return 0
		}
}

func main() {
	reverse(1534236469)
	//fmt.Println(int(math.Pow(2, 31)))
}
