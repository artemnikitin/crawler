package crawler

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// GetListOfURL returns list or URL on a page
func GetListOfURL(link string) ([]string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(link)
	if err != nil {
		return nil, err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	return parseHTML(resp.Body, link)
}
