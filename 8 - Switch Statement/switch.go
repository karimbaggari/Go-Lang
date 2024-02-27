package switche

import "fmt"


func main() {
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Some other day")
	}
	
}