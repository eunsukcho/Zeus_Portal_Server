package requestLayer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"zeus/models"
	"errors"

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
		&models.Authdetails {
			ClientId: auth.ClientId,
			ClientSecret: auth.ClientSecret,
			AdminId:  auth.AdminId,
			AdminPw:  auth.AdminPw,
			TokenUrl: auth.TokenUrl,
	}}, nil
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
	log.Printf("[DEBUG] Fetching API Client - UserListApi")
	resp, err := client.Get(
		"https://docker.jointree.co.kr:8443/auth/admin/realms/parthenon/users",
	)
	if err != nil {
		log.Println("Connection Error")
		return nil, errConnFail
	}
	if resp.StatusCode != 200 {
		return nil, err
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

func (auth *AuthInfo) RequestRegisterUserApi(ctx context.Context, user models.RegisterUserInfo, client *http.Client) (string, error) {

	log.Printf("[DEBUG] Fetching API Client - RegisterUserApi")

	fmt.Println(user)
	//user.ClientRoles
	ubytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(ubytes)

	resp, err := client.Post(
		"https://docker.jointree.co.kr:8443/auth/admin/realms/parthenon/users",
		"application/json",
		buff,
	)
	if resp.StatusCode != 200 {
		return "", err
	}

	if err != nil {
		return "", errConnFail
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	str := string(respBody)
	fmt.Println("str : " + str)
	return str, nil
}
