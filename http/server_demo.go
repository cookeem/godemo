package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// hello world, the web server
func HelloServer(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, time.Now(), req.Header, req.Host)
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
