package sword

func Bubble(n []int) []int {
	for i := 0; i < len(n)-1; i++ {
		isSorted := true
		for j := 0; j < len(n)-1-i; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
	return n
}

func Cocktail(n []int) []int {
	return n
}
