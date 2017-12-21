package main

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	"time"
	"sync"
	"strings"
)

type Item struct {
	name string
	price float32
}


func PCCaseGear() {
	doc, err := goquery.NewDocument("http://www.pccasegear.com")
	if err != nil {
		log.Fatal(err)
	}
	var links []string
	// Find the review items
	doc.Find(".left_cat_menu a").Each(func(i int, s *goquery.Selection) {
		// For each item fo
		link, found := s.Attr("href")
		if found {
			links = append(links, link)
		}
	})
	var products [] <-chan string
	clink := producer(links)
	for i := 0; i != 1; i++ {
		products = append(products, consumer(clink))
	}

	for p := range merge(products...) {
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
			time.Sleep(2 * time.Second)
			doc, err := goquery.NewDocument(l)
			if err != nil {
				log.Fatal(err)
			}
			doc.Find(".media-list li").Each(func(i int, s *goquery.Selection) {
				product := strings.Trim(s.Find(".media-heading").Text(), " \n")
				price := s.Find(".btn-for-mobile .product-price").Text()
				out <- product+"||"+price
			})
		}
	}()
	return out
}

func producer(links []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, l := range links {
			out <- l
		}
	}()
	return out
}

func main() {
	PCCaseGear()
}
