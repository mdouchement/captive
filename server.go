package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/ncsi.txt", func(w http.ResponseWriter, r *http.Request) {
		trace(w, r)
		fmt.Fprintf(w, "Microsoft NCSI")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		trace(w, r)
		fmt.Fprintf(w, "<HTML><HEAD><TITLE>Success</TITLE></HEAD><BODY>Success</BODY></HTML>")
	})

	listen := os.Getenv("CAPTIVE_LISTEN")
	if listen == "" {
		listen = ":8080"
	}
	log.Println("Listening on", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}

func trace(w http.ResponseWriter, r *http.Request) {
	requestID := strconv.FormatInt(time.Now().UnixNano(), 36)

	log.Printf("[%s] %s %s", requestID, r.Method, r.RequestURI)
	log.Printf("[%s] From %s", requestID, r.RemoteAddr)
	log.Printf("[%s] Content-Length %d", requestID, r.ContentLength)

	headers, err := json.Marshal(r.Header)
	if err != nil {
		log.Printf("[%s] %v", requestID, r.Header)
	} else {
		log.Printf("[%s] %s", requestID, headers)
	}
}
