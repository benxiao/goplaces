package main

import "fmt"

type Node struct {
	next *Node
	value int
}


type LinkedList struct {
	data chan int
	length int
}

func (ll *LinkedList)New(capactiy int){
	ll.data = make(chan int, capactiy)
	ll.length = 0
}

func (ll *LinkedList)add(val int){
	ll.data<-val
	ll.length++
}

func (ll *LinkedList)pop() int {
	ll.length--
	return <-ll.data
}

func (ll *LinkedList) empty() bool {
	return ll.length == 0
}

func main() {
	ll := LinkedList{}
	ll.New(100)
	ll.add(10)
	ll.add(20)
	ll.add(30)
	for !ll.empty(){
		fmt.Println(ll.pop())
	}




}
