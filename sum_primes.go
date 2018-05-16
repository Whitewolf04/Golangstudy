package main

import "fmt"

func checkPrimeNumber(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}

func numgen() chan int {
	number := make(chan int)
	go func() {
		for j := 2; j < 2000000; j++ {
			number <- j
		}
		close(number)
	}()
	return number
}

func calculator(n chan int) chan int {
	total := make(chan int)
	go func() {
		for v := range n {
			if checkPrimeNumber(v) == true {
				total <- v
			}
		}
		close(total)
	}()
	return total
}

func main() {
	var total int
	for h := range calculator(numgen()) {
		total += h
		fmt.Println(total)
	}
}
