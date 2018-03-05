package Language_Fundamentals

import "fmt"

var X int = 0

func increment() int {
	X--
	return X
}

func main() {
	fmt.Println(increment())
}
