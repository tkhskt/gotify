package test

import (
	"log"
	"testing"

	"github.com/gericass/gotify/gotify"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func TestAuth(t *testing.T) {
	clientID := "6c9f61b6ffb5420ca8995290b82a88a9"
	clientSecret := "982ddafad2a84ea190deb9fa7a1e426f"
	callbackURI := "https://localhost:3000/callback/"
	auth := gotify.Set(clientID, clientSecret, callbackURI)

	actual := auth.Auth()
	expected := "https://accounts.spotify.com/authorize/?client_id=" + auth.ClientID + "&response_type=" + "code" + "&redirect_uri=" + auth.CallbackURI +
		"&scope=user-read-private%20user-library-read%20user-follow-read"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func testAccess(c echo.Context) error {
	clientID := "6c9f61b6ffb5420ca8995290b82a88a9"
	clientSecret := "982ddafad2a84ea190deb9fa7a1e426f"
	callbackURI := "https://localhost:3000/callback/"
	auth := gotify.Set(clientID, clientSecret, callbackURI)
	req := c.Request()
	token, _ := auth.CallbackHandler(req)
	log.Println(token.AccessToken)
	return nil
}

func TestCallbackHandler(t *testing.T) {
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/callback/", testAccess)

	//HTTPSを使うのに必要！
	e.Logger.Fatal(e.StartTLS(":3000", "cert.pem", "key.pem"))

}
