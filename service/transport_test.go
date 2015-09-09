package service

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestDecodeCountRequest(t *testing.T) {
	var r http.Request

	r.Body = ioutil.NopCloser(strings.NewReader("{\"s\" : \"Hello, World?\"}"))
	req, err := decodeCountRequest(&r)
	if err != nil {
		t.Error("Failed to decode valid count request")
	}
	if req.(countRequest).S != "Hello, World?" {
		t.Error("Failed to transcribe count request correctly")
	}
}

func TestDecodeInvalidCountRequest(t *testing.T) {
	var r http.Request

	r.Body = ioutil.NopCloser(strings.NewReader("{\"s\" : 4}"))
	req, err := decodeCountRequest(&r)
	// log.Printf("%s\n", req)
	if err == nil {
		t.Error("Failed to report error when decoding invalid count request")
	}

	if req != nil {
		t.Error("Returned request for invalid JSON encoding")
	}
}

func TestDecodeUppercaseRequest(t *testing.T) {
	var r http.Request

	r.Body = ioutil.NopCloser(strings.NewReader("{\"s\" : \"Hello, World?\"}"))
	req, err := decodeUppercaseRequest(&r)
	if err != nil {
		t.Error("Failed to decode valid uppercase request")
	}
	if req.(uppercaseRequest).S != "Hello, World?" {
		t.Error("Failed to transcribe uppercase request correctly")
	}
}

func TestDecodeInvalidUppercaseRequest(t *testing.T) {
	var r http.Request

	r.Body = ioutil.NopCloser(strings.NewReader("{\"s\" : 4}"))
	req, err := decodeUppercaseRequest(&r)
	// log.Printf("%s\n", req)
	if err == nil {
		t.Error("Failed to report error when decoding invalid uppercase request")
	}

	if req != nil {
		t.Error("Returned request for invalid JSON encoding")
	}
}

type mockResponseWriter struct {
	buffer bytes.Buffer
}

func (rw *mockResponseWriter) Header() http.Header {
	var h http.Header
	return h
}

func (rw *mockResponseWriter) Write(b []byte) (int, error) {
	return rw.buffer.Write(b)
}

func (rw *mockResponseWriter) WriteHeader(c int) {

}

func TestEncodeCountRequest(t *testing.T) {
	w := new(mockResponseWriter)
	r := countResponse{5, ""}

	e := encodeResponse(w, r)

	if e != nil {
		t.Error("Failed to encode valid count response")
	}

	if "{\"v\":5}\n" != w.buffer.String() {
		t.Error("Failed to encode valid count response correctly")
	}
}

func TestEncodeErrorCountRequest(t *testing.T) {
	w := new(mockResponseWriter)
	r := countResponse{-1, errEmpty.Error()}
	e := encodeResponse(w, r)
	if e != nil {
		t.Error("Failed to encode valid error count response")
	}

	if "{\"v\":-1,\"err\":\"Empty String\"}\n" != w.buffer.String() {
		t.Error("Failed to encode valid error count response")
	}
}

func TestEncodeUppercaseRequest(t *testing.T) {
	w := new(mockResponseWriter)
	r := uppercaseResponse{"HELLO, WORLD?", ""}

	e := encodeResponse(w, r)

	if e != nil {
		t.Error("Failed to encode valid count response")
	}

	if "{\"v\":\"HELLO, WORLD?\"}\n" != w.buffer.String() {
		t.Error("Failed to encode valid count response correctly")
	}
}

func TestEncodeErrorUpperRequest(t *testing.T) {
	w := new(mockResponseWriter)
	r := uppercaseResponse{"", errEmpty.Error()}
	e := encodeResponse(w, r)
	if e != nil {
		t.Error("Failed to encode valid error count response")
	}

	if "{\"err\":\"Empty String\"}\n" != w.buffer.String() {
		t.Error("Failed to encode valid error count response")
	}
}
