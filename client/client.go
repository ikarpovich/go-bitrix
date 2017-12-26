package client

import "gopkg.in/resty.v1"

type Client struct {
	client *resty.Client
}

func NewClient() *Client {
	return &Client{
		client: resty.DefaultClient,
	}
}