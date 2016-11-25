package crawler

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

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

func parseHTML(body io.ReadCloser, baseURL string) ([]string, error) {
	var result []string

	t := html.NewTokenizer(body)
	for {
		tkn := t.Next()
		switch {
		case tkn == html.ErrorToken:
			//return result, errors.New("Error while parsing HTML")
			return result, nil
		case tkn == html.StartTagToken:
			tmp := t.Token()
			if tmp.Data == "a" {
				for _, v := range tmp.Attr {
					if v.Key == "href" {
						if strings.HasPrefix(v.Val, "http://") || strings.HasPrefix(v.Val, "https://") {
							result = append(result, v.Val)
						} else {
							result = append(result, baseURL+v.Val)
						}
						break
					}
				}
			}
		}
	}
}
