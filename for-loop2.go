package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}
	//println(i) --> Not valid to use i here

	var j int

	for ; j < 5; j++ {
		println(j)
	}
	println(j) // this way is working
}
