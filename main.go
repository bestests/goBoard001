package main

import (
	"github.com/goBoard001/board"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func jsHandler(c echo.Context) error {
	fileNm := c.Param("fileNm")
	return c.File("js/" + fileNm)
}

func cssHandler(c echo.Context) error {
	fileNm := c.Param("fileNm")
	return c.File("css/" + fileNm)
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
	// JS FILE
	e.GET("/js/:fileNm", jsHandler)
	// CSS FILE
	e.GET("/css/:fileNm", cssHandler)

	e.GET("/", homeHandler)

	// Board
	e.GET("/board/list", board.ListHandler)
	e.GET("/board/getList", board.GetListHandler)
	e.GET("/board/view/:idx", board.ViewHandler)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
