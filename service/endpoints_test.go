package service

import (
	"log"
	"testing"
)

func TestCountEndport(t *testing.T) {
	svc := stringService{}
	ep := makeCountEndpoint(svc)
	req := countRequest{"Hello, World?"}
	resp, e := ep(nil, req)
	log.Printf("%d\n", resp.(countResponse).V)

	if e != nil {
		t.Error("Unexpected error when calling through end point")
	}

	if resp.(countResponse).V != 13 {
		t.Error("Unexpected result calling through count endpoint")
	}

	if resp.(countResponse).Err != "" {
		t.Error("Unexpected error when calling through count endpoint")
	}
}

func TestEmptyCountEndport(t *testing.T) {
	svc := stringService{}
	ep := makeCountEndpoint(svc)
	req := countRequest{""}
	resp, e := ep(nil, req)
	log.Printf("%d\n", resp.(countResponse).V)

	if e != nil {
		t.Error("Unexpected error when calling through end point")
	}

	if resp.(countResponse).V != -1 {
		t.Error("Unexpected result calling through count endpoint with emtpy string")
	}

	if resp.(countResponse).Err != errEmpty.Error() {
		t.Error("Expected error when calling through count endpoint with empty string")
	}
}

func TestUppercaseEndport(t *testing.T) {
	svc := stringService{}
	ep := makeUppercaseEndpoint(svc)
	req := uppercaseRequest{"Hello, World?"}
	resp, e := ep(nil, req)
	log.Printf("%d\n", resp.(uppercaseResponse).V)

	if e != nil {
		t.Error("Unexpected error when calling through end point")
	}

	if resp.(uppercaseResponse).V != "HELLO, WORLD?" {
		t.Error("Unexpected result calling through uppercase endpoint")
	}

	if resp.(uppercaseResponse).Err != "" {
		t.Error("Unexpected error when calling through uppercase endpoint")
	}
}

func TestEmptyUppercaseEndport(t *testing.T) {
	svc := stringService{}
	ep := makeUppercaseEndpoint(svc)
	req := uppercaseRequest{""}
	resp, e := ep(nil, req)
	log.Printf("%d\n", resp.(uppercaseResponse).V)

	if e != nil {
		t.Error("Unexpected error when calling through end point")
	}

	if resp.(uppercaseResponse).V != "" {
		t.Error("Unexpected result calling through uppercase endpoint with emtpy string")
	}

	if resp.(uppercaseResponse).Err != errEmpty.Error() {
		t.Error("Expected error when calling through uppercase endpoint with empty string")
	}
}
