package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	rowIndex++
	sourceArr := make([]int, rowIndex)
	sourceArr[0] = 1
	sourceArr[1] = 1
	for i := 2; i < rowIndex; i++ {
		tempArr := make([]int,rowIndex)
		tempArr[0] = 1
		tempArr[i] = 1
		for j := 1; j < i; j++ {
			tempArr[j] = sourceArr[j-1] + sourceArr[j]
		}
		sourceArr = tempArr
	}
	return sourceArr
}

func main() {
	fmt.Println(getRow(1))
}
