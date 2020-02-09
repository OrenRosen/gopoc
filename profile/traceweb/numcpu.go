package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

func findNumCPU(topic string, docs []string) int {
	var found int32

	g := runtime.NumCPU() // new
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, len(docs))
	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	for i := 0; i < g; i ++ {

		go func() {

			var lfound int32 // new

			defer func() {
				atomic.AddInt32(&found, lfound)
				wg.Done()
			}()

			for doc := range ch {

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
			}
		}()
	}

	wg.Wait()
	return int(found)
}
