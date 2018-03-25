package main

import "fmt"

func main() {
	primenumbers := make([]int, 100)
	var sum int
	for i := 0; i < 2000000; i++ {
		if i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0 {
			primenumbers = append(primenumbers, i)
		}
	}
	fmt.Println(primenumbers)
	for _, v := range primenumbers {
		sum = sum + v + 3 + 5 + 2 + 7
	}
	fmt.Println(sum)
}
