package main

import "fmt"

func main() {
	slice := []int{1, 2, 3} //in this way is Go managing the underline array
	fmt.Println(slice)

	slice = append(slice, 4) //Go create a new underline array, then copy the data from the previous one with a new element
	fmt.Println(slice)

	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)

	s2 := slice[1:]
	fmt.Println(s2)

	s3 := slice[:2]
	fmt.Println(s3)

	s4 := slice[1:2] //not including index2
	fmt.Println(s4)
}
