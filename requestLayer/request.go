package requestLayer

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "strings"
	"time"
	"zeus/models"

	"golang.org/x/oauth2"
)

var OAuthConf *oauth2.Config

type AuthInfo struct {
	*models.Authdetails
	*KeycloakApiInfo
	Token *oauth2.Token
}

var errConnFail = errors.New("Connection Failed")
var tokenErr = errors.New("Invalid Token")

func NewAuthInfo(auth models.Authdetails) (*AuthInfo, error) {

	return &AuthInfo{
		&models.Authdetails{
			ClientId:     auth.ClientId,
			ClientSecret: auth.ClientSecret,
			AdminId:      auth.AdminId,
			AdminPw:      auth.AdminPw,
			TokenUrl:     auth.TokenUrl,
		},
		SettingKeycloakInfo("ope"),
		nil,
	}, nil
}

func InputAuthInit(inputAuth models.Authdetails, auth *AuthInfo) (*AuthInfo, bool, error) {
	fmt.Println("inputAuth : ", inputAuth)

	switch {
	case inputAuth.ClientId != auth.ClientId:
		return nil, false, nil
	case inputAuth.ClientSecret != auth.ClientSecret:
		return nil, false, nil
	case inputAuth.AdminId != auth.AdminId:
		return nil, false, nil
	case inputAuth.AdminPw != auth.AdminPw:
		return nil, false, nil
	case inputAuth.TokenUrl != auth.TokenUrl:
		return nil, false, nil
	}
	return nil, true, nil
}

func GetClient(ctx context.Context, token *oauth2.Token) *http.Client {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Timeout: 10 * time.Second, Transport: tr}

	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	var tokeValid = token.Valid()

	fmt.Println("tokeValid : ", tokeValid)

	fmt.Println("GetClient Token : ", token)

	client := OAuthConf.Client(ctx, token)

	return client
}

func (auth *AuthInfo) GetApiClientTokenSource(ctx context.Context) *oauth2.Token {
	OAuthConf = &oauth2.Config{
		ClientID:     auth.ClientId,
		ClientSecret: auth.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: auth.TokenUrl,
		},
	}
	log.Println("[DEBUG] Obtaining Tokensource for user %s", auth.AdminId)

	token, err := OAuthConf.PasswordCredentialsToken(ctx, auth.AdminId, auth.AdminPw)

	if err != nil {
		fmt.Println("GetApiClientTokenSource Error : ", err.Error())
	} else {
		fmt.Println("GetApiClientTokenSource :", token)
	}

	return token
}

