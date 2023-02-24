package main

import "fmt"

////---------Global constants
//const is use to define constant
//if the name of that starts with capital letter than it is exportable and if not it can't be exported outside
const LoginToken string = "abc"

// to define miltiple const we can use this syntax it also work for variables and imports(in go generally imports are automatic)
// const(
// 	abc string
// 	bcd string
// )
////------GLOBAL VARIABLES------
//Global variables are defined outside of a function, usually on top of the program
var glob int = 22

func main() {

	fmt.Println(glob)
// string
	var username string = "Yashwardhan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
// bool
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
//int
	var LoginToken int = 1234
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
// now the best part like in python we don't need to write big lines to declare + initilaize the variables 
//similarly we can do it in go
	uName := "Yash"
	fmt.Println(uName)
//like this go automatically read the data and decides its datatype


// and if we declare any variable or const in program and dont use it it will show error

}
