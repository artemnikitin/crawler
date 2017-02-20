package main

import (
	"fmt"
	"log"

	"github.com/artemnikitin/crawler"
)

func main() {
	urls, err := crawler.GetListOfURLWithLimitedConcurrency("https://www.facebook.com/", 2, 150)
	if err != nil {
		log.Fatal("Can't extract URLs from page ", err)
	}

	fmt.Println("Total # of URL:", len(urls))
	crawler.WriteToFile(urls, "result_of_crawler_with_limited_concurrency.txt")
}
