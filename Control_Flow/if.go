package main

import "fmt"

func main() {
	var x float64
	fmt.Println("How old are you?")
	fmt.Scan(&x)
	if x <= 20 && x > 0 {
		fmt.Println("You are young")
	} else if x > 20 && x <= 50 {
		fmt.Println("You are good")
	} else if x > 50 && x <= 100 {
		fmt.Println("You are fk old!")
	} else {
		fmt.Println("You are lying or you are stupid af")
	}
}
