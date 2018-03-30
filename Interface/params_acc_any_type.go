package main

import "fmt"

type shit struct {
	Smell string
}

type dogshit struct {
	shit
	Solid bool
}

type bullshit struct {
	shit
	Big bool
}

func ShitCharacteristics(a interface{}) {
	fmt.Println(a)
}

func main() {
	Eric := dogshit{shit{"crap"}, true}
	Dylan := bullshit{shit{"nice"}, false}
	ShitCharacteristics(Eric)
	ShitCharacteristics(Dylan)
}
