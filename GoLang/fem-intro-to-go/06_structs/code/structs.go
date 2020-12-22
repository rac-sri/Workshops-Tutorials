package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

type Group struct {
	role string
	users []User
	newwestUser ser
	spaceAvailable bool
}

// // User is a user type
// type User struct {
// ID					       int
// FirstName, LastName, Email string
// }

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}

	fmt.Println(u.FirstName)
	fmt.Println(u)
}
