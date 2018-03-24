package main

import "fmt"

func main() {
	myGreetings := map[int]string{
		1: "Hello",
		2: "Bello",
		3: "Khello",
		4: "Jello",
		5: "Hi",
	}

	for key, val := range myGreetings {
		fmt.Println(key, "-", val)
	}
}
