package main

import (
	"log"
	"net/http"

	"github.com/davidkbainbridge/bp2-template/service"
)

func main() {
	log.Println("Hello World")
	service.Register()
	log.Fatal(http.ListenAndServe(":8901", nil))
}
