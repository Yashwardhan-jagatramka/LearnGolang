package main

import "fmt"

func main() {
	fmt.Println("Welcome to for Loop")

	for i:=1;i<10;i++{
		fmt.Println(i)
	}
	
	slice := make([]string,5)

	slice[0] = "abc"
	slice[1] = "bcd"
	slice[2] = "cde"
	slice[3] = "def"
	slice[4] = "efg"

	for i := range slice{
		fmt.Println(slice[i])
	}

	for i,alp := range slice{
		fmt.Printf("Index is %v and value is %v.\n",i,alp)
	}

	//-----creating while loop using for loop
	j := 1
	for j < 10{
		//break condition
		// if j==2 {
		// 	break
		// }
		//continue condition
		// if j==2 {
		// 	j++
		// 	continue
		// }
		fmt.Print(j," ")
		j++
	}
}