package controlFlow

import "fmt"

func main() {
	x := 5
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is not greater than 10")
	}
}