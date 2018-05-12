package main

import "fmt"

func main() {
	a := 67
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a
	fmt.Println(b)
	fmt.Println(*b) // This is dereferencing the value address

	*b = 42        // This changes the value in this memory address into 42
	fmt.Println(a) // As a result, a changes into 42
}
