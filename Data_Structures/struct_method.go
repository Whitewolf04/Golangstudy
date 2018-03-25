package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func (p Person) fullname() string {
	return p.First + p.Last
}

func main() {
	p1 := Person{"Tuan ", "To", 17}
	p2 := Person{"Le ", "Tran", 17}
	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
}
