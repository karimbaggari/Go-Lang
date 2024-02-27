package pointers

import (
	"fmt"
)

func main() {
	var p *int
	x := 5
	p = &x

	fmt.Println(*p) // Prints the value stored at the memory location pointed by p

}