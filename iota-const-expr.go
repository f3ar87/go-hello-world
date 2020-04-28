package main

import "fmt"

const (
	first  = 1
	second = "second"

	uno      = iota
	due      = iota + iota + iota
	tre      = 2 << iota //00110101<<2	11010100	Left shift --> in questo caso dovrebbe essere 10 << 4 (iota parte da 0) --> 100000 = 32
	quattro  = iota + 2
	quattro2 //se non specifico nulla riusa la stringa sopra, ma in questo caso iota cambia
)

const (
	nuovo1 = iota
	nuovo2
	nuovo3
)

func main() {
	fmt.Println(first, second)
	fmt.Println(uno, due, tre)
	fmt.Println(quattro, quattro2)
	fmt.Println(nuovo1, nuovo2, nuovo3)
}
