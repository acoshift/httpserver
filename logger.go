package main

import (
	"log"
	"net"
	"net/http"
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nw := &logResponseWriter{
			ResponseWriter: w,
		}

		remoteHost, _, _ := net.SplitHostPort(r.RemoteAddr)
		if remoteHost == "" {
			remoteHost = r.RemoteAddr
		}

		defer func() {
			log.Printf("%s %d %s %s", remoteHost, nw.status, r.Method, r.RequestURI)
		}()
		h.ServeHTTP(nw, r)
	})
}

type logResponseWriter struct {
	http.ResponseWriter

	wroteHeader bool
	status      int
}

func (w *logResponseWriter) WriteHeader(status int) {
	if w.wroteHeader {
		return
	}
	w.wroteHeader = true
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *logResponseWriter) Write(p []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(p)
}
