package credentials

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

var OAuthConf *oauth2.Config

type UserInfo struct {
	Username  string `binding:"required" json:"username"`
	FirstName string `binding:"required" json:"firstName"`
	LastName  string `binding:"required" json:"lastName"`
	Enabled   string `binding:"required" json:"enabled"`
	Email     string `binding:"required" json:"email"`
}

type Authdetails struct {
	//ClientInfo   authInterface.ClientInfoInterface
	APIClient    string
	APISecret    string
	UserName     string
	password     string
	Account      string
	APIURL       string
	OrbitURL     string
	currentToken oauth2.Token
}

func (auth *Authdetails) ClientInit() *Authdetails {

	return &Authdetails{
		APIClient: "admin-cli",
		APISecret: "7eceed48-073d-4c47-bb30-aae22ac14366",
		UserName:  "admin",
		password:  "admin",
		APIURL:    "http://192.168.0.118:9090/auth/realms/master/protocol/openid-connect/token",
	}
}

func (auth *Authdetails) RequestApi(ctx context.Context) error {

	httpClient := &http.Client{Timeout: 10 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	client, err := auth.GetClient(ctx)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Fetching API Client")
	log.Printf("[DEBUG] Client : ", client)

	user := UserInfo{}
	fmt.Println("addUser :", user)
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
	if err == nil {
		str := string(respBody)
		fmt.Println("str : " + str)
	}
	return nil
}

func (auth *Authdetails) GetClient(ctx context.Context) (*http.Client, error) {

	OAuthConf = &oauth2.Config{
		ClientID:     auth.APIClient,
		ClientSecret: auth.APISecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: auth.APIURL,
		},
	}

	log.Println("[DEBUG] Obtaining Tokensource for user %s", auth.UserName)

	token, err := OAuthConf.PasswordCredentialsToken(ctx, auth.UserName, auth.password)

	if err != nil {
		return nil, err
	}

	client := OAuthConf.Client(ctx, token)
	return client, nil
}
