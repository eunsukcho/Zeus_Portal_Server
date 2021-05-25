package requestLayer

import (
	"context"
	"zeus/models"

	"golang.org/x/oauth2"
)

type RequestLayer interface {
	GetApiClientTokenSource(context.Context) *oauth2.Token
	RequestUserListApi(context.Context, *oauth2.Token) ([]models.ResponseUserInfo, int, error)
	RequestUserListByGroupApi(context.Context, string, *oauth2.Token) ([]models.ResponseUserInfo, int, error)
	RequestOneUserApi(context.Context, string, *oauth2.Token) (models.ResponseUserInfo, int, error)
	RequestRegisterUserApi(context.Context, models.RegisterUserInfo, *oauth2.Token) (string, int, error)
	DeleteUserApi(context.Context, string, *oauth2.Token) (string, int, error)
	UpdateUserApi(context.Context, models.RegisterUserInfo, *oauth2.Token) (string, int, error)
	UpdateUserCredentialsApi(context.Context, string, *oauth2.Token) (string, int, error)

	RequestGroupListApi(context.Context, string, *oauth2.Token) ([]models.ResGroupInfo, int, error)
	RequestRegisterGroupsApi(context.Context, models.ReqToken, *oauth2.Token) (string, int, error)
}
