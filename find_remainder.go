package main

import "fmt"

func main() {
	var so1 int
	var so2 int
	fmt.Print("Nhap so lon\t")
	fmt.Scan(&so1)
	fmt.Print("Nhap so nho\t")
	fmt.Scan(&so2)
	switch {
	case so1 > so2:
		fmt.Println("Ket qua", so1%so2)
	case so1 < so2:
		fmt.Println("Error")
	}
}
