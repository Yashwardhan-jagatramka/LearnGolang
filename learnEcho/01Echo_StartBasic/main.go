package main

import (
	"fmt"
	"net/http"
)

//now net/http means there is a package called net inside it we have http
//http helps us to do web based tasks

type MyWebServeHttp bool

//now we have created MyWebServeHttp and then implemented same method as of handler,
// so we can use it as handler

func (m MyWebServeHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello we are learning echo. and Request is = %+v", r)
	//this method uses responseWriter to write my message and also showing my Request
}
func main() {
	//now ListenAndServe is the starting point to setup http web request
	// ListenAndServe takes to parameter first is string and second is handler
	// with string we pass the url and handler takes care of all the operations/web requests
	var k MyWebServeHttp
	http.ListenAndServe("localhost:8080", k)
	// ListenAndServe runs infinite loop to keep accepting the requests

	//now to understand handler we need to see how its implemented
	//type handler interface{
	//		ServeHttp(ResponseWriter, *Requests)
	//}
	// this is how handler was implemented now if we create any thing which have same method as in handler then we can use it as handler
	// example in line 8

	//now to use it we need to declare it see line 24 and then pass it in listen and serve as handler
}
