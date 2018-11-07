package main

import "fmt"

func numberGen()chan int{ //generate numbers
	numberList := make(chan int)
	go func() {
		for i := 0; i < 1000000; i++{
			numberList <- i
		}
		close(numberList)
	}()
	return numberList
}

func checkPrimeNumber(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}

func findPrime(list chan int)chan int{
	primeNumbers := make(chan int)
	go func() {
		for v := range list{
			if checkPrimeNumber(v) == true{
				primeNumbers <- v
			}
		}
		close(primeNumbers)
	}()
	return primeNumbers
}

func CalculateBiggestPrime(list chan int)chan int{
	primeFactors := make(chan int)
	go func() {
		for k := range list{
			if  600851475143%k == 0{
				primeFactors <- k
			}
		}
		close(primeFactors)
	}()
	return primeFactors
}

func main(){
	var primeFactors int
	for h := range CalculateBiggestPrime(findPrime(numberGen())){
		primeFactors = h
		fmt.Println(primeFactors)
	}
}


