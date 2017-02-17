package main

import (
	"fmt"
	"log"

	"github.com/artemnikitin/crawler"
)

func main() {
	urls, err := crawler.GetListOfURLWithDepth("http://google.com", 3)
	if err != nil {
		log.Fatal("Can't extract URLs from page ", err)
	}

	fmt.Println("Total # of URL:", len(urls))
	if len(urls) <= 100 {
		fmt.Println("URL:")
		for _, v := range urls {
			fmt.Println(v)
		}
	}
}
