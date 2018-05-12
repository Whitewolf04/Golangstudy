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

func shitCharacteristics(a interface{}) {
	fmt.Println(a)
}

func main() {
	Eric := dogshit{shit{"crap"}, true}
	Dylan := bullshit{shit{"nice"}, false}
	shitCharacteristics(Eric)
	shitCharacteristics(Dylan)
}
