package main

import (
	"sync"
	"bufio"
	"os"
	"log"
	"net"
	"strings"
	"fmt"
)

type lookup struct {
	name string
	err error
	result bool
}


var wg sync.WaitGroup
var _in = make(chan lookup)
var _out = make(chan lookup)

func main(){
	wg.Add(1);
	// read from Stdin
	go func(){
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			_in <- lookup{name:s.Text()}
		}
		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(_in)
		wg.Done()
	}()

	// processing with goroutines

	for i:=0; i!=1000; i++ {
		wg.Add(1)
		go func(){
			for l := range _in {
				nss, err := net.LookupNS(l.name)
				if err != nil {
					l.err = err
				} else {
					for _, ns := range nss {
						if strings.HasSuffix(ns.Host, ".ns.cloudflare.com"){
							l.result = true
							break
						}
					}
				}
				_out <- l
			}
			wg.Done()
		}()
	}

	// print out the result to stdout
	wg.Add(1)
	go func(){
		for l := range _out {
			fmt.Println(l.name, l.result)
		}
		close(_out)
		wg.Done()
	}()

	wg.Wait()
}


