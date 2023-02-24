package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")
	//we are creating a file and then writing in it
	content := "Hello this is content of file"
	// content needs to be pushed in file

	file, err := os.Create("./myFirstGoFile.txt")
	defer file.Close()

	if err!= nil{
		panic(err)
	}

	length, err := io.WriteString(file,content)
	if err!= nil{
		panic(err)
	}

	fmt.Println("length is :",length)
	readFile(file.Name())
}

func readFile(filename string){
	data,err := ioutil.ReadFile(filename)
	if err!=nil{
		panic(err)
	}
	fmt.Println("The readable data in file is : ", string(data))
}