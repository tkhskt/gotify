/**
Don't Edit
*/

package main

import (
	"github.com/gericass/go-spotify/gotify"
	"github.com/gericass/go-spotify/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const clientID = ""
const clientSecret = ""
const callbackURI = "https://localhost:3000/callback/"

func (g *handler.Got) first(c echo.Context) error {
	g.Gt.AuthURL()
	return nil
}

func callback(c echo.Context) error {

	return nil
}

func main() {
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	got := &handler.Got{gotify.Set(clientID, clientSecret, callbackURI)}

	e.GET("/", got.first)
	e.GET("/callback/", handler.CallbackHandler)

	//HTTPSを使うのに必要！
	e.Logger.Fatal(e.StartTLS(":3000", "cert.pem", "key.pem"))
}
