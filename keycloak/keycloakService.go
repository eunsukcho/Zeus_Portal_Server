package keycloak

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetAccessToken(grantType string, get_refresh_token string) (string, string, int, int) {

	//h_accessToken := InitAccessTokenInfo()
	r_accessToken := InitResponseAccessTokenInfo()

	res := request_AccessToken(grantType, get_refresh_token)

	expired_in := r_accessToken.parsing_Response_Expires_in(res)
	refresh_expires_in := r_accessToken.parsing_Response_RefreshExpires_in(res)
	access_token := r_accessToken.parsing_Response_AccessToken(res)
	refresh_token := r_accessToken.parsing_Response_RefreshToken(res)

	return access_token, refresh_token, expired_in, refresh_expires_in
}

func request_AccessToken(grantType string, refresh_token string) string {

	data := url.Values{}
	data.Set("client_id", "admin-cli")
	data.Set("client_secret", "7eceed48-073d-4c47-bb30-aae22ac14366")
	data.Set("grant_type", grantType)
	data.Set("username", "admin")
	data.Set("password", "admin")

	if grantType == "refresh_token" && refresh_token != "" {
		data.Set("refresh_token", refresh_token)
	}

	req, err := http.NewRequest("POST", "http://192.168.0.118:9090/auth/realms/master/protocol/openid-connect/token", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		return str
	}
	return ""
}
