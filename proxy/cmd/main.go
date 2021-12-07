package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"

	"proxy-server/lib"
)

func main() {
	var (
		pemPath string
		proto   = "http"
		keyPath string
	)

	flag.StringVar(&pemPath, "pem", "server.pem", "path to pem file")
	flag.StringVar(&keyPath, "key", "server.key", "path to key file")
	flag.StringVar(&proto, "proto", "http", "Proxy protocol (http or https)")
	flag.Parse()

	if proto != "http" && proto != "https" {
		log.Fatal("Protocol must be either http or https")
	}

	server := &http.Server{
		Addr: ":8888",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				lib.HandleTunneling(w, r)
			} else {
				lib.HandleHTTP(w, r)
			}
		}),
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	if proto == "https" {
		log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}
