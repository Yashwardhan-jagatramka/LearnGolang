package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	//var ptr *int
	//fmt.Println(ptr) -----> this will give us <nil>

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Address of myNumber is:", ptr)
	fmt.Println("Value stored in myNumber is: ",*ptr)

	*ptr = *ptr + 2
	fmt.Println("Changed value in pointer is : ", *ptr)
}