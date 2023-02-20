package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to Methods ")

	Yash := User{"Yash" ,"Yash@go.dev",true,22}

	fmt.Println(Yash)
	fmt.Printf("%+v\n",Yash)
	fmt.Printf("Name is %v.\n" ,Yash.Name)
	Yash.GetStatus()
	Yash.NewEmail()

}

func (u User) GetStatus(){
	fmt.Println("Is user Active :", u.Status)
}

func (u User) NewEmail(){
	u.Email = "test@go.dev"
	fmt.Println("New email for this user is : ", u.Email)
}