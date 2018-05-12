package main

import "fmt"

const poundsToKilograms float64 = 0.454

func main() {
	var Pounds float64
	fmt.Print("How much do you weight(lbs): ")
	fmt.Scan(&Pounds)
	Kilograms := Pounds * poundsToKilograms
	fmt.Println(Pounds, "lbs is", Kilograms, "kg")
}
