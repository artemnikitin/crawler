package crawler

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
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

func parseHTML(body io.Reader, baseURL string) ([]string, error) {
	var result []string

	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	var f func(*html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, v := range node.Attr {
				if v.Key == "href" {
					if strings.HasPrefix(v.Val, "http") {
						result = append(result, v.Val)
					} else {
						result = append(result, baseURL+v.Val)
					}
					break
				}
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return result, nil
}
