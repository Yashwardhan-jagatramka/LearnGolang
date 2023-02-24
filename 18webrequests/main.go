package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web requests")

}
func PerformGetRequest() {
	const myurl = "http://localhost:1000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status Code is: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	responseString.Write(content)

	//fmt.Println(bytecount)
	fmt.Println(responseString.String())
}