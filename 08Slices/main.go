package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
//---------------Way 1----------------//
	var fruits = []string{"Apple", "Tomato", "Peach"}
	fmt.Println(fruits)

	fruits = append(fruits, "Mango", "Banana")
	fmt.Println(fruits)
//---------------Way 2-----------------//
	
	highScores := make([]int,4)
	highScores[0] = 345
	highScores[1] = 567
	highScores[2] = 876
	highScores[3] = 761

	highScores = append(highScores, 666,555,777)

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

//-----------remove element from specific index-----------
	var courses = []string{"reactjs","javascript","kotlin","swift","golang"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}