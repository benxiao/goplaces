package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"io"
)


func open(name string, out chan<-string){
	f, err := os.Open("hello.txt")
	defer f.Close()
	if err != nil {
		log.Fatalln("cannot find file")
	}
	reader := bufio.NewReader(f)
	var line string
	for {
		line, err = reader.ReadString('\n')
		out<-line
		if err == io.EOF {break}
	}
	close(out)
}


func main() {
	out := make(chan string)
	go open("hello.txt", out)
	for line := range out {
		fmt.Print(line)
	}
}
