package routing

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	DEFAULT_PORT              = "8080"
	CF_FORWARDED_URL_HEADER   = "X-Cf-Forwarded-Url"
	CF_PROXY_SIGNATURE_HEADER = "X-Cf-Proxy-Signature"
)

type MapForwardURL func(string) string

//NewProxy
func NewProxy(transport http.RoundTripper, forwardURL MapForwardURL) http.Handler {
	reverseProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			forwardedURL := forwardURL(req.Header.Get(CF_FORWARDED_URL_HEADER))
			url, err := url.Parse(forwardedURL)
			if err != nil {
				log.Fatalln(err.Error())
			}
			log.Println(url)
			req.URL = url
			req.Host = url.Host
		},
		Transport: transport,
	}
	return reverseProxy
}
