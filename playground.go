package main

import "fmt"

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater than 10")
	}
	return fmt.Sprint("x is less than 10")
}

func main() {
	x := 8
	fmt.Println(evalInt(x))
}
