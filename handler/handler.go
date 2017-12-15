/**
Don't Edit
*/

package handler

import (
	"net/http"

	"github.com/gericass/go-spotify/gotify"
	"github.com/gericass/go-spotify/gotify/response"

	"github.com/labstack/echo"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expire_in"`
	RefreshToken string `json:"refresh_token"`
}

var Auth *gotify.Client

type Got struct {
	Gt *gotify.Client
}

func TestCallbackHandler(c echo.Context) error {
	token, err := Auth.CallbackHandler(c.Request())
	if err != nil {
		return err
	}

	res, err := response.GetUserStatus(token)
	if err != nil {
		return nil
	}
	return c.String(http.StatusOK, res)
}
