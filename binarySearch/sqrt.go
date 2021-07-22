package binarySearch

// 计算一个数的平方根, 误差必须小于1e-5
//func Sqrt(x float64) float64 {
//	low := 0.0
//	high := x
//	guess := (low + high) / 2
//	for guess*guess-x <= 1e-5 || guess*guess-x >= -(1e-5) {
//		if guess*guess == x {
//			return guess
//		} else if guess*guess-x > 0 {
//			high = guess
//			guess = (low + high) / 2
//		} else {
//			low = guess
//			guess = (low + high) / 2
//		}
//	}
//	return guess
//}

func Sqrt(x float64) float64 {
	low := 0.0
	high := x
	guess := (low+high)/2
	for guess*guess-x<=-1e-5 || guess*guess-x>=1e-5 {
		if guess*guess == x {
			return guess
		}else if guess*guess>x {
			high = guess
			guess =(low+high)/2
		}else {
			low = guess
			guess = (low+high)/2
		}
	}
	return guess
}















