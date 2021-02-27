package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	const url = "http://www.google.com/robots.txt"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
