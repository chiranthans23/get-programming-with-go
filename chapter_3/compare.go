package chapter_3

import "fmt"

//CompareValues comapres two operands with the operator and prints the corresponding results
func ComapreValues(a int, b int, operator string) {

	switch operator {
	case "<":
		if a < b {
			fmt.Printf("%v is less than %v\n", a, b)
		} else {
			fmt.Printf("%v is not lesser than %v\n", a, b)
		}
	case ">":
		if a > b {

			fmt.Printf("%v is greater than %v\n", a, b)
		} else {
			fmt.Printf("%v is not greater than %v\n", a, b)
		}
	case "=", "==":
		if a == b {
			fmt.Printf("%v is equal to %v\n", a, b)
		} else {
			fmt.Printf("%v is not equal to %v\n", a, b)
		}
	default:
		fmt.Printf("Our basic comparator is not devised to compare %v and %v with %v operator\n", a, b, operator)
	}

}
