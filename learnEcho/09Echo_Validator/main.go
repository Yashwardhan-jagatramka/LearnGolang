package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

var courses = []map[int]string{{1: "GOLANG"}, {2: "ECHO"}, {3: "C++"}}

// now since we are taking input from user we need to validate it before processing it
// so that our system is free from attacks
// so i m using validator in previous post example function
func main() {
	e := echo.New()
	e.POST("/courses", PostHandFunc)
}
func PostHandFunc(e echo.Context) error {
	// we need to declare this for validation
	v := validator.New()
	//to use validator we need to implement it in Post Function -- look at PostHandFunc
	type body struct {
		// to use validator we need to do this
		Name string `json:"course_name" validate:"required,min=4"`
	}
	var reqBody body
	if err := e.Bind(&reqBody); err != nil {
		return err
	}
	//we call validator here
	if err := v.Struct(reqBody); err != nil {
		return err
	}
	course := map[int]string{
		len(courses) + 1: reqBody.Name,
	}
	courses = append(courses, course)
	return e.JSON(http.StatusOK, course)
}
