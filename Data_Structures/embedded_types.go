package main

import "fmt"

type doubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   45,
		},
		LicenseToKill: true,
	}
	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "Moneypie",
			Age:   19,
		},
		LicenseToKill: false,
	}
	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)
}
