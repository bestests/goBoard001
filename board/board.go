package board

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

var (
	boardsMap   = make(map[int]*Board)
	boardsSlice = make([]*Board, 0)
)

type (
	// Board - board type
	Board struct {
		Idx     int
		Title   string
		Content string
		RegDate string
		ModDate string
	}

	// Result - response result
	Result struct {
		Cnt  int
		Data []*Board
	}
)

// NowTime - get now time at Seoul
func NowTime() (string, time.Time) {
	t := time.Now()

	loc, _ := time.LoadLocation("Asia/Seoul")

	t = t.In(loc)

	str := t.Format("2006-01-02 15:04:05.000")

	return str, t
}

// NewBoard - create new board
func newBoard(idx int, title, content string) *Board {

	timeStr, _ := NowTime()

	return &Board{
		Idx:     idx,
		Title:   title,
		Content: content,
		RegDate: timeStr,
		ModDate: timeStr,
	}
}

func getResult() *Result {
	keys := make([]int, 0)

	for key := range boardsMap {
		keys = append(keys, key)
	}

	cnt := len(keys)

	if cnt == 0 {
		return &Result{
			Cnt:  0,
			Data: nil,
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	bSlice := make([]*Board, 0)

	for _, key := range keys {
		bSlice = append(bSlice, boardsMap[key])
	}

	return &Result{
		Cnt:  cnt,
		Data: bSlice,
	}
}

// ListHandler - list page
func ListHandler(c echo.Context) error {
	return c.File("views/board/list.html")
}

// GetListHandler - get list data
func GetListHandler(c echo.Context) error {

	boardsMap[1] = newBoard(1, "title1", "content1")
	boardsMap[2] = newBoard(2, "title2", "content2")
	boardsMap[3] = newBoard(3, "title3", "content3")

	result := getResult()

	return c.JSON(http.StatusOK, result)
}

// ViewHandler - view page
func ViewHandler(c echo.Context) error {
	return c.File("views/board/view.html")
}

// GetBoardHandler - get a board
func GetBoardHandler(c echo.Context) error {
	idxStr := c.QueryParam("idx")
	idx, _ := strconv.Atoi(idxStr)

	return c.JSON(http.StatusOK, boardsMap[idx])
}
