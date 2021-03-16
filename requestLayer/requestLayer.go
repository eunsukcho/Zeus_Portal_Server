package requestLayer

import (
	"context"
	"net/http"
	"zeus/models"

	"golang.org/x/oauth2"
)

type RequestLayer interface {
	GetApiClientTokenSource(context.Context) *oauth2.Token
	RequestUserListApi(context.Context, *http.Client) ([]models.ResponseUserInfo, error)
	RequestRegisterUserApi(context.Context, models.RegisterUserInfo, *http.Client) (string, error)
}
