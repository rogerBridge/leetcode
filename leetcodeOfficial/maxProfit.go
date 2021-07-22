package leetcodeofficial

import "fmt"

func MaxProfit(priceList []int) int {
	var profit int = 0
	priceListLen := len(priceList)
	for i := 1; i < priceListLen; i++ {
		if priceList[i] > priceList[i-1] {
			profit += priceList[i] - priceList[i-1]
		}
	}
	fmt.Printf("I had found the max profit: %d\n", profit)
	return profit
}
