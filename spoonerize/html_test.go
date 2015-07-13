package spoonerize

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func addTags(s string) string {
	return "<html><head></head><body>" + s + "</body></html>"
}

func assertHTMLSpoonerized(t *testing.T, s1 string, s2 string) {
	s1 = addTags(s1)
	s2 = addTags(s2)
	b, _ := ioutil.ReadAll(SpoonerizeHTML(bytes.NewBufferString(s1), ""))
	assertEqual(t, string(b), s2)
}

func assertHTMLSpoonerizedExtra(t *testing.T, s1 string, s2 string, extra string) {
	s1 = addTags(s1)
	s2 = addTags(s2)
	b, _ := ioutil.ReadAll(SpoonerizeHTML(bytes.NewBufferString(s1), extra))
	assertEqual(t, string(b), s2)
}

func TestSpoonerizeHTML(t *testing.T) {
	assertHTMLSpoonerized(t, "hello world", "wello horld")
	assertHTMLSpoonerized(t, "<div>hello world</div>", "<div>wello horld</div>")
	assertHTMLSpoonerized(t, "<div>Hello world</div>", "<div>Wello horld</div>")
	assertHTMLSpoonerized(t, "<div class=\"hello world\"></div>", "<div class=\"hello world\"></div>")
	assertHTMLSpoonerized(t, "<div><span>hello world</div>", "<div><span>wello horld</span></div>")
	assertHTMLSpoonerized(t,
		"<div><span>hello world</span></div>",
		"<div><span>wello horld</span></div>")
	assertHTMLSpoonerized(t,
		"<div><span>hello world</span></div><div><span>scrape block</span></div>",
		"<div><span>wello horld</span></div><div><span>blape scrock</span></div>")
	assertHTMLSpoonerized(t, "<script>hello world</script>", "<script>hello world</script>")
	assertHTMLSpoonerized(t, "<style>hello world</style>", "<style>hello world</style>")

	// Make sure Close doesn't error.
	SpoonerizeHTML(bytes.NewBufferString(""), "").Close()
}

func TestSpoonerizeHTMLWithExtra(t *testing.T) {
	assertHTMLSpoonerizedExtra(t, "hello world", "wello horldfoo", "foo")
	assertHTMLSpoonerizedExtra(t,
		"<div>hello world</div>",
		"<div>wello horld</div><span>foo bar</span>",
		"<span>foo bar</span>")
	assertHTMLSpoonerizedExtra(t,
		"<div>hello world</div>",
		"<div>wello horld</div><span>foo bar</span><div>baz</div>",
		"<span>foo bar</span><div>baz</div>")
}
