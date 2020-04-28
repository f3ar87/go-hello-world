package main

import (
	"fmt"
)

func main() {
	const pi float64 = 3.1415 //constant must be initialize when declared, cannot be a function, it must be evaluated when declared
	const pi2 = 3.1415

	fmt.Println(pi)

	//pi = 1 --> can't assign value to a constant

	const c = 3

	fmt.Println(c + 3)   //implicitly type constraint
	fmt.Println(c + 1.2) //implicitly type constraint

	const d int = 3

	fmt.Println(d + 3)
	fmt.Println(float32(d) + 1.2)
}
