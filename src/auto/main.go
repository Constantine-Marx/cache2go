package main

import "fmt"

func main() {
	a := 0
	name := ""
	fmt.Println("请输入书籍名字:")
	fmt.Scan(&name)
	fmt.Println("请输入书籍最大页码:")
	fmt.Scan(&a)
	// fmt.Println(name,a)
	var nums[50]int
	fmt.Println(nums)
	cnt := 1
	for i := range nums {
		jg,_ := fmt.Scan("%d",&nums[i])
		fmt.Println(nums)
		if nums[i] >= a && jg != 1{
			break
		}
	}
	cnt = 1
	for{
		if nums[cnt+1] <= a{
		fmt.Printf("《%s》P%d - P%d\n",name,nums[cnt],nums[cnt+1])
		}else{
			break
		}
	}
}