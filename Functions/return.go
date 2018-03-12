package main

import "fmt"

func main() {
	fmt.Println(Hello("Jane", "Doe"))
}

func Hello(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
