package main

import "fmt"

func checkPrimeNumber(num int) bool {
	for i := 2; i < 2000000; i++ {
		if num%i != 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkPrimeNumber(12))
}
