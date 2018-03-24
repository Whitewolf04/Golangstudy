package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4}
	yourSlice := []int{5, 6, 7, 8}
	mySlice = append(mySlice, yourSlice...)
	fmt.Println(mySlice)
	mySlice = append(mySlice[:2], yourSlice[3:]...) //delete from a slice
	fmt.Println(mySlice)
}
