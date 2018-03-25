package main

import "fmt"

type jungle struct {
	monster string
	amount  int
}

func main() {
	p1 := jungle{"alistar", 4}
	fmt.Println(p1)
	fmt.Printf("%T \n", p1)
	fmt.Println(p1.amount)
	fmt.Println(p1.monster)
}
