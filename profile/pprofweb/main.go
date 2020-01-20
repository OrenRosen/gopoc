package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"
)

var s, s1, s2, s3 []byte

func main() {
	http.Handle("/work", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("working")
		leakyFunction()
		fmt.Fprintf(w, "DONE")
	}))

	http.ListenAndServe(":6060", nil)
}

func leakyFunction() {
	log.Println("leakyFunction")
	allocS1()
	allocS2()
	allocS3()

	for i:= 0; i < 2 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s = append(s, 's')
	}
}

func allocS1() {
	log.Println("allocS1")
	for i:= 0; i < 2 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s1 = append(s1, 's')
	}
}

func allocS2() {
	log.Println("allocS2")
	for i:= 0; i < 2 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s2 = append(s2, 's')
	}
}

func allocS3() {
	log.Println("allocS3")
	for i:= 0; i < 2 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s3 = append(s3, 's')
	}
}

