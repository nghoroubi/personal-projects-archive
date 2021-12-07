package lib

import (
	"io"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// HandleTunneling ...
func HandleTunneling(w http.ResponseWriter, r *http.Request) {
	destConn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go Transfer(destConn, clientConn)
	go Transfer(clientConn, destConn)
}

// Transfer ...
func Transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	_, _ = io.Copy(destination, source)
}

// HandleHTTP ...
func HandleHTTP(w http.ResponseWriter, req *http.Request) {

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer func() {
		logrus.StandardLogger().
			WithField("request_body",req.Body).
			WithField("request_url",req.URL.String()).
			WithField("request_headers",req.Header).
			WithField("method",req.Method).
			WithField("response",resp.Body).
			Infoln()
		resp.Body.Close()
	}()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
