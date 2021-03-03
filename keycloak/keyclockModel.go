package keycloak

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type KeyCloak interface {
	Get_Header_AccessTokenInfo()
	Request_AccessToken()
}

type request_Access_Token struct {
	client_id        string `binding:"required" json:clientId`
	client_secret    string `binding:"required" json:clientSecret`
	access_token_url string `binding:"required" json:accessTokenUrl`
	grant_type       string `binding:"required" json:grantType`
	username         string `binding:"required" json:username`
	password         string `binding:"required" json:password`
}

type response_Access_Token struct {
	Access_token       string `json:"access_token, string"`
	Expires_in         int    `json:"expires_in, int"`
	Refresh_expires_in int    `json:"refresh_expires_in, int"`
	Refresh_token      string `json:"refresh_token, string"`
	Token_type         string `json:"token_type, string"`
	Not_before_policy  int    `json:"not-before-policy, int"`
	Session_state      string `json:"s"ession_state, string"`
	Scope              string `json:"scope, string"`
}

func InitAccessTokenInfo() *request_Access_Token {
	rat := request_Access_Token{
		"admin-cli",
		"7eceed48-073d-4c47-bb30-aae22ac14366",
		"http://192.168.0.118:9090/auth/realms/master/protocol/openid-connect/token",
		"password",
		"admin",
		"admin",
	}
	return &rat
}

func InitResponseAccessTokenInfo() *response_Access_Token {
	rat := response_Access_Token{}
	return &rat
}

func (rat request_Access_Token) Get_Header_AccessTokenInfo() request_Access_Token {
	return rat
}

func (rat request_Access_Token) request_AccessToken(grantType string) string {
	fmt.Println("AccessToken Request Info : ", rat)

	data := url.Values{}
	data.Set("client_id", rat.client_id)
	data.Set("client_secret", rat.client_secret)
	data.Set("grant_type", grantType)
	data.Set("username", rat.username)
	data.Set("password", rat.password)

	req, err := http.NewRequest("POST", rat.access_token_url, strings.NewReader(data.Encode()))
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

func (response response_Access_Token) parsing_Response_Expires_in(response_str string) int {
	json.Unmarshal([]byte(response_str), &response)

	return response.Expires_in
}

func (response response_Access_Token) parsing_Response_RefreshExpires_in(response_str string) int {
	json.Unmarshal([]byte(response_str), &response)

	return response.Refresh_expires_in
}

func (response response_Access_Token) parsing_Response_AccessToken(response_str string) string {
	json.Unmarshal([]byte(response_str), &response)

	return response.Access_token
}

func (response response_Access_Token) parsing_Response_RefreshToken(response_str string) string {
	json.Unmarshal([]byte(response_str), &response)

	return response.Refresh_token
}
