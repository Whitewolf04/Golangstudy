package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type boat struct {
	vehicle
	Length int
}

type plane struct {
	vehicle
	Jet bool
}

func main() {
	prius := car{}
	tacoma := car{}
	ferrari := car{}
	boeing787 := plane{}
	boeing747 := plane{}
	airbusA350 := plane{}
	nautique := boat{}
	malibu := boat{}
	rides := []vehicles{prius, tacoma, ferrari, boeing747, boeing787, airbusA350, nautique, malibu}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}
