package main

import (
	"log"
	"net/http"

	"github.com/davidkbainbridge/bp2-template/service"
	"github.com/davidkbainbridge/bp2-template/hooks"
)

func main() {
	log.Println("Hello World")
	updates := make(chan interface{})
	go hooks.HandleHooks(updates)
	service.Register()
	log.Fatal(http.ListenAndServe(":8901", nil))
}