func (auth *AuthInfo) RequestUserListApi(ctx context.Context, token *oauth2.Token) ([]models.ResponseUserInfo, error) {
	log.Printf("[DEBUG] Fetching API Client - Request UserList Api")
	fmt.Println("userendpoint :", auth.UserEndpoint)

	client := GetClient(ctx, token)
	resp, err := client.Get(
		auth.UserEndpoint,
	)

	if err != nil || resp.StatusCode != 200 {
		log.Println("Client Connection Error")
		return nil, errConnFail
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil && resp.StatusCode == 200 {
		userinfo := &[]models.ResponseUserInfo{}
		_ = json.Unmarshal([]byte(string(respBody)), userinfo)

		return *userinfo, nil
	}

	return nil, errConnFail

}

func (auth *AuthInfo) RequestUserListByGroupApi(ctx context.Context, groupId string, token *oauth2.Token) ([]models.ResponseUserInfo, error) {
	log.Printf("[DEBUG] Fetching API Client - Request UserList Api")

	client := GetClient(ctx, token)
	resp, err := client.Get(
		auth.GroupEndpoint + groupId + "/members",
	)

	if err != nil || resp.StatusCode != 200 {
		log.Println("Client Connection Error")

		return nil, errConnFail
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil && resp.StatusCode == 200 {
		userinfo := &[]models.ResponseUserInfo{}
		_ = json.Unmarshal([]byte(string(respBody)), userinfo)

		return *userinfo, nil
	}

	return nil, errConnFail

}

func (auth *AuthInfo) RequestOneUserApi(ctx context.Context, user string, token *oauth2.Token) (models.ResponseUserInfo, error) {
	log.Printf("[DEBUG] Fetching API Client - Request One User Api")

	client := GetClient(ctx, token)
	resp, err := client.Get(
		auth.UserEndpoint + user,
	)

	if err != nil || resp.StatusCode != 200 {
		log.Println("Client Connection Error")
		return models.ResponseUserInfo{}, errConnFail
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil && resp.StatusCode == 200 {
		userinfo := &models.ResponseUserInfo{}
		_ = json.Unmarshal([]byte(string(respBody)), userinfo)

		return *userinfo, nil
	}

	return models.ResponseUserInfo{}, errConnFail
}

func (auth *AuthInfo) RequestRegisterUserApi(ctx context.Context, user models.RegisterUserInfo, token *oauth2.Token) (string, int, error) {

	log.Printf("[DEBUG] Fetching API Client - Request Register User Api")

	fmt.Println(user)
	ubytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(ubytes)

	client := GetClient(ctx, token)
	resp, err := client.Post(
		auth.UserEndpoint,
		"application/json",
		buff,
	)
	fmt.Println(resp.StatusCode)

	if err != nil {
		fmt.Println("register error")
		return "", 0, errConnFail
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 201 {
		fmt.Println(string(respBody))
	}

	fmt.Println(string(respBody))
	if err != nil {
		return "", 0, err
	}
	return string(respBody), resp.StatusCode, nil
}

func (auth *AuthInfo) DeleteUserApi(ctx context.Context, user string, token *oauth2.Token) (string, error) {
	log.Printf("[DEBUG] Fetching API Client - Delete User Api")

	client := GetClient(ctx, token)

	req, err := http.NewRequest(
		"DELETE",
		auth.UserEndpoint+user,
		nil,
	)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "error", err
	}

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 204 {
		t, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("str : " + string(t))

		return "error", errConnFail
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	return string(respBody), nil
}

func (auth *AuthInfo) UpdateUserApi(ctx context.Context, user models.RegisterUserInfo, token *oauth2.Token) (string, error) {

	log.Printf("[DEBUG] Fetching API Client - Request Update User Api")

	fmt.Println(user)

	ubytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(ubytes)

	client := GetClient(ctx, token)

	req, err := http.NewRequest(
		"PUT",
		auth.UserEndpoint+user.ID,
		buff,
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 204 {
		fmt.Println("register error")
		fmt.Println(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errConnFail
	}
	return string(respBody), nil
}

func (auth *AuthInfo) UpdateUserCredentialsApi(ctx context.Context, user string, token *oauth2.Token) (string, error) {
	client := GetClient(ctx, token)

	log.Printf("[DEBUG] Fetching API Client - Request Update User Api")

	ubytes, _ := json.Marshal([]string{"UPDATE_PASSWORD"})
	buff := bytes.NewBuffer(ubytes)

	req, err := http.NewRequest(
		"PUT",
		auth.UserEndpoint+user+"/execute-actions-email",
		buff,
	)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "error", err
	}

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 204 {
		t, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("str : " + string(t))

		return "error", errConnFail
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	return string(respBody), nil
}

func (auth *AuthInfo) RequestGroupListApi(ctx context.Context, group string, token *oauth2.Token) ([]models.ResGroupInfo, error) {
	client := GetClient(ctx, token)

	log.Printf("[DEBUG] Fetching API Client - User Groups Api (List --)")
	var requesturl string
	if group == "all" {
		requesturl = auth.GroupEndpoint + "?briefRepresentation=false"
	} else {
		requesturl = auth.GroupEndpoint + "?search=" + group
	}
	fmt.Println("requesturl : ", requesturl)
	resp, err := client.Get(
		requesturl,
	)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("RequestGroupListApi err : ", err)
		log.Println("Client Connection Error")
		return nil, errConnFail
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil && resp.StatusCode == 200 {
		groups := &[]models.ResGroupInfo{}
		_ = json.Unmarshal([]byte(string(respBody)), groups)

		return *groups, nil
	}
	defer resp.Body.Close()

	return nil, err
}

func (auth *AuthInfo) RequestRegisterGroupsApi(ctx context.Context, group models.ReqToken, token *oauth2.Token) (string, error) {
	client := GetClient(ctx, token)
	log.Printf("[DEBUG] Fetching API Client - Register Groups Api")

	fmt.Println("GROUP : ", group)
	fmt.Println("GROUP UPDATE URL : ", auth.GroupEndpoint+group.Id)
	ubytes, _ := json.Marshal(group)
	buff := bytes.NewBuffer(ubytes)

	req, err := http.NewRequest(
		"PUT",
		auth.GroupEndpoint+group.Id,
		buff,
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 204 {
		t, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("str : " + string(t))

		return "error", errConnFail
	}
	defer resp.Body.Close()

	return "", nil
}
