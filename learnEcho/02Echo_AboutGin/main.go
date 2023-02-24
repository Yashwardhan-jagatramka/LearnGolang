package main

import (
	"fmt"
	"net/http"
)

type MyWebServeHttp bool

func (m MyWebServeHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello we are learning echo. and Request is = %+v", r)
}
func main() {
	var k MyWebServeHttp
	http.ListenAndServe("localhost:8080", k)
}

//This is previous example which we saw in StartBasic but when we are creating it we saw after any update in code we need to rerun,
//now to tackle this we use Gin
