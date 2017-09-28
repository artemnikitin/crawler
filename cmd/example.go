package main

import (
	"fmt"
	"log"
	"os"

	"github.com/artemnikitin/crawler"
)

func main() {
	urls, err := crawler.Do("https://www.facebook.com/", 2, 150)
	if err != nil {
		log.Fatal("Can't extract URLs from page ", err)
	}

	fmt.Println("Total # of URL:", len(urls))
	writeToFile(urls, "result_of_crawler_with_limited_concurrency.txt")
}

func writeToFile(list []string, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	for _, v := range list {
		file.WriteString(v)
		file.WriteString("\n")
	}

	file.Sync()
}
