// --------HttpMethods
package main

import (
	"fmt"
	"net/http"
)

// now to achieve different kind of behaviour for different methods like GET, POST....etc
// we can apply this type of conditions
// what it does is it checks the request method and behave accordingly
// below is very simple example
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "Hello Get from Login")
	}
	if r.Method == "POST" {
		fmt.Fprintln(w, "Hello post from Login")
	}
}

func main() {
	http.HandleFunc("/login", login)
}
