package main

import (
	"fmt")

func main() {
	SliceOfNumber := make([]int, 0)
	for x := 1; x < 10000000000; x++ {
		if x%2 == 0 && x%3 == 0 && x%4 == 0 && x%5 == 0 && x%7 == 0 && x%9 == 0 && x%8 == 0 && x%11 == 0 && x%13 == 0 && x%16 == 0 && x%17 == 0 && x%19 == 0 {
			v := append(SliceOfNumber, x)
			fmt.Println(v)
		}
	}
}
