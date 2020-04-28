package main

import "fmt"

func main() {
	//with struct we must first define and then initialize
	//fields are fixed at runtime
	//posso definire struct anche fuori dal main
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user

	fmt.Println(u) //field are initialize to their zero-value; int = 0, string = ""

	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)

	fmt.Println(u.ID)
	fmt.Println(u.FirstName)
	fmt.Println(u.LastName)

	u2 := user{ID: 2, FirstName: "Den", LastName: "Morris"}
	fmt.Println(u2)

	u3 := user{ID: 3,
		FirstName: "Danika",
		LastName:  "Mori", //occhio alla virgola che deve essere messa per forza, altrimenti avrei dovuto mettere la parentesi graffa qui
	}
	fmt.Println(u3)
}
