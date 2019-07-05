package main

import (
	"net/http"
  "strconv"
  "github.com/labstack/echo"
)

func main() {
  e := echo.New()

  e.GET("/", func(c echo.Context) error {
		var string = "Hello World!"
    return c.JSON(http.StatusOK, string)
  })

  e.GET("/reiwa", func(c echo.Context) error {
		var string = "リクエストパラメータで西暦を入れてね。"
    return c.JSON(http.StatusOK, string)
  })

  e.GET("/reiwa/:seireki", func(c echo.Context) error {
		var seireki, _ = strconv.Atoi(c.Param("seireki"))
		var string = ""
		if seireki < 2019 {
			string = "西暦"+strconv.Itoa(seireki)+"年 は 令和 ではありません。"
		} else {
			var reiwa = strconv.Itoa(seireki - 2018);
			string = "西暦"+strconv.Itoa(seireki)+"年 は 令和 "+reiwa+"年 です。"
		}
    return c.JSON(http.StatusOK, string)
  })

  e.Logger.Fatal(e.Start(":3000"))
}
