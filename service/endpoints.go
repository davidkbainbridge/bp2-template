package service

import (
	"github.com/davidkbainbridge/bp2-template/hooks"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		ip, err := hooks.GetMyIP()
		if err != nil {
			ip = "0.0.0.0"
		}
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, ip, err.Error()}, nil
		}
		return uppercaseResponse{v, ip, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		ip, err := hooks.GetMyIP()
		if err != nil {
			ip = "0.0.0.0"
		}
		v, err := svc.Count(req.S)
		if err != nil {
			return countResponse{v, ip, err.Error()}, nil
		}
		return countResponse{v, ip, ""}, nil
	}
}
