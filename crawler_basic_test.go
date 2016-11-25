package crawler

import (
	"io"
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
		{"invalid HTML", strings.NewReader(InvalidHTML), 0},
		{"valid HTML", strings.NewReader(ValidHTML), ValidLinks},
		{"valid HTML without URL", strings.NewReader(ValidHTMLNoURL), 0},
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
