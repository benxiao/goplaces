package main

type Rectangle struct{
	Name string
	Width, Height float64
}

func(r Rectangle) Area() float64 {
	return r.Height * r.Width
}


func main() {



}


