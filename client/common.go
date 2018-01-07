package client

import (
	"github.com/ikarpovich/go-bitrix/types"
	"gopkg.in/resty.v1"
)

func (c *Client) Methods(request *types.MethodsRequest) (*resty.Response, error) {
	return c.Do("methods", request)
}
