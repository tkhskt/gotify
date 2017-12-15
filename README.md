# Go-Spotify

Spotify WEB API wrapper

## Require

- Go 1.9~

## Usage

`go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost`

`go get github.com/gericass/go-spotify`

## example

sample by using `echo`

```main.go
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

func main() {
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.Auth = gotify.Set(clientID, clientSecret, callbackURI)

	e.GET("/", handler.ExampleHandler)
	e.GET("/callback/", handler.ExampleCallbackHandler)

	// Require SSL
	e.Logger.Fatal(e.StartTLS(":3000", "cert.pem", "key.pem"))
}

```

```handler.go
package handler

import (
	"net/http"

	"github.com/gericass/go-spotify/gotify"
	"github.com/gericass/go-spotify/gotify/response"

	"github.com/labstack/echo"
)

var Auth *gotify.Client

func ExampleHandler(c echo.Context) error {
	url := Auth.AuthURL()
	return c.Redirect(301, url)
}

func ExampleCallbackHandler(c echo.Context) error {
	token, err := Auth.Token(c.Request())
	if err != nil {
		return err
	}

	res, err := response.GetUserStatus(token)
	if err != nil {
		return nil
	}
	return c.String(http.StatusOK, res)
}

```
