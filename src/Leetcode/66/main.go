package main

import "fmt"

func plusOne(digits []int) []int {
	ln := len(digits) - 1
	digits[ln]++
	temp := make([]int, ln+2)

	for i := ln; i > 0; i-- {
		if digits[i] >= 10 {
			digits[i-1] += digits[i] / 10
			digits[i] %= 10
		}
	}

	for i := 0; i <= ln; i++ {
		temp[i] = 0
	}
	for i := ln; i >= 0; i-- {
		if digits[i] >= 10 {
			temp[i] = digits[i] / 10
			digits[i] %= 10
		}
		temp[i+1] = digits[i]
		fmt.Println(temp)
		fmt.Println(digits)
		fmt.Println("**************")
	}

	if temp[0] == 0 {
		return temp[:]
	} else {
		return temp
	}
}

func main() {
	test := [5]int{9, 9, 9, 9, 9}
	fmt.Println(plusOne(test[:]))
}
