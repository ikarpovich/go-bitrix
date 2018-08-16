package types

type MethodsRequest struct {
	Full bool `json:"full" form:"full"`
	Scope string `json:"scope" form:"scope"`
}

type MethodsResponse struct {
	Response
	Result []string
}