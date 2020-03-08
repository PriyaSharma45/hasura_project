package main

import (
	"hasura/router"
	"hasura/service"
	"log"
	"net/http"
)

func main() {

	sqlHandler := new(router.SQLHandler)
	var err error
	sqlHandler.PostgresClient, err = service.PostgresConnection()
	if err != nil {
		log.Fatal("couldn't connect to database ", err)
	}

	ginRouter := router.GetRouterEngine(sqlHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: ginRouter,
	}

	server.ListenAndServe()
}
