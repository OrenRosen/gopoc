package main

import (
	"os"
	//"os"
	//"runtime/pprof"
	"runtime/pprof"
	"time"
)

var s, s1, s2, s3 []byte

func main() {
	leakyFunction()
}

func leakyFunction() {
	allocS1()
	allocS2()

	for i:= 0; i < 10 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s = append(s, 's')
	}

	_ = pprof.WriteHeapProfile(os.Stdout)
}

func allocS1() {
	for i:= 0; i < 5 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s1 = append(s1, 's')
	}
}


func allocS2() {
	for i:= 0; i < 5 * 1024 * 1024; i++{
		if i == 3000 { time.Sleep(500 * time.Millisecond) }
		s2 = append(s2, 's')
	}
}


