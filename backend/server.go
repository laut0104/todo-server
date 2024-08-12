package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	e.GET("/test", test)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}

func test(c echo.Context) error {
	println("test")
	return c.NoContent(http.StatusOK)
}
