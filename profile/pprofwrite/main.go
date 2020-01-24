package main

var s, s1, s2 []byte

func main() {
	leakyFunction()
}

func leakyFunction() {
	for i:= 0; i < 10 * 1024 * 1024; i++{
		s = append(s, 's')
	}

	allocS1()
	allocS2()
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
