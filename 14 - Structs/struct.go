package structExample

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	person := Person{name: "Alice", age: 30}
	fmt.Println(person)
}
