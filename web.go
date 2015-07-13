package main

import (
	"github.com/clarkduvall/nackerhews/spoonerize"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	rev := httputil.NewSingleHostReverseProxy(spoonerize.ParsedBaseURL)
	rev.Transport = &SneakyTransport{http.DefaultTransport}
	http.HandleFunc("/", rev.ServeHTTP)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

type SneakyTransport struct {
	http.RoundTripper
}

func (t *SneakyTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	req.RequestURI = ""
	req.Host = ""

	// Make sure the response isn't gzipped.
	req.Header.Del("Accept-Encoding")

	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	for _, v := range resp.Header["Content-Type"] {
		if strings.Contains(v, "text/html") {
			oldBody := resp.Body
			resp.Body = spoonerize.SpoonerizeHTML(oldBody)
			oldBody.Close()
			break
		}
	}
	return resp, nil
}
