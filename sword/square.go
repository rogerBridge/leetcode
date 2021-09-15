package sword

import (
	"fmt"
	"math"
)

func IfSquareOfTwo(nums int) bool {
	i := 0
	for {
		if int(math.Pow(2, float64(i))) < nums {
			i++
			continue
		}else if int(math.Pow(2, float64(i))) > nums {
			return false
		}
		fmt.Println("i: ", i)
		return true
	}
}

func IfSquareOfTwo1(nums int) bool {
	switch nums & (nums-1) {
	case 0:
		return true
	default:
		return false
	}
}