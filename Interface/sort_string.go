package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Bla", "Bithc", "Flask", "Hsalf", "Ksalf"}
	//t := sort.StringSlice(s)
	//fmt.Printf("%T \n", t)
	fmt.Printf("s converted to StringSlice: %T\n", sort.StringSlice(s))
	fmt.Println(sort.StringSlice(s).Search("Flask"))
	//sort.Sort(sort.StringSlice(s))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
