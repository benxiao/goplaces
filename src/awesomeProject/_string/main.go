package main

import (

	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Item struct {
	name string
	price int32
}



func open(file string)<-chan string{
	out := make(chan string)

	go func() {
		f, err := os.Open(file)

		defer close(out) // close the out channel, so the for loop terminates
		defer f.Close() //close the file pointer, so there is no resource leak

		if err != nil { // cannot open file successfully
			fmt.Println(err)
			return
		}
		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {

				break // encounter error while reading the file, leave!
			}
			out<-line
		}
	}()
	return out
}


func main() {
	fp := open("data/pccasegear.dat")

	for line := range fp {
		array := strings.Split(line, "||")
		if len(array) == 2 {
			name := array[0]
			s_price := strings.Trim(array[1], "$\n ")
			i_price, err := strconv.ParseInt(s_price, 10, 32)
			if err == nil {
				fmt.Println(name, i_price)
			}

		}
	}
}
