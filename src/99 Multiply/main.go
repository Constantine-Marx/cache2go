package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	calMuti(10)
	// fmt.Printf("%20d",123);
}

func calMuti(levelNum int) {
	rand.Seed(time.Now().Unix())
	y := rand.Intn(levelNum)
	randOpr := true
	for i := 0; i < 9; i++ {
		for j := 0; j < i; j++ {
			y = rand.Intn(levelNum)
			printMuti("src/99 Multiply/pro.txt", "src/99 Multiply/ans.txt", j+1, i+1, (i+1)*(j+1), y, false, randOpr)
		}
		printMuti("src/99 Multiply/pro.txt", "src/99 Multiply/ans.txt", i+1, i+1, (i+1)*(i+1), y, true, randOpr)
	}
}

func printMuti(proName, ansName string, x, y, z, randNum int, endl, randOpr bool) { //randopr -- rand or no rand
	
	proPath, proErr := os.OpenFile(proName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	ansPath, ansErr := os.OpenFile(ansName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if proErr != nil {
		fmt.Println("Pro file Open Fail!")
		return
	}
	
	if ansErr != nil {
		fmt.Println("Ans file Open Fail!")
		return
	}

	if randOpr { //随机挖空模式
		if randNum < 5 {
			if endl {
				fmt.Fprintf(proPath, "%d × %d = %-2s\n", x, y,"__")
			} else {
				fmt.Fprintf(proPath, "%d × %d = %-2s ", x, y,"__")
			}
		} else {
			if endl {
				fmt.Fprintf(proPath, "%d × %d = %-2d\n", x, y, z)
			} else {
				fmt.Fprintf(proPath, "%d × %d = %-2d ", x, y, z)
			}
		}
	} else { //非随机模式
		if endl {
			fmt.Fprintf(proPath, "%d × %d = %-2s\n", x, y,"__")
		} else {
			fmt.Fprintf(proPath, "%d × %d = %-2s ", x, y,"__")
		}
	}

	//输出答案文件
	if endl {
		fmt.Fprintf(ansPath, "%d × %d = %-2d\n", x, y, z)
	} else {
		fmt.Fprintf(ansPath, "%d × %d = %-2d ", x, y, z)
	}
}
