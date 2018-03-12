package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}
func world() {
	fmt.Println("World")
}
func main() {
	defer world()
	hello()
}
