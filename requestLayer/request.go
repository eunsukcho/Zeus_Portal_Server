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

	"golang.org/x/oauth2"
)

var OAuthConf *oauth2.Config

type AuthInfo struct {
	*models.Authdetails
}

func NewAuthInfo() (*AuthInfo, error) {

	return &AuthInfo{
		&models.Authdetails{
			APIClient: "admin-cli",
			APISecret: "7eceed48-073d-4c47-bb30-aae22ac14366",
			UserName:  "admin",
			Password:  "admin",
			APIURL:    "http://192.168.0.118:9090/auth/realms/master/protocol/openid-connect/token",
		}}, nil
}

func DefaultClient() {
	//client := oauth2.Transport
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
		ClientID:     auth.APIClient,
		ClientSecret: auth.APISecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: auth.APIURL,
		},
	}

	log.Println("[DEBUG] Obtaining Tokensource for user %s", auth.UserName)

	token, err := OAuthConf.PasswordCredentialsToken(ctx, auth.UserName, auth.Password)
	if err != nil {
		panic(err)
	}
	auth.CurrentToken = OAuthConf.TokenSource(ctx, token)
	return token
}

func (auth *AuthInfo) RequestUserListApi(ctx context.Context, client *http.Client) ([]models.ResponseUserInfo, error) {
	log.Printf("[DEBUG] Fetching API Client - UserListApi")
	resp, err := client.Get(
		"http://192.168.0.118:9090/auth/admin/realms/parthenon/users",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		userinfo := &[]models.ResponseUserInfo{}
		_ = json.Unmarshal([]byte(string(respBody)), userinfo)

		return *userinfo, nil
	}
	return nil, nil

}

func (auth *AuthInfo) RequestRegisterUserApi(ctx context.Context, user models.RegisterUserInfo, client *http.Client) (string, error) {

	log.Printf("[DEBUG] Fetching API Client - RegisterUserApi")

	fmt.Println(user)
	ubytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(ubytes)

	resp, err := client.Post(
		"http://192.168.0.118:9090/auth/admin/realms/parthenon/users",
		"application/json",
		buff,
	)
	if err != nil {
		log.Fatal(err)
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
