package Endpoint

import (
	"code/Service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid int`json:"uid"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func GetUserEndpoint(UserService Service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := UserService.GetName(r.Uid)
		return UserResponse{Result:result}, nil
	}
}

