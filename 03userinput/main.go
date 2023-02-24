package main

import (
	"bufio"
	"fmt"
	"os"
)
// we have 2 method to take user input 
// first using os and bufio packages below is the example
func main() {
	fmt.Println("Welcome to user input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number")

	input, _ := reader.ReadString('\n')
	fmt.Println("Number is :", input)
}