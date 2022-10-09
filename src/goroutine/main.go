package main

import (
	"fmt"
	"sync"
)

var x = 0
var wgp sync.WaitGroup

//add x plus 1
func add() {
	for i := 0; i < 50000; i++ {
		x++
	}
	wgp.Done()
}

func mgr() {
	x++
}

func main() {
	wgp.Add(2)
	go add()
	go add()
	wgp.Wait()
	fmt.Println(x)
}
