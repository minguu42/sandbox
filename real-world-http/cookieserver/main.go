package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body>2回目以降</body></html>")
	} else {
		fmt.Fprintf(w, "<html><body>初回</body></html>")
	}
}

func main() {
	var s http.Server
	http.HandleFunc("/", handler)
	log.Println("start http server on :18888")
	s.Addr = ":18888"
	log.Println(s.ListenAndServe())
}
