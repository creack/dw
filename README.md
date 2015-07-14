# Do While loop

## Example

```go
package main

import "github.com/creack/dw"

func main() {
	counter := 5
	factorial := 1

	dw.Do(func() {
		factorial *= counter
		counter--
	}).While(func() bool { return counter > 0 })

	println("factorial of 5 is", factorial)
}
```
