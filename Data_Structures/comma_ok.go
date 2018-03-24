package main

import "fmt"

func main() {
	myGreetings := map[int]string{
		0: "Bonjour",
		1: "Chao",
		2: "Dunkirk",
		3: "Nihao",
		4: "Camsamita",
	}
	//delete(myGreetings, 2)
	if val, exist := myGreetings[2]; exist {
		delete(myGreetings, 2)
		fmt.Println("val: ", val)
		fmt.Println("exist: ", exist)
	} else {
		fmt.Println("That value does not exist")
		fmt.Println("val: ", val)
		fmt.Println("exist: ", exist)
	}
}
