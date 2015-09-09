package service

// StringService - Defines the interface of a simple Go-Kit based RPC service
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) (int, error)
}

// UppercaseRequest - Request message for Uppercase RPC
type UppercaseRequest struct {
	S string `json:"s"`
}

// CountRequest - Request message for Count RPC
type CountRequest struct {
	S string `json:"s"`
}
