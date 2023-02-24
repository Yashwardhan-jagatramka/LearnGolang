package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}
func main() {
	var courses = []map[int]string{{1: "GOLANG"}, {2: "ECHO"}, {3: "C++"}}
	var e = echo.New()
	e.POST("/courses", func(ee echo.Context) error {
		v := validator.New()
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
