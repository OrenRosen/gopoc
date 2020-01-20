package main

import (
	"fmt"
	"log"
	"net/http"
)



func main() {
	http.Handle("/find", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		docs := make([]string, 1000)
		for i := range docs {
			docs[i] = "newsfeed.xml"
		}

		topic := "some_topic"

		n := findNumCPU(topic, docs)
		log.Printf("Found %d documents with topic %s", n, topic)
		_, _ = fmt.Fprintf(w, "Found %d documents with topic %s", n, topic)
	}))

	log.Println("listening on :6060")
	_ = http.ListenAndServe(":6060", nil)
}

