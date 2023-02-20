package main

import (
	"awesomeProject/Packages"
	"awesomeProject/Polymorphism"
)

func main() {
	p := Packages.Person{
		FirstName: "Yashwardhan",
		LastName:  "Jagatramka",
		Age:       22,
	}
	p.Walk()

	var c Polymorphism.Circle = Polymorphism.Circle{}
	var r Polymorphism.Rectangle = Polymorphism.Rectangle{}

	c.SetR(5)
	r.SetLB(10, 10)

	c.Area()
	r.Area()

}
