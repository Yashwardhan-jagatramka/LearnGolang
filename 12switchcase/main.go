package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switchCase")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Move 1 block")
	case 2:
		fmt.Println("Move 2 block")
	case 3:
		fmt.Println("Move 3 block")
	case 4:
		fmt.Println("Move 4 block")
	case 5:
		fmt.Println("Move 5 block")
	case 6:
		fmt.Println("Move 6 block & roll again")
	default:
		fmt.Println("What was that")
	}
}
