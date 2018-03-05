package Language_Fundamentals

import "fmt"

func main() {
	x := 547 / 17
	fmt.Println(x)

	var y *int = &x
	*y = 547 % 17
	fmt.Println(x)
}
