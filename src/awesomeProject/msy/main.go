package main

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	"sync"
	"strings"
	"os"
	"bufio"
	"strconv"
)

type Item struct {
	name string
	price float32
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

func MSY() {

	clink := producer()
	var products [] <-chan string
	for i := 0; i != 20; i++ {
		products = append(products, consumer(clink))
	}

	for p := range merge(products...) {
		//fmt.Sprintf(p)
		fmt.Println(p)
	}
}

func merge(cs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan string) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}


func consumer(links <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)

		for l := range links {
			//time.Sleep(2 * time.Second)
			doc, err := goquery.NewDocument(l)
			if err != nil {
				log.Fatal(err)
			}
			doc.Find(".ajax_block_product").Each(func(i int, s *goquery.Selection) {
				product, exist := s.Find("a").Attr("title")
				if !exist {
					fmt.Println("couldn't find product")
					return
				}

				price := s.Find(".price").Text()
				out <- product+"||"+price
			})
		}
	}()
	return out
}

func producer() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for line := range open("data/cats.dat"){
			array := strings.Split(line, ",")
			s_number := strings.Trim(array[1], " \n")
			s_base := strings.Trim(array[0], " ")
			i_number, err := strconv.ParseInt(s_number, 10, 64)
			if err == nil {
				for i:=1; i!=int(i_number)+1; i++{
					url := fmt.Sprintf("%v#/page-%v", s_base, i)
					out<-url
				}
			}
		}

	}()
	return out
}

func main() {
	MSY()
}
