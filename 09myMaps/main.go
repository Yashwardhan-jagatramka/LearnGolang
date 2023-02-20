package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["GO"] = "Golang"
	languages["PY"] = "Python"
	languages["JS"] = "JavaScript"

	fmt.Println(languages)
	fmt.Println(languages["GO"])

	delete(languages, "JS")
	fmt.Println(languages)
}