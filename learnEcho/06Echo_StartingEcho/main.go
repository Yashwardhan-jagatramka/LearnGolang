package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/get", gethandler)
	// lets understand this e.GET()
	//to find out how its implemented ctrl + click on it
	//it needs 3 parameters first is path,handlerfunc ,middleware
	//path is - its a string type and the string which we pass is when call runs the handlerfunc.

	e.Logger.Fatal(e.Start(":8080"))
	//we can just use e.start(":8080") to run the http server
	//but using logger.fatal is good practice because when some errors come then it stops the process after addressing it

}

func gethandler(c echo.Context) error {
	return c.String(http.StatusOK, "Echo Get HandlerFunc")
}
