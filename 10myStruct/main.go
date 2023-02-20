package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to struct")

	Yash := User{"Yash" ,"Yash@go.dev",true,22}

	fmt.Println(Yash)
	fmt.Printf("%+v\n",Yash)
	fmt.Printf("Name is %v." ,Yash.Name)

}
