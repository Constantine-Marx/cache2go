package main

import "fmt"

func generate(numRows int) [][]int {
	var nums [][]int
	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		nums = append(nums, temp)
		nums[i][0] = 1
		nums[i][i] = 1
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			nums[i][j]  = nums[i-1][j-1] + nums[i-1][j]
		}
	}
	for i := 0;i< numRows;i++{
		fmt.Println(nums[i])
	}
	return nums
}

func main() {
	generate(10)
}
