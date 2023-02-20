package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer")

	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}