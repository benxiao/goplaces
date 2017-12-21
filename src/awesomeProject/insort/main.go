package main

import "fmt"

func isort(array []int, comp func(int, int) bool){
	for i:=1; i!=len(array); i++ {
		j := i
		val := array[i]
		for j > 0 && comp(val, array[j-1]){
			array[j] = array[j-1]
			j -= 1
		}
		array[j] = val
	}
}





func main() {
	array := []int{5,1,2,3,4,5,10}
	isort(array, func(x int, y int) bool {return x < y})
	fmt.Println(array)
}
