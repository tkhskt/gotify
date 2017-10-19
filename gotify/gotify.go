package gotify

import (
	"github.com/labstack/echo"
	"github.com/go-spotify/handler"
)



func OAuthHandler(clientID string, clientSecret string, callbackURL string) (t *Tokens){
	e := echo.New()
	e.StartTLS(":3000", "cert.pem", "key.pem")
	e.GET(callbackURL, handler.CallbackHandler)

}
