package main

import (
	"os"
	"runtime/pprof"
	"time"
)

var s, s1, s2 []byte

func main() {
	leakyFunction()

	pprof.WriteHeapProfile(os.Stdout)
}

func leakyFunction() {
	for i:= 0; i < 10 * 1024 * 1024; i++{
		s = append(s, 's')
	}

	allocS1()
	allocS2()
	allocS3()
}

func allocS1() {
	for i:= 0; i < 10 * 1024 * 1024; i++{
		s1 = append(s1, 's')
	}
}


func allocS2() {
	for i:= 0; i < 10 * 1024 * 1024; i++{
		s2 = append(s2, 's')
	}
}

func allocS3() {
	for i:= 0; i < 2 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
	}
}

