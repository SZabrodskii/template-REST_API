package main

import (
	"github.com/SZabrodskii/REST_API"
	"github.com/SZabrodskii/REST_API/pkg/handler"
	"log"
)

func main() {
	handler := new(handler.Handler)
	srv := &rest_api.Server{}
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatal("error occured when trying to run http.server: %s", err.Error())
	}
}
