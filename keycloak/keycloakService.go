package keycloak

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SetHeader(c *gin.Context) {
	Test()
	now := time.Now().Local()

	accessToken, err := c.Cookie("access_token")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Cookie Access Token 0 :", accessToken)

	if accessToken == "" {
		accessToken := GetAccessToken(c, "password")
		fmt.Println("Access Token 0 :", accessToken)
		return
	}

	strTime, _ := c.Cookie("access_token_expired_time")
	expiredTime, _ := time.Parse(time.RFC3339, strTime)

	refresh_strTime, _ := c.Cookie("refresh_token_expired_time")
	refresh_expiredTime, _ := time.Parse(time.RFC3339, refresh_strTime)

	if expiredTime.After(now) {
		fmt.Println("after expired time")

		accessToken := GetAccessToken(c, "refresh_token")
		fmt.Println("Access Token 1 :", accessToken)

		if refresh_expiredTime.After(now) {
			accessToken := GetAccessToken(c, "password")
			fmt.Println("Access Token 2 :", accessToken)
		}

		return
	} else {
		fmt.Println("before expired time")

		accessToken, _ := c.Cookie("access_token")
		fmt.Println("Access Token : ", accessToken)

		return
	}

}
func GetAccessToken(c *gin.Context, grantType string) string {

	//h_accessToken := InitAccessTokenInfo()
	r_accessToken := InitResponseAccessTokenInfo()

	res := request_AccessToken(grantType, c)

	expired_in := r_accessToken.parsing_Response_Expires_in(res)
	refresh_expires_in := r_accessToken.parsing_Response_RefreshExpires_in(res)
	access_token := r_accessToken.parsing_Response_AccessToken(res)
	refresh_token := r_accessToken.parsing_Response_RefreshToken(res)

	c.SetCookie("access_token", access_token, 3600, "/", "http://127.0.0.1:3000", false, true)
	c.SetCookie("refresh_token", refresh_token, 3600, "/", "http://127.0.0.1:3000", false, true)
	c.SetCookie("access_token_expired_time", strconv.Itoa(expired_in), 3600, "/", "http://127.0.0.1:3000", false, true)
	c.SetCookie("refresh_token_expired_time", strconv.Itoa(refresh_expires_in), 3600, "/", "http://127.0.0.1:3000", false, true)
	return access_token
}

func request_AccessToken(grantType string, c *gin.Context) string {

	data := url.Values{}
	data.Set("client_id", "admin-cli")
	data.Set("client_secret", "7eceed48-073d-4c47-bb30-aae22ac14366")
	data.Set("grant_type", grantType)
	data.Set("username", "admin")
	data.Set("password", "admin")

	refreshToken, _ := c.Cookie("refresh_token")
	if grantType == "refresh_token" {
		data.Set("refresh_token", refreshToken)
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
