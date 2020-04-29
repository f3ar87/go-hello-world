package main

func main() {
	var i int

	for i < 5 {
		i++
		println(i)
		if i == 1 {
			continue
		}
		if i == 3 {
			break
		}
		println("Continuing..")
	}
}
