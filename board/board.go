package board

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var boards = make(map[int]*Board)

// Board - board type
type Board struct {
	Idx     int
	Title   string
	Content string
	RegDate string
	ModDate string
}

// NowTime - get now time at Seoul
func NowTime() (string, time.Time) {
	t := time.Now()

	loc, _ := time.LoadLocation("Asia/Seoul")

	t = t.In(loc)

	str := t.Format("2006-01-02 15:04:05.000")

	return str, t
}

// NewBoard - create new board
func NewBoard(idx int, title, content string) *Board {

	timeStr, _ := NowTime()

	return &Board{
		Idx:     idx,
		Title:   title,
		Content: content,
		RegDate: timeStr,
		ModDate: timeStr,
	}
}

// ListHandler - list
func ListHandler(c echo.Context) error {
	return c.File("views/board/list.html")
}

// GetListHandler - get list data
func GetListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, boards)
}
