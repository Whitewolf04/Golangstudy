package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func (p person) fullname() string {
	return p.First + p.Last
}

func main() {
	p1 := person{"Tuan ", "To", 17}
	p2 := person{"Le ", "Tran", 17}
	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
}
