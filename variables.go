package main

import (
	"fmt"
)

//single line comment

/*
multi line
comment
*/
func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)
	//se le variabili non vengono usate viene segnalato un errore dal compilatore

	//use implicit initialization syntax --> imply data type by assegnation
	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	//Multiple assignment
	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
