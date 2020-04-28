package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3} //value-type var

	slice := arr[:] //kind of a pointer; used to work with array in a more dynamically manner

	fmt.Println(arr, slice)

	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice)
}
