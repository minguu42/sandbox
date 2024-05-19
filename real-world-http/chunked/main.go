package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handlerChunkedResponse(w http.ResponseWriter, r *http.Request) {
	c := http.NewResponseController(w)
	for i := range 10 {
		fmt.Fprintf(w, "Chunk #%d\n", i)
		c.Flush()
		time.Sleep(500 * time.Millisecond)
	}
	c.Flush()
}

func main() {
	var s http.Server
	http.HandleFunc("/chunked", handlerChunkedResponse)
	log.Println("start http listening :18888")
	s.Addr = ":18888"
	log.Println(s.ListenAndServe())
}
