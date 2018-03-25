package main

import "fmt"

type view struct {
	front  int
	behind int
	wing   int
}

type judge struct {
	view
	judgement bool
}

func (v view) Greeting() {
	fmt.Println("Hello")
}

func (j judge) Greeting() {
	fmt.Println("Lo con cac")
}

func main() {
	v1 := judge{
		view: view{
			front:  45,
			behind: 65,
			wing:   54,
		},
		judgement: true,
	}
	v2 := judge{
		view: view{
			front:  34,
			behind: 23,
			wing:   90,
		},
		judgement: false,
	}
	v1.Greeting()
	v2.view.Greeting()
}
