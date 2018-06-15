package main

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/datianshi/route-service-proxy/routing"
)

func main() {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	url := func(a string) string {
		return a
	}

	//This is to accormodate the issue DNS can not updated duto idle connections
	//https://github.com/golang/go/issues/23427
	go func() {
		for {
			transport.CloseIdleConnections()
			time.Sleep(5 * time.Second)
		}
	}()
	server := &http.Server{
		Addr:    ":8080",
		Handler: routing.NewProxy(transport, url),
	}
	server.ListenAndServe()
}
