package main

import "fmt"

func main() {
	for i := 5000; i <= 5100; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
}

func second() {
	for j := 50; j <= 140; j++ {
		fmt.Printf("%v - %v - %v\n", j, string(j), []byte(string(j)))
	}
}
