package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type listfun struct {
}

func (l listfun) a(head *node)(sum int){
	temp := head
	sum = 0
	for temp != nil{
		sum++
	}
	return
}

func (l listfun) b(head *node){
	temp := head
	for temp != nil{
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func (l listfun) c(head *node,target int)bool{
	temp := head
	for temp != nil{
		if temp.value == target{
			return true
		}
	}
	return false
}

func (l listfun) d(head *node,x int){
	temp := head
	for temp.next != nil{
		temp = temp.next
	}
	p := new(node)
	p.value = x
	temp.next = p
}

func (l listfun) e(head *node,x int){
	temp := head
	for temp.next.value != x{
		temp = temp.next
	}
	temp.next = temp.next.next
}

func main() {
	head := new(node)
	head.value = 0
	head.next = nil
	temp := head
	for i := 1; i < 11; i++ {
		p := new(node)
		p.next = nil
		p.value = i
		temp.next = p
		temp = p

	}

	var l listfun
	l.a(head)
}
