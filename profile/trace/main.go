package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)



func main() {
	docs := make([]string, 1000)
	for i := range docs {
		docs[i] = "newsfeed.xml"
	}

	topic := "some_topic"

	n := find(topic, docs)
	log.Printf("Found %d documents with topic %s", n, topic)
}

func find(topic string, docs []string) int {
	var found int

	for _, doc := range docs {
		f, err := os.OpenFile(doc, os.O_RDONLY, 0)
		if err != nil {
			log.Printf("Opening doc [%s]: error: %v", doc, err)
			return 0
		}

		data, err := ioutil.ReadAll(f)
		if err != nil {
			f.Close()
			log.Printf("Reading doc [%s]: error: %v", doc, err)
			return 0
		}
		f.Close()

		var d rssResponseXML
		if err := xml.Unmarshal(data, &d); err != nil {
			log.Printf("Decoding doc [%s]: error: %v", doc, err)
			return 0
		}

		for _, item := range d.Items {
			if  strings.Contains(item.Title, topic) {
				found++
				continue
			}

			if  strings.Contains(item.Description, topic) {
				found++
			}
		}
	}

	return found
}