package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello World"
	}
}

func main() {
	greet := makeGreeter() //in order to write a func in a func, we must assign it to a variable(an anonymous func)
	fmt.Println(greet())
}
