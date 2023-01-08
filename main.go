package main

import (
	"log"
	"net/http"

	"github.com/bharath79/golang/router"
)

func main() {
	r := router.Router()

	log.Fatal(http.ListenAndServe(":4000",r))
}