package main

import (
	"github.com/Whitewolf04/Golangstudy/UsefulTools"
	"fmt"
)

func findPrimeNums()[]int{
	num := make([]int, 10003)
	for i := 2; i <= 10000000; i++{
		if UsefulTools.CheckPrimeNumber(i) == true{
			num = append(num, i)
		}
	}
	return num
}

func main(){
	fmt.Println(findPrimeNums()[10000:10001])
}
