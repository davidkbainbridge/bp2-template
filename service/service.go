package service

import (
	"errors"
	"strings"
)

// StringService - Defines the interface of a simple Go-Kit based RPC service
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) (int, error)
}

type stringService struct{}

var errEmpty = errors.New("Empty String")

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) (int, error) {
	if s == "" {
		return -1, errEmpty
	}
	return len(s), nil
}

// uppercaseRequest - Request message for Uppercase RPC
type uppercaseRequest struct {
	S string `json:"s"`
}

// uppercaseResponse - Response message for Uppercase RPC
type uppercaseResponse struct {
	V   string `json:"v,omitempty"`
	By  string `json:"by,omitempty"`
	Err string `json:"err,omitempty"`
}

// countRequest - Request message for Count RPC
type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V   int    `json:"v,omitempty"`
	By  string `json:"by,omitempty"`
	Err string `json:"err,omitempty"`
}
