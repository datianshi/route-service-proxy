package main

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

var routeMaping map[string]string

const (
	CF_FORWARDED_URL_HEADER     string = "X-Cf-Forwarded-Url"
	CF_PROXY_SIGNATURE_HEADER   string = "X-Cf-Proxy-Signature"
	CF_PROXY_SIGNATURE_METADATA string = "X-CF-Proxy-Metadata"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", ProxyTo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//ProxyTo Proxy the request to another application
func ProxyTo(res http.ResponseWriter, req *http.Request) {
	log.Println("Dump the http request:")
	request, _ := httputil.DumpRequest(req, false)
	log.Println(string(request))
	forwardedURL := req.Header.Get(CF_FORWARDED_URL_HEADER)
	sigHeader := req.Header.Get(CF_PROXY_SIGNATURE_HEADER)
	sigMetadata := req.Header.Get(CF_PROXY_SIGNATURE_METADATA)

	res.Header().Set(CF_FORWARDED_URL_HEADER, forwardedURL)
	res.Header().Set(CF_PROXY_SIGNATURE_HEADER, sigHeader)
	res.Header().Set(CF_PROXY_SIGNATURE_METADATA, sigMetadata)

	res.WriteHeader(200)
}
