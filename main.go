/**
 Don't Edit
 */

package main

import (
	"github.com/go-spotify/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.OAuthHandler)
	e.GET("/callback/", handler.CallbackHandler)

	//HTTPSを使うのに必要！
	e.Logger.Fatal(e.StartTLS(":3000", "cert.pem", "key.pem"))
}
