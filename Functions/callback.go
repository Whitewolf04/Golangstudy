package main

import "fmt"

func wrapper(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	wrapper([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

//callback pass func as an argument
