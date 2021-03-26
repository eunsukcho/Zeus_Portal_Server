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
	RequestOneUserApi(context.Context, string, *http.Client) (models.ResponseUserInfo, error)
	RequestRegisterUserApi(context.Context, models.RegisterUserInfo, *http.Client) (string, error)
	DeleteUserApi(context.Context, string, *http.Client) (string, error)
	UpdateUserApi(context.Context, models.RegisterUserInfo, *http.Client) (string, error)
	UpdateUserCredentialsApi(context.Context, string, *http.Client) (string, error)

	RequestGroupListApi(context.Context, *http.Client) ([]models.ResGroupInfo, error)
	RequestRegisterGroupsApi(context.Context, models.ReqToken, *http.Client) (string, error)
}
