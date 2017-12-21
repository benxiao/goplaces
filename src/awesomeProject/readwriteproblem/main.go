package main

import (
	"fmt"
	"time"
)

func writer(buffer chan int32){
	for _, ch := range "it just works!"{
		buffer <- ch
		time.Sleep(100 * time.Millisecond)
	}
	close(buffer)
}

func reader(buffer chan int32, done chan bool){
	for ch := range buffer{
		fmt.Printf("%c", ch)
		time.Sleep(2000 * time.Millisecond)
	}
	done <- true
}

func main() {
	done := make(chan bool)
	buffer := make(chan int32, 10)
	go writer(buffer)
	go reader(buffer, done)
	<-done
}
