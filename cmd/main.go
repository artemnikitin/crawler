package main

import (
	"fmt"
	"log"

	"github.com/artemnikitin/crawler"
)

func main() {
	urls, err := crawler.GetListOfURL("http://google.com")
	if err != nil {
		log.Fatal("Can't extract URL from page ", err)
	}

	fmt.Println("Page http://google.com contains next URL:")
	for _, v := range urls {
		fmt.Println(v)
	}
}
