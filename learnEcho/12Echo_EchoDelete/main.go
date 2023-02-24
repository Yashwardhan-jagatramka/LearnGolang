package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/courses/:id", DeleteHandFunc)
	e.Logger.Fatal(e.Start("localhost:7000"))
}

// to delete again we have to check for availability of data as we did in get
func DeleteHandFunc(e echo.Context) error {
	courses := []map[int]string{{1: "GOLANG"}, {2: "ECHO"}, {3: "C++"}}
	var course map[int]string
	// we need index of the element that needs to be deleted
	var index int
	for i, c := range courses {
		for k := range c {
			cID, err := strconv.Atoi(e.Param("id"))
			if err != nil {
				return err
			}
			if cID == k {
				course = c
				//when we find the data to be deleted we store its index
				index = i
			}
		}
	}
	if course == nil {
		return e.JSON(http.StatusNotFound, "Sorry!. This Course is not Available")
	}
	//after checking availability and storing the index where data to be deleted is present we delete the data
	//here i created my own method since this is not inbulid for maps
	mapSlice := func(s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
	}
	courses = mapSlice(courses, index)
	return e.JSON(http.StatusOK, course)
}
