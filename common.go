package crawler

import (
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

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
					if strings.Contains(v.Val, "javascript:") {
						break
					}
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

// WriteToFile allows to write list of URL to file
func WriteToFile(list []string, name string) {
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
