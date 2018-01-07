package client

import (
	"gopkg.in/resty.v1"
	"net/url"
	"github.com/pkg/errors"
	"fmt"
	"os"
)

type Client struct {
	client *resty.Client
	oAuth *OAuthData
	Url *url.URL
}

type OAuthData struct {
	AuthToken string
	RefreshToken string
}

type WebhookAuthData struct {
	UserID int
	Secret string
}

func NewClientWithOAuth(intranetUrl, authToken, refreshToken string) (*Client, error) {
	u, err := url.Parse(intranetUrl)
	if err != nil {
		return nil, errors.Wrap(err, "Error parsing B24 URL")
	}

	return &Client{
		client: resty.DefaultClient,
		Url:    u,
		oAuth: &OAuthData{
			AuthToken:    authToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func NewEnvClientWithOauth() (*Client, error) {
	return NewClientWithOAuth(
		os.Getenv("BITRIX_URL"),
		os.Getenv("BITRIX_AUTH_TOKEN"),
		os.Getenv("BITRIX_REFRESH_TOKEN"))
}

func (c *Client) Do(method string, data interface{}) (*resty.Response, error) {
	params := map[string]string{
		"auth": c.oAuth.AuthToken,
	};

	resty.SetHostURL(c.Url.String())
	resty.SetHeader("Accept", "application/json")
	req := resty.R();
	resp, err := req.SetQueryParams(params).
		SetBody(data).
		Post(fmt.Sprintf("/rest/%s", method))

	return resp, err
}
