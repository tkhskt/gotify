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

func first(c echo.Context) error {
	url := handler.Auth.AuthURL()
	return c.Redirect(301, url)
}

func callback(c echo.Context) error {

	return nil
}

func main() {
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.Auth = gotify.Set(clientID, clientSecret, callbackURI)

	e.GET("/", first)
	e.GET("/callback/", handler.TestCallbackHandler)

	//HTTPSを使うのに必要！
	e.Logger.Fatal(e.StartTLS(":3000", "cert.pem", "key.pem"))
}
