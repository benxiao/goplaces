package main

import "fmt"

type Point struct {x, y int}

func (p *Point) SetX(x int){ p.x = x}
func (p *Point) SetY(y int){ p.y = y}


func main() {

	p := Point{10, 29}
	fmt.Println(p)
	p.SetX(20)
	fmt.Println(p)
}
