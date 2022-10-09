package main

/* 使用goroutine和channel实现一个计算int64随机数个位数的程序
1、开启一个goroutine循环生成int64类型的随机数，发送到jobChan
2、开启24个goroutine从jobChan中取出随机数，计算各位数的和，将结果发送到resultChan
3、主goroutine从resultChan取出结果斌并打印到终端输出
*/

type Job struct{
	x int64
}

type Result struct{
	Job
	result int64
}

var jobChan = make(chan *Job,100)
var resultChan = make(chan *Result,100)

func woker(wk <-chan *Job )

func main() {
	
}