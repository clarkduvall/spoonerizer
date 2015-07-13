package main

import (
	"bytes"
	"github.com/clarkduvall/spoonerizer/spoonerize"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type SneakyTransport struct {
	http.RoundTripper
	ExtraHTML string
}

func (t *SneakyTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	// Use default value for Host.
	req.Host = ""

	// Make sure the response isn't gzipped.
	req.Header.Del("Accept-Encoding")

	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return
	}

	for _, v := range resp.Header["Content-Type"] {
		if strings.Contains(v, "text/html") {
			oldBody := resp.Body
			resp.Body = spoonerize.SpoonerizeHTML(oldBody, t.ExtraHTML)
			oldBody.Close()
			break
		}
	}
	return
}

func extraHTML() string {
	html, err := ioutil.ReadFile("extra.html")
	if err != nil {
		return ""
	}

	var b bytes.Buffer
	t := template.Must(template.New("extra").Parse(string(html)))
	t.Execute(&b, os.Getenv("GA_TRACKING_ID"))
	return b.String()
}

func main() {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		log.Fatal("Please set the BASE_URL environment variable.")
	}

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	rev := httputil.NewSingleHostReverseProxy(parsedURL)
	rev.Transport = &SneakyTransport{http.DefaultTransport, extraHTML()}

	http.ListenAndServe(":"+os.Getenv("PORT"), rev)
}
