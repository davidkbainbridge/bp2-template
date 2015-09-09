package service

import (
	"testing"
)

func TestCountEmpty(t *testing.T) {
	var svc stringService
	v, e := svc.Count("")
	if v != -1 || e != errEmpty {
		t.Error("Failed when passed an empty string.")
	}
}

func TestCountGood(t *testing.T) {
	var svc stringService
	v, e := svc.Count("Hello, World?")
	if v != 13 || e != nil {
		t.Error("Failed to return proper count.")
	}
}

func TestUppercaseEmpty(t *testing.T) {
	var svc stringService
	v, e := svc.Uppercase("")
	if v != "" || e != errEmpty {
		t.Error("Failed when passed an empty string.")
	}
}

func TestUppercaseGood(t *testing.T) {
	var svc stringService
	v, e := svc.Uppercase("Hello, World?")
	if v != "HELLO, WORLD?" || e != nil {
		t.Error("Failed to return proper uppercase string.")
	}
}
