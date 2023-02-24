package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var courses = []map[int]string{{1: "GOLANG"}, {2: "ECHO"}, {3: "C++"}}

func main() {
	e := echo.New()
	e.GET("/courses/:id", GetHandFunc2)
	e.POST("/courses", PostHandFunc)
	e.Logger.Fatal(e.Start("localhost:7000"))
}
func GetHandFunc2(e echo.Context) error {

	var course map[int]string
	for _, c := range courses {
		for k := range c {
			cID, err := strconv.Atoi(e.Param("id"))
			if err != nil {
				return err
			}
			if cID == k {
				course = c
			}
		}
	}
	if course == nil {
		return e.JSON(http.StatusNotFound, "Sorry!. This Course is not Available")
	}
	return e.JSON(http.StatusOK, course)
}

// starting post
func PostHandFunc(e echo.Context) error {

	//creating a body struct with Name(string)
	//we created this body struct because when we use post method we need to pass a particular format.
	type body struct {
		Name string `json:"course_name"`
		//json:"course_name" will get maped to Name when body struct is converted into json
	}
	// we declare reqBody of type body to take user request
	var reqBody body
	// now we bind the user request to reqBody
	// note- we pass pointer of reqBody to take user request else it won't work
	// since bind return error we handle error if any
	if err := e.Bind(&reqBody); err != nil {
		return err
	}
	//now we declare course and populate it with data user provided
	course := map[int]string{
		len(courses) + 1: reqBody.Name,
	}
	// and then appending it to courses
	courses = append(courses, course)
	//after successfully appending it we return course not courses because user requested for course
	return e.JSON(http.StatusOK, course)
}
