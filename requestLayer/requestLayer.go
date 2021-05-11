package requestLayer

import (
	"context"
	"zeus/models"

	"golang.org/x/oauth2"
)

type RequestLayer interface {
	GetApiClientTokenSource(context.Context) *oauth2.Token
	RequestUserListApi(context.Context, *oauth2.Token) ([]models.ResponseUserInfo, error)
	RequestUserListByGroupApi(context.Context, string, *oauth2.Token) ([]models.ResponseUserInfo, error)
	RequestOneUserApi(context.Context, string, *oauth2.Token) (models.ResponseUserInfo, error)
	RequestRegisterUserApi(context.Context, models.RegisterUserInfo, *oauth2.Token) (string, int, error)
	DeleteUserApi(context.Context, string, *oauth2.Token) (string, error)
	UpdateUserApi(context.Context, models.RegisterUserInfo, *oauth2.Token) (string, error)
	UpdateUserCredentialsApi(context.Context, string, *oauth2.Token) (string, error)

	RequestGroupListApi(context.Context, string, *oauth2.Token) ([]models.ResGroupInfo, error)
	RequestRegisterGroupsApi(context.Context, models.ReqToken, *oauth2.Token) (string, error)
}
