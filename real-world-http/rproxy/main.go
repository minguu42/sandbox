package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	director := func(r *http.Request) {
		r.URL.Scheme = "http"
		r.URL.Host = ":8080"
	}
	modifier := func(r *http.Response) error {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("reading body error: %w", err)
		}
		newBody := bytes.NewBuffer(body)
		newBody.WriteString("via Proxy\n")
		r.Body = io.NopCloser(newBody)
		r.Header.Set("Content-Length", strconv.Itoa(newBody.Len()))
		return nil
	}
	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}
	server := http.Server{
		Addr:    "127.0.0.1:9000",
		Handler: rp,
	}
	log.Println("Listening on http://127.0.0.1:9000")
	log.Fatal(server.ListenAndServe())
}
