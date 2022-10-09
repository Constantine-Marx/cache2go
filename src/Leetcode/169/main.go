package main

import "fmt"

//归并排序
func merge(nums []int, l, r, mid int) {
	i := l
	j := mid + 1
	tmp := make([]int, len(nums))
	p := 0

	for i < mid && j < r {
		if nums[i] < nums[j] {
			tmp[p] = nums[i]
			i++
			p++
		} else {
			tmp[p] = nums[j]
			j++
			p++
		}
	}
	for i <= mid {
		tmp[p] = nums[i]
		i++
		p++
	}
	for j <= r {
		tmp[p] = nums[j]
		j++
		p++
	}
	fmt.Println(nums)
	for t := 0; t < len(nums); t++ {
		nums[t] = tmp[t]
	}
}

func sort(nums []int, l,r int){
	if l >= r {
		return
	}

	mid := l + (r - l) / 2
	sort(nums,l,mid)
	sort(nums,mid+1,r)
	merge(nums,l,r,mid)

}

func test(nums []int){
	nums[1] = 23
}

func test1(nums []int){
	test(nums)
}

func main() {
	var nums = []int{2,3,5,1,34,12,42,12}
	sort(nums,0,len(nums)-1)
	// test1(nums)
	fmt.Println(nums)
}
