package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type TestServer struct {
	Version int64
}

func (ts *TestServer) StartServer() {
	http.HandleFunc("/", ts.handler)
	http.HandleFunc("/healthcheck", handler1)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func (ts *TestServer) handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-Type", "application/json")
	fmt.Fprintf(w, `{"version":%d,"time":"%v"}`, ts.Version, time.Now())
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy")
}
