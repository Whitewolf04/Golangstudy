package main

import "fmt"

func main() {
	greet("John")
	greet("Casey")
}

func greet(name string) {
	fmt.Println("Hello", name)
}
