package Courses

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var courses = []map[int]string{{1: "GOLANG"}, {2: "ECHO"}, {3: "C++"}}

func getCourses(c echo.Context) error {
	return c.JSON(http.StatusOK, courses)
}
func getCourse(e echo.Context) error {
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
func deleteCourses(e echo.Context) error {
	var course map[int]string
	var index int
	for i, c := range courses {
		for k := range c {
			cID, err := strconv.Atoi(e.Param("id"))
			if err != nil {
				return err
			}
			if cID == k {
				course = c
				index = i
			}
		}
	}
	if course == nil {
		return e.JSON(http.StatusNotFound, "Sorry!. This Course is not Available")
	}
	mapSlice := func(s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
	}
	courses = mapSlice(courses, index)
	return e.JSON(http.StatusOK, course)
}
func updateCourses(ee echo.Context) error {
	var course map[int]string
	cID, err := strconv.Atoi(ee.Param("id"))
	if err != nil {
		return err
	}
	for _, c := range courses {
		for k := range c {
			if cID == k {
				course = c
			}
		}
	}
	if course == nil {
		return ee.JSON(http.StatusNotFound, "Sorry!. Product not found")
	}
	//as we have done in post and validation we need to do it here also
	type body struct {
		Name string `json:"course_name" validate:"required,min=4"`
	}
	var reqBody body
	e.Validator = &ProductValidator{validator: v}
	if err := ee.Bind(&reqBody); err != nil {
		return err
	}
	if err := ee.Validate(reqBody); err != nil {
		return err
	}
	//after all validation we should edit it
	course[cID] = reqBody.Name
	//and at the end return the final edited version
	return ee.JSON(http.StatusOK, course)
}
func createCourses(e echo.Context) error {
	type body struct {
		Name string `json:"course_name"`
	}
	var reqBody body
	if err := e.Bind(&reqBody); err != nil {
		return err
	}
	course := map[int]string{
		len(courses) + 1: reqBody.Name,
	}
	courses = append(courses, course)
	return e.JSON(http.StatusOK, course)
}
