package main

import (
	"crypto/tls"
	"net/http"

	"github.com/datianshi/route-service-proxy/routing"
)

func main() {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	url := func(a string) string {
		return a
	}
	server := &http.Server{
		Addr:    ":8080",
		Handler: routing.NewProxy(transport, url),
	}
	server.ListenAndServe()
}
