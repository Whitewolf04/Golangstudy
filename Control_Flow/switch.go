package main

import "fmt"

func thuhai() {
	switch "Eli" {
	case "John":
		fmt.Println("What's up")
	case "Dylan":
		fmt.Println("Ditme")
		fallthrough
	case "Eric":
		fmt.Println("Ngul")
	case "Eli", "Hobeum":
		fmt.Println("Wassup")
	default:
		fmt.Println("Where are your friends")
	}
}

func main() {
	var o string
	fmt.Println("Ten ban la gi")
	fmt.Scan(&o)
	switch {
	case (o) == "Tim":
		fmt.Println("Hello")
	case len(o) == 2:
		fmt.Println("Ditbo")
	default:
		fmt.Println("THoi")
	}
}
