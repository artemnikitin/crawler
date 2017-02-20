package crawler

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseHTML(t *testing.T) {
	cases := []struct {
		Desc  string
		Text  io.Reader
		Count int
	}{
		{"empty string", strings.NewReader(""), 0},
		{"not HTML", strings.NewReader("3f323grgrf2frfr4"), 0},
		{"invalid HTML", strings.NewReader(invalidHTML), 0},
		{"valid HTML", strings.NewReader(validHTML), validLinks},
		{"valid HTML", strings.NewReader(validHTMLWithURLWithoutHTTP), validLinks},
		{"valid HTML with javascript", strings.NewReader(withJavascriptURL), validLinks},
		{"valid HTML without URL", strings.NewReader(validHTMLNoURL), 0},
	}

	for _, v := range cases {
		t.Run(v.Desc, func(t *testing.T) {
			links, err := parseHTML(v.Text, "")
			if err != nil {
				t.Errorf("Get error: %s", err.Error())
				t.FailNow()
			}
			if len(links) != v.Count {
				t.Errorf("Expected: %d, actual: %d", v.Count, len(links))
			}
		})
	}

}

func TestWriteToFile(t *testing.T) {
	name := "test_correct.txt"
	list := []string{"first", "second", "third"}

	WriteToFile(list, name)

	golden, _ := ioutil.ReadFile("testdata")
	file, _ := ioutil.ReadFile(name)
	if !bytes.Equal(golden, file) {
		t.Error("Files should be equal!")
	}
}

func TestWriteToFileIncorrect(t *testing.T) {
	cases := []struct {
		Desc string
		List []string
		Name string
	}{
		{
			Desc: "Empty list",
			List: []string{},
			Name: "test_empty.txt",
		},
		{
			Desc: "Empty file name",
			List: []string{"first", "second", "third"},
			Name: "",
		},
		{
			Desc: "Different content",
			List: []string{"a", "b", "c"},
			Name: "test_different.txt",
		},
	}

	for _, v := range cases {
		t.Run(v.Desc, func(t *testing.T) {
			WriteToFile(v.List, v.Name)
			golden, _ := ioutil.ReadFile("testdata")
			file, _ := ioutil.ReadFile(v.Name)
			if bytes.Equal(golden, file) {
				t.Error("Files shouldn't be equal!")
			}
		})
	}
}
