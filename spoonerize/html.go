package spoonerize

import (
	"bytes"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"net/url"
)

const BaseURL = "https://news.ycombinator.com"

var ParsedBaseURL, _ = url.Parse(BaseURL)

type bufferCloser struct {
	bytes.Buffer
}

func (b *bufferCloser) Close() error {
	return nil
}

func SpoonerizeHTML(r io.Reader) io.ReadCloser {
	doc, _ := html.Parse(r)
	var f func(*html.Node)
	f = func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			n.Data = string(Spoonerize([]byte(n.Data)))
		case html.ElementNode:
			switch n.DataAtom {
			case atom.Style, atom.Script:
				return
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	d := &bufferCloser{}
	html.Render(d, doc)
	return d
}
