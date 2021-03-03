package user

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
	"zeus/keycloak"

	"github.com/gin-gonic/gin"
)

func Register_user(c *gin.Context, r *http.Request, w http.ResponseWriter) {
	now := time.Now().Local()

	strTime, err := getCookie(r, "access_token_expired_time")
	expiredTime, _ := time.Parse(time.RFC3339, strTime)

	refresh_strTime, err := getCookie(r, "refresh_token_expired_time")
	refresh_expiredTime, _ := time.Parse(time.RFC3339, strTime)

	if expiredTime.After(now) {
		fmt.Println("after expired time")

		if refresh_expiredTime.After(now) {
			keycloak.GetAccessToken(w, "password")
		}

	} else {
		fmt.Println("before expired time")

		accessToken, _ := getCookie(r, "access_token")
	}

	//keycloak.GetAccessToken("password")

	bindUser := InitUserInfo()
	c.BindJSON(bindUser)
}

func getCookie(r *http.Request, cookieName string) (token string, err error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return
	}
	tokb, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		return
	}
	token = string(tokb)
	return
}
