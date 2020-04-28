package main

import "fmt"

func main() {
	var firstName *string = new(string)

	*firstName = "Arthur"

	fmt.Println(firstName)  //0xc000010200
	fmt.Println(*firstName) //Arthur

	firstName2 := "Arthur2"
	fmt.Println(firstName2) //Arthur2

	ptr := &firstName2
	fmt.Println(ptr, *ptr) //0xc000010220 Arthur2

	firstName2 = "Tricia"
	fmt.Println(ptr, *ptr) //0xc000010220 Tricia
}
