package Packages

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) Walk() {
	fmt.Println(p.FirstName + " " + p.LastName + " is Walking")
}
