package main

import "fmt"

func main() {
	animals := []Animal{Dog{}, Cat{}}
	for _, a := range animals{
		fmt.Println(a.Speak())
	}
}


type Animal interface {
	Speak() string
}


type Dog struct {}
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {}
func (c Cat) Speak() string {
	return "Meow!"
}

