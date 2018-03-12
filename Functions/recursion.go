package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

func main() {
	fmt.Println(factorial(4))
}

//Although recursion seems to be the answer, there maybe other answer because recursion cost a lot of memory to run
