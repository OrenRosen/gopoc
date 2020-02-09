package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

func findFanoutLocal(topic string, docs []string) int {
	var found int32

	g := len(docs)
	var wg sync.WaitGroup
	wg.Add(g)

	for _, doc := range docs {

		go func(doc string) {

			var lfound int32 // new

			defer func() {
				atomic.AddInt32(&found, lfound)
				wg.Done()
			}()

			// open the file
			f, err := os.OpenFile(doc, os.O_RDONLY, 0)
			if err != nil {
				log.Printf("Opening doc [%s]: error: %v", doc, err)
				return
			}

			// read file
			data, err := ioutil.ReadAll(f)
			if err != nil {
				_ = f.Close()
				log.Printf("Reading doc [%s]: error: %v", doc, err)
				return
			}
			_ = f.Close()

			// decode
			var d rssResponseXML
			if err := xml.Unmarshal(data, &d); err != nil {
				log.Printf("Decoding doc [%s]: error: %v", doc, err)
				return
			}

			// find
			for _, item := range d.Items {
				if strings.Contains(item.Title, topic) {
					lfound++ // new
					continue
				}

				if strings.Contains(item.Description, topic) {
					lfound++ // new
				}
			}
		}(doc)
	}

	wg.Wait()
	return int(found)
}
