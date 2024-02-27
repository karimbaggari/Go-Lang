package slices

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}

	for index, value := range slice {
        fmt.Println(index, value)
    }
}
