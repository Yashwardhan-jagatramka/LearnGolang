package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	//port := os.Getenv("MY_APP_PORT")
	//if port == "" {
	//	port = "8080"
	//}
	//this can be used to set a default port from the code if MY_APP_PORT is not available
	// we can set MY_APP_PORT in the terminal
	//to use this we need to start server with this code
	//e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
	e := echo.New()
	//created new echo

	//starting get
	e.GET("/courses/:id", GetHandFunc)
	e.Logger.Fatal(e.Start("localhost:7000"))
	// this is to start the server

}

func GetHandFunc(e echo.Context) error {
	courses := []map[int]string{{1: "GOLANG"}, {2: "ECHO"}, {3: "C++"}}
	//this is the data in map form
	var course map[int]string
	//created a var course which will contain only one key/value pair

	//in this loop we are only taking key/value in c
	for _, c := range courses {
		//in this loop we are taking key in k from c
		for k := range c {
			cID, err := strconv.Atoi(e.Param("id"))
			//in this cID will have that int which user has given as input
			//e.Param return string so we converted it into int

			//handling error
			if err != nil {
				return err
			}

			//here we check if cID is equal to k if yes we put the value of c in course
			if cID == k {
				course = c
			}
		}
	}
	// after looping through every possible case and course is still empty then it means used inputed id is not present in data
	if course == nil {
		return e.JSON(http.StatusNotFound, "Sorry!. This Course is not Available")
	}

	//if course is not empty we return it
	return e.JSON(http.StatusOK, course)
}
