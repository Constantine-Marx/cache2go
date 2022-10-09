package main

import (
	"fmt"
)

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	Curr_profit := 0
	minprice := prices[0]
	for _,price := range prices{
		minprice = max(minprice,price)
		Curr_profit = max(Curr_profit,price - minprice)
	}
	return Curr_profit
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(nums))
}


