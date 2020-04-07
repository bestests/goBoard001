package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func jsHandler(c echo.Context) error {
	fileNm := c.Param("fileNm")
	return c.File("js/" + fileNm)
}

func homeHandler(c echo.Context) error {
	return c.File("views/index.html")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route
	e.GET("/js/:fileNm", jsHandler)

	e.GET("/", homeHandler)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
