package main

import "fmt"

func numGenerator()[]int{
	numList := make([]int,101)
	for i:= 0; i <= 100; i++{
		numList = append(numList, i)
	}
	return numList
}

func square(num int) int{
	var result int
	result = num*num
	return result
}

func squareSum(numlist []int)[]int{
	resultList := make([]int, 101)
	for _,k := range numlist{
		resultList = append(resultList, square(k))
	}
	return resultList
}

func findSum(numList []int) int{
	var sumOfList int
	for _,l := range numList{
		sumOfList += l
	}
	return sumOfList
}

func main(){
	var sumOfSquare int
	var sumOfList int
	sumOfList = findSum(numGenerator())
	sumOfSquare = findSum(squareSum(numGenerator()))
	fmt.Println(square(sumOfList)-sumOfSquare)
}
