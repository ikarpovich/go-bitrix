package client

import (
	"github.com/ikarpovich/go-bitrix/types"
)

func (c *Client) Methods(request *types.MethodsRequest) (*types.MethodsResponse, error) {
	resp, err := c.Do("methods", request, &types.MethodsResponse{})
	return resp.(*types.MethodsResponse), err
}
