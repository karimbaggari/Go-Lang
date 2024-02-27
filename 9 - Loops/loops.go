package loops

import "fmt"

func main() {

	var slice []int

        slice = make([]int, 5)

		for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	
	for index, value := range slice {
		fmt.Println(index, value)
	}
	
	for {
		// Infinite loop
	}
	
}