package main

import (
	"fmt"
	"sort"
)

func comp(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	for i := 0; i < 3; i++ {
		if a[i] != b[i] {
			return true
		}
	}
	return false
}

func cal(tol [][]int, pos []int) bool {
	ln := len(tol)
	if pos == nil || ln == 0{
		return true
	}
	for i := 0; i < ln; i++ {
		if comp(tol[i], pos) {
			return true
		}
	}
	return false
}

func threeSum(nums []int) [][]int {
	ans := make([][]int,0)
	if len(nums) < 4 {
		return make([][]int, 1)
	}
	ln := len(nums)
	pos := make([][]int, 10000)
	cnt := 0
	for i := 0; i < ln; i++ {
		for j := i+1; j < ln; j++ {
			for k := j+1; k < ln; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					a := []int{nums[i], nums[j], nums[k]}
					pos[cnt] = a
					cnt++
				}
			}
		}
	}
	for i:=0;i<cnt;i++{
		if cal(ans,pos[i]){
			ans = append(ans, pos[i])
		}
	}
	return ans
}

func RightthreeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue//To prevent the repeat
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return results
}
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
