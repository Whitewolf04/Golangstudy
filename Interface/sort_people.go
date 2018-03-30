package main

import (
	"fmt"
	"sort"
)

type animal []string

func (a animal) Len() int           { return len(a) }
func (a animal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a animal) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	type people []string
	studyGroup := people{"Zoey", "Brad", "Fred", "Hudson", "Tim", "Anson"}
	sort.Slice(studyGroup, func(i, j int) bool {
		return studyGroup[i] > studyGroup[j]
	})
	fmt.Println(studyGroup)

	s := animal{"Monkey", "Gorilla", "Giraffe", "Dog", "Lion", "Tiger", "White Wolf"}
	sort.Sort(s)
	fmt.Println(s)
	fmt.Println(s.Len())
}
