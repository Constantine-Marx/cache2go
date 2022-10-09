package main

import (
	"fmt"
	"time"
)

func first(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		sum++
	}
}

func second(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n*n; j++ {
			sum++
		}
	}
}

func third(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n*n; j++ {
			sum++
		}
	}
}

func forth(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			sum++
		}
	}
}

func fifth(n int) {
	// sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i*i; j++ {
			// for k := 0; k < j; k++ {
			// 	sum++
			// }
		}
	}
}

func sixth(n int) {
	sum := 0
	for i := 1; i < n; i++ {
		for j := 1; j < i*i; j++ {
			if j%i == 0 {
				for k := 0; k < j; k++ {
					sum++
				}
			}
		}
	}
}

func test(n int) {
	start1 := time.Now()
	first(n)
	cost1 := time.Since(start1)
	fmt.Printf("N = %d,The (1) time cost=%v\n", n, cost1)

	start2 := time.Now()
	second(n)
	cost2 := time.Since(start2)
	fmt.Printf("N = %d,The (2) time cost=%v\n", n, cost2)

	start3 := time.Now()
	third(n)
	cost3 := time.Since(start3)
	fmt.Printf("N = %d,The (3) time cost=%v\n", n, cost3)

	start4 := time.Now()
	forth(n)
	cost4 := time.Since(start4)
	fmt.Printf("N = %d,The (4) time cost=%v\n", n, cost4)

	start5 := time.Now()
	fifth(n)
	cost5 := time.Since(start5)
	fmt.Printf("N = %d,The (5) time cost=%v\n",n,cost5)

	start6 := time.Now()
	sixth(n)
	cost6 := time.Since(start6)
	fmt.Printf("N = %d,The (6) time cost=%v\n", n, cost6)
}

func main() {
	for i := 1;i<=1024; i *= 2 {
		test(i)
		fmt.Println("---------------")
	}
}
