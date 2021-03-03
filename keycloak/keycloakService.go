package keycloak

import (
	"encoding/base64"
	"net/http"
	"strconv"
)

func GetAccessToken(w http.ResponseWriter, grantType string) {

	h_accessToken := InitAccessTokenInfo()
	r_accessToken := InitResponseAccessTokenInfo()

	res := h_accessToken.request_AccessToken(grantType)

	expired_in := r_accessToken.parsing_Response_Expires_in(res)
	refresh_expires_in := r_accessToken.parsing_Response_RefreshExpires_in(res)
	access_token := r_accessToken.parsing_Response_AccessToken(res)
	//refresh_token := r_accessToken.parsing_Response_RefreshToken(res)
	setTokenCookie(w, access_token, "access_token")
	setCookie(w, expired_in, "access_token_expired_time")
	setCookie(w, refresh_expires_in, "refresh_token_expired_time")
}

func setTokenCookie(w http.ResponseWriter, value string, cookieName string) {
	tok64 := base64.StdEncoding.EncodeToString([]byte(value))

	cookie := http.Cookie{
		Name:     cookieName,
		Value:    tok64,
		HttpOnly: true,
		Secure:   false, //use true for production
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	return
}

func setCookie(w http.ResponseWriter, expired_time int, cookieName string) {
	tok64 := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(expired_time)))

	cookie := http.Cookie{
		Name:     cookieName,
		Value:    tok64,
		HttpOnly: true,
		Secure:   false, //use true for production
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	return
}
