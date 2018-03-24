package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)
	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Cap: ", cap(mySlice), "Value: ", mySlice[i])
	}
}

//append is used when the value exceeds the length of the slice
