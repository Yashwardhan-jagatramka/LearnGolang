package main

import "fmt"

func greet(){
	fmt.Println("Hello")
}

func adder(a int, b int) int {
	return a+b
}

func proAdder(values ...int) int{
	total := 0

	for _,val := range values{
		total += val
	}

	return total
}

func main() {
	fmt.Println("Welcome to functions")

	greet()
	result := adder(3,5)
	proRes := proAdder(3,5,6,7,8,9)

	fmt.Println(result)
	fmt.Println(proRes)
}