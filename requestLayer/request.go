package requestLayer

import (
	"bytes"
	"context"
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
		}}, nil
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

func GetClient(ctx context.Context, auth *AuthInfo) (*http.Client, error) {

	httpClient := &http.Client{Timeout: 10 * time.Second}

	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	token := auth.GetApiClientTokenSource(ctx)

	client := OAuthConf.Client(ctx, token)
	return client, nil
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
		panic(err)
	}

	return token
}

func (auth *AuthInfo) RequestUserListApi(ctx context.Context, client *http.Client) ([]models.ResponseUserInfo, error) {
	log.Printf("[DEBUG] Fetching API Client - Request UserList Api")
	resp, err := client.Get(
		"https://docker.jointree.co.kr:8443/auth/admin/realms/parthenon/users/",
	)

	if err != nil || resp.StatusCode != 200 {
		log.Println("Connection Error")
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

func (auth *AuthInfo) RequestOneUserApi(ctx context.Context, user string, client *http.Client) (models.ResponseUserInfo, error) {
	log.Printf("[DEBUG] Fetching API Client - Request One User Api")

	resp, err := client.Get(
		"https://docker.jointree.co.kr:8443/auth/admin/realms/parthenon/users/" + user,
	)

	if err != nil || resp.StatusCode != 200 {
		log.Println("Connection Error")
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

func (auth *AuthInfo) RequestRegisterUserApi(ctx context.Context, user models.RegisterUserInfo, client *http.Client) (string, error) {

	log.Printf("[DEBUG] Fetching API Client - Request Register User Api")

	fmt.Println(user)
	ubytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(ubytes)

	resp, err := client.Post(
		"https://docker.jointree.co.kr:8443/auth/admin/realms/parthenon/users",
		"application/json",
		buff,
	)
	fmt.Println(resp.StatusCode)
	if err != nil || resp.StatusCode != 201 {
		fmt.Println("register error")
		fmt.Println(err)

	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}

func (auth *AuthInfo) DeleteUserApi(ctx context.Context, user string, client *http.Client) (string, error) {
	log.Printf("[DEBUG] Fetching API Client - Delete User Api")

	req, err := http.NewRequest(
		"DELETE",
		"https://docker.jointree.co.kr:8443/auth/admin/realms/parthenon/users/"+user,
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

		return "error", err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	return string(respBody), nil
}

func (auth *AuthInfo) UpdateUserApi(ctx context.Context, user models.RegisterUserInfo, client *http.Client) (string, error) {

	log.Printf("[DEBUG] Fetching API Client - Request Update User Api")

	fmt.Println(user)

	ubytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(ubytes)

	req, err := http.NewRequest(
		"PUT",
		"https://docker.jointree.co.kr:8443/auth/admin/realms/parthenon/users/"+user.ID,
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
		return "", err
	}
	return string(respBody), nil
}

func (auth *AuthInfo) UpdateUserCredentialsApi(ctx context.Context, user string, client *http.Client) (string, error) {

	log.Printf("[DEBUG] Fetching API Client - Request Update User Api")

	ubytes, _ := json.Marshal([]string{"UPDATE_PASSWORD"})
	buff := bytes.NewBuffer(ubytes)

	req, err := http.NewRequest(
		"PUT",
		"https://docker.jointree.co.kr:8443/auth/admin/realms/parthenon/users/"+user+"/execute-actions-email?redirect_uri?http://192.168.0.102:4201",
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

		return "error", err
	}
	respBody, err := ioutil.ReadAll(resp.Body)

	return string(respBody), nil
}

func (auth *AuthInfo) RequestGroupListApi(ctx context.Context, client *http.Client) ([]models.ResGroupInfo, error) {
	log.Printf("[DEBUG] Fetching API Client - User Groups Api")
	resp, err := client.Get(
		"https://docker.jointree.co.kr:8443/auth/admin/realms/parthenon/groups?briefRepresentation=false",
		//"http://192.168.0.118:9090/auth/admin/realms/parthenon/groups?briefRepresentation=false",
	)
	if err != nil || resp.StatusCode != 200 {
		log.Println("Connection Error")
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil && resp.StatusCode == 200 {
		groups := &[]models.ResGroupInfo{}
		_ = json.Unmarshal([]byte(string(respBody)), groups)

		return *groups, nil
	}
	return nil, err
}

func (auth *AuthInfo) RequestRegisterGroupsApi(ctx context.Context, group models.ReqToken, client *http.Client) (string, error) {

	log.Printf("[DEBUG] Fetching API Client - Register Groups Api")

	fmt.Println("GROUP : ", group)
	ubytes, _ := json.Marshal(group)
	buff := bytes.NewBuffer(ubytes)

	req, err := http.NewRequest(
		"PUT",
		"https://docker.jointree.co.kr:8443/auth/admin/realms/parthenon/groups/"+group.Id,
		//"http://192.168.0.118:9090/auth/admin/realms/parthenon/groups/"+group.Id,
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

		return "error", err
	}
	defer resp.Body.Close()

	return "", nil
}
