package client

import "github.com/ikarpovich/go-bitrix/types"
import "github.com/ikarpovich/go-bitrix/types/landing"

func (c *Client) LandingRepoRegister(request *landing.RepoRegisterRequest) (*types.IntResponse, error) {
	resp, err := c.DoRaw("landing.repo.register", request, &types.IntResponse{})
	return resp.Result().(*types.IntResponse), err
}

func (c *Client) LandingRepoUnregister(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.repo.unregister", data, &types.Response{})
	return resp.Result().(*types.Response), err
}

func (c *Client) LandingRepoGetappinfo(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("landing.repo.getappinfo", data, &types.Response{})
	return resp.Result().(*types.Response), err
}

