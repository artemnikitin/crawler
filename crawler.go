package crawler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Do returns list of all URL with specified depth of crawling
// and allow to limit simultaneous number of goroutines
func Do(link string, depth, ccy int) ([]string, error) {
	var result []string
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	list, err := getAndParse(client, link)
	if err != nil {
		return nil, err
	}
	result = append(result, list...)

	for i := 1; i <= depth; i++ {
		fmt.Println("Run:", i)
		list = processListURL(client, ccy, list)
		result = append(result, list...)
	}

	return deleteDoubles(result), nil
}

func getAndParse(client *http.Client, link string) ([]string, error) {
	resp, err := client.Get(link)
	if err != nil {
		return nil, err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	lst, err := parseHTML(resp.Body, link)
	if err != nil {
		return nil, err
	}

	return lst, nil
}

func processListURL(client *http.Client, ccy int, list []string) []string {
	var result []string
	var m sync.Mutex
	semaphore := make(chan bool, ccy)

	for _, url := range list {
		semaphore <- true
		go func(url string) {
			list, err := getAndParse(client, url)
			if err == nil {
				m.Lock()
				result = append(result, list...)
				m.Unlock()
			}
			<-semaphore
		}(url)
	}

	for i := 0; i < cap(semaphore); i++ {
		semaphore <- true
	}

	return result
}

func deleteDoubles(list []string) []string {
	tempmap := make(map[string]string, 0)
	for _, v := range list {
		tempmap[v] = v
	}

	temp := make([]string, 0)
	for k := range tempmap {
		temp = append(temp, k)
	}

	return temp
}
