package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {

	u1 := User{
		ID:        1,
		FirstName: "Danika",
		LastName:  "Mori",
	}
	u2 := User{
		ID:        2,
		FirstName: "Beppe",
		LastName:  "Maniglia",
	}

	if u1.ID == u2.ID {
		println("Same User!")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user.")
	} else {
		println("Different users!")
	}

}
