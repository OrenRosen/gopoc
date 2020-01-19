package main

import (
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//"log"
	//"net/http"
	//"fmt"

	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var s, s1, s2, s3 []byte

func main() {
	runtime.SetCPUProfileRate(1)
	leakyFunction()
	if err := pprof.WriteHeapProfile(os.Stdout); err != nil {
		panic("AHHHHH")
	}
}

func leakyFunction() {
	allocS1()
	allocS2()
	allocS3()

	for i:= 0; i < 2 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s = append(s, 's')
	}
}

func allocS1() {
	//s := make([]byte, 1)
	for i:= 0; i < 2 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s1 = append(s1, 's')
	}
	//allocHuge()
}

func allocS2() {
	for i:= 0; i < 2 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s2 = append(s2, 's')
	}
}

func allocS3() {
	for i:= 0; i < 2 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s3 = append(s3, 's')
	}
}

