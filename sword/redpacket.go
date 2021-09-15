package sword

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Redpacket(amount int, userNum int) ([]int, error) {
	if amount < 1 || userNum < 1 || amount/userNum < 1 {
		return []int{}, fmt.Errorf("input invalid")
	}

	rand.Seed(time.Now().UnixNano())

	moneyArray := make([]int, userNum)
	index := 0
	var n int
	d := make(map[int]struct{})
	for index < userNum-1 {
		n = rand.Intn(amount)
		if _, ok := d[n]; ok || n == 0 {
			continue
		}
		d[n] = struct{}{}
		moneyArray[index] = n
		index++
	}
	moneyArray[userNum-1] = amount
	// sort moneyArray
	sort.Ints(moneyArray)

	for i := len(moneyArray) - 1; i > 0; i-- {
		moneyArray[i] = moneyArray[i] - moneyArray[i-1]
	}
	return moneyArray, nil
}
