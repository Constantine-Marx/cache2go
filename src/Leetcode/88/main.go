package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	j := 0
	for i := 0; i < m; i++ {
		if nums1[i] <= nums2[j] {
			if nums1[i+1] >= nums2[j] {
				temp := nums1[:i+1]
				temp1 := nums1[i+1 : m-n]
				temp = append(temp, nums2[j])
				temp = append(temp, temp1...)
				j++
				// nums1 = temp
			}
		}
	}
	return nums1
}

func main() {
	nums1 := [6]int{1, 2, 3, 0, 0, 0}
	nums2 := [3]int{2, 5, 6}
	fmt.Println(merge(nums1[:],6,nums2[:],3))
}
