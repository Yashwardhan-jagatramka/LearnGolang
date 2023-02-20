package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number")

	input, _ := reader.ReadString('\n')
	fmt.Println("Number is :", input)
}