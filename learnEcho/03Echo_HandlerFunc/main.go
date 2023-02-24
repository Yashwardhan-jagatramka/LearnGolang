// -----handler function
// as we created handler of our own type similarly we can create handler function also
package main

import (
	"fmt"
	"net/http"
)

// LAS = ListenAndServe
// we created HandFunc which has ResponseWriter and Request so this can be pased as handler function in LAS
// to pass it as handler we use http.handlerFunc(name of function we created).
// http.handlerFunc is used to convert Handler Function into Handler
func HandFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello we are learning echo. and Request is = %+v", r)
}

// so now the there is one problem, with different url we are recieving same output,
// localhost:8080/login is same as localhost:8080/welcome
// so to solve this
// we created two different handler function down below
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
	//and then to execute it we have to call it as i call it below
	http.HandleFunc("/login", login)
	http.HandleFunc("/welcome", welcome)
	//now as i passed /welcome as parameter then after localhost:8080/welcome welcome function will be executed
	//but if we run localhost:8080/welcome/yash then it will show error 404
	// to solve this we can pass /welcome/ in HandleFunc

	//http.ListenAndServe("localhost:8080", http.HandlerFunc(HandFunc))
	http.ListenAndServe("localhost:8080", nil)
}
