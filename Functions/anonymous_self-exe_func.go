package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello World")
	}()
}
//An anonymous self-exe func exe itself with the parens after it
