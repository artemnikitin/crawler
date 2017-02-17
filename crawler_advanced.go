package crawler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// GetListOfURLWithDepth returns list of all URL with specified depth of crawling
func GetListOfURLWithDepth(link string, depth int) ([]string, error) {
	var resmap map[string]string
	var result []string
	var temp []string
	var temp2 []string
	var wg sync.WaitGroup
	var m sync.Mutex

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

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
	result = lst
	temp = lst

	for i := 1; i <= depth; i++ {
		fmt.Println("Run:", i)
		wg.Add(len(temp))
		for _, v := range temp {
			go func(v string) {
				resp, err := client.Get(v)
				if err == nil {
					lst, err := parseHTML(resp.Body, link)
					if err == nil {
						m.Lock()
						temp2 = append(temp2, lst...)
						m.Unlock()
					}
					io.Copy(ioutil.Discard, resp.Body)
					resp.Body.Close()
				}
				wg.Done()
			}(v)
		}
		wg.Wait()

		temp = temp2
		temp2 = make([]string, 0)
		result = append(result, temp...)
	}

	resmap = make(map[string]string, 0)
	for _, v := range result {
		resmap[v] = v
	}
	result = make([]string, 0)
	for k := range resmap {
		result = append(result, k)
	}

	return result, nil
}
