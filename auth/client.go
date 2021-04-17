package auth

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"

	bca "github.com/ianeinser/bca-api-go"
)

//Client is used to invoke BCA OAuth 2.0 API
type Client struct {
	Client       bca.APIImplementation
	ClientID     string
	ClientSecret string
}

//NewClient is used to initialize new auth.Client
func NewClient(config bca.Config) Client {
	return Client{
		Client:       bca.NewAPI(config),
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
	}
}

//GetToken is used to get OAuth 2.0 token
func (c *Client) GetToken(ctx context.Context) (*bca.AuthToken, error) {
	path := "/api/oauth/token"

	header := http.Header{}
	header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.ClientID+":"+c.ClientSecret)))

	data := url.Values{}
	data.Add("grant_type", "client_credentials")

	var ptr_authToken *bca.AuthToken
	if err := c.Client.CallRaw("POST", path, "application/x-www-form-urlencoded",
		header, strings.NewReader(data.Encode()), ptr_authToken); err != nil {
		return ptr_authToken, err
	}
	return ptr_authToken, nil
}
