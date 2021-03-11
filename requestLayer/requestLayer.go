package requestLayer

import (
	"context"
	"net/http"
	"zeus/models"

	"golang.org/x/oauth2"
)

type RequestLayer interface {
	//GetClient(context.Context) (*http.Client, error)
	GetApiClientTokenSource(context.Context) *oauth2.Token
	RequestRegisterUserApi(context.Context, models.UserInfo, *http.Client) error
}
