package authUserInterface

import "zeus/user"

type ClientInfoInterface interface {
	UserInit() *user.UserInfo
}
