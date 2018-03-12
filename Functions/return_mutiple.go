package main

import "fmt"

func main() {
	fmt.Println(hi("John", "Casey"))
}

func hi(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
