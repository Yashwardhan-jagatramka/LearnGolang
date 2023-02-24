package Courses

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

func Start() {
	port := os.Getenv("MY_COURSES_ECHO")
	if port == "" {
		port = "7000"
	}

	e.GET("/courses", getCourses)
	e.GET("/courses/:id", getCourse)
	e.DELETE("/courses/:id", deleteCourses)
	e.PUT("/courses/:id", updateCourses)
	e.POST("/courses", createCourses)

	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
