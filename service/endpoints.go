package service

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v, err := svc.Count(req.S)
		if err != nil {
			return countResponse{v, err.Error()}, nil
		}
		return countResponse{v, ""}, nil
	}
}
