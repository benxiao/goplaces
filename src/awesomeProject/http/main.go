package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("https://www.nytimes.com")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	bytes, err :=ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))

}
