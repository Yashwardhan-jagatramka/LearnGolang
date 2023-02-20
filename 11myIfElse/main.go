package main

import "fmt"

func main() {
	fmt.Println("Welcome to If/Else")

	loginCounts := 23

	var result string

	if loginCounts < 10{
		result = "Not Active"
	}else if loginCounts > 10{
		result = "Very Active"
	}else{
		result = "Just Active"
	}

	fmt.Println(result)

	if 9%2==0 {
		fmt.Println("9 is even")
	}else{
		fmt.Println("9 is odd")
	}

	if a:=3 ; a < 10{
		fmt.Println("a is less that 10")
	}else{
		fmt.Println("a is greater than 10")
	}
}