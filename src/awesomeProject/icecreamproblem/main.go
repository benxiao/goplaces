package main

import (
	"math/rand"
	"fmt"
	"time"
)


type IceCreamCone struct {

}

var requestedInspection = make(chan IceCreamCone)
var conePassedInspection = make(chan bool)
var receivedPayment = make(chan int)
var finishedBilling = make(chan bool)
var customerWalkAway = make(chan bool)

func manager(){
	for {
		<-requestedInspection
		inspect()
		f := rand.Float64()
		if f > 0.5{
			conePassedInspection <- true
		}else{
			conePassedInspection <- false;
		}
	}
}

func clerk(servedThisCone chan bool){
	for {
		requestedInspection <- makeCone()
		if <-conePassedInspection {break}
	}
	servedThisCone<-true
}

func customer(numOfCones int){
	servedOneCone := make(chan bool)
	for i:=0; i!=numOfCones; i++ {
		go clerk(servedOneCone)
	}
	//wait on the cones
	for i:=0; i!=numOfCones; i++ {
		<-servedOneCone
	}
	receivedPayment<-10
	payBill()

	<-finishedBilling
	customerWalkAway<-true
}

func cashier(){
	for{
		money := <-receivedPayment
		processPayment(money)
		finishedBilling<-true
	}
}

func makeCone() IceCreamCone{
	time.Sleep(time.Millisecond * 500)
	fmt.Println("clerk: make cone")
	return IceCreamCone{}
}

func inspect(){
	time.Sleep(time.Millisecond * 150)
	fmt.Println("manager: inspect cone")
}


func payBill(){
	time.Sleep(time.Millisecond * 400)
	fmt.Println("customer: pay bill")
}

func processPayment(amount int){
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("cashier: received payment")
}

func main() {
	rand.Seed(0)
	start := time.Now()
	go manager()
	n := 10




	for i:=0; i!=n; i++ {
		go customer(rand.Intn(4)+1)
	}
	go cashier()
	for i:=0; i!=n; i++ {
		<-customerWalkAway
	}
	fmt.Println(time.Since(start))
}
