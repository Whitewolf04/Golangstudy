package main

import "fmt"

func main() {
	fmt.Println(aloha("Jane", "Doe"))
}

func aloha(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
