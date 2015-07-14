package dw

import "fmt"

func Example() {
	counter := 5
	factorial := 1

	Do(func() {
		factorial *= counter
		counter--
	}).While(func() bool { return counter > 0 })

	fmt.Println("factorial of 5 is", factorial)

	// output: factorial of 5 is 120
}
