package main

import (
	"fmt"
	"net/http"
)

func HandFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello we are learning echo. and Request is = %+v", r)
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>
		   Login Page
		</head>
		<body>
			<h1>
			  Enter Login id and Password.
			</h1>
		</body>
	</html>
 `)
}
func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>
		   Welcome Page
		</head>
		<body>
			<h1>
			  Hello there!. Welcome to Echo.
			</h1>
		</body>
	</html>
 `)
}
func main() {
	//http.HandleFunc("/login", login)
	//http.HandleFunc("/welcome", welcome)
	//as we are calling handleFunc we can do same stuff with Handle also
	http.Handle("/login", http.HandlerFunc(login))
	http.Handle("/welcome", http.HandlerFunc(welcome))
	//its doing same stuff but HandleFunc takes Handle func as input parameter whereas Handle take Handler as input parameter

	http.ListenAndServe("localhost:8080", nil)
}
