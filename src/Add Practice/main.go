package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 通用的文件打开函数(综合和 Create 和 Open的作用)
// OpenFile第二个参数 flag 有如下可选项
//    O_RDONLY  文件以只读模式打开
//	  O_WRONLY  文件以只写模式打开
//	  O_RDWR   文件以读写模式打开
//	  O_APPEND 追加写入
//	  O_CREATE 文件不存在时创建
//	  O_EXCL   和 O_CREATE 配合使用,创建的文件必须不存在
//	  O_SYNC   开启同步 I/O
//	  O_TRUNC  打开时截断常规可写文件

func main() {
	rand.Seed(time.Now().Unix())
	val := 1
	cnt := 0
	for i := 0; i < 50; {
		x := rand.Intn(10)
		y := rand.Intn(10)
		if (x+y)-10 == val {
			if cnt < 10 {
				wriInt(x, y, x+y, 1, "./src/Add Practice/pro.txt", "src\\Add Practice\\ans.txt")
				cnt++
				i++
				fmt.Printf("%dth write\n", i)
			} else {
				cnt = 0
				val++
				continue
			}
		} else {
			continue
		}
	}

	tempPath,_ := os.OpenFile("./src/Add Practice/pro.txt",os.O_RDWR|os.O_APPEND,0644)
	fmt.Fprintln(tempPath)
	fmt.Fprintln(tempPath)
	fmt.Fprintln(tempPath)
	tempPath,_ = os.OpenFile("src\\Add Practice\\ans.txt",os.O_RDWR|os.O_APPEND,0644)
	fmt.Fprintln(tempPath)
	fmt.Fprintln(tempPath)
	fmt.Fprintln(tempPath)
	for i := 0; i < 50; {
		x := rand.Intn(10)
		y := rand.Intn(10)
		if x==0||y==0{
			continue
		}
		if (x+y)%10 <= 5 && x+y>10 {
			wriInt(x, y, x+y, 1, "./src/Add Practice/pro.txt", "src\\Add Practice\\ans.txt")
			// fmt.Println(x," ",y," ",x+y)
			i++
		}
	}

}

func wriInt(x, y, z, opr int, proName, ansName string) { //opr --- 1-add 0-sub
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
	// fmt.Println("Open Success!")
	if opr == 1 { //add
		_, proErr := fmt.Fprintf(proPath, "%d + %d = __\n", x, y)
		_, ansErr := fmt.Fprintf(ansPath, "%d + %d = %d\n", x, y, z)
		if proErr != nil {
			fmt.Println("pro Write Fail!")
			return
		}
		if ansErr != nil {
			fmt.Println("ans Write Fail!")
			return
		}
	} else if opr == 0 {
		_, proErr := fmt.Fprintf(proPath, "%d - %d = __\n", x, y)
		_, ansErr := fmt.Fprintf(ansPath, "%d - %d = %d\n", x, y, z)
		if proErr != nil {
			fmt.Println("pro Write Fail!")
			return
		}
		if ansErr != nil {
			fmt.Println("ans Write Fail!")
			return
		}
	} else {
		fmt.Println("opr is undefined!")
		return
	}
	// fmt.Println("Write Success!")
}
