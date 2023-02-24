package main

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

//now if we want to edit data that is already present we need to use PUT method
func main() {
	courses := []map[int]string{{1: "GOLANG"}, {2: "ECHO"}, {3: "C++"}}
	e := echo.New()
	v := validator.New()

	//as we check for the avability of data in get same we need to check in put,
	// because if data if not present then how can we edit it
	e.PUT("/courses/:id", func(ee echo.Context) error {
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
	})
	e.POST("/courses", func(ee echo.Context) error {
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
		//if err := v.Struct(reqBody); err != nil {
		//	return err
		//}
		course := map[int]string{
			len(courses) + 1: reqBody.Name,
		}
		courses = append(courses, course)
		return ee.JSON(http.StatusOK, course)
	})
}
