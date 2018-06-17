package main

import (
	"log"
	"net/http"
	"github.com/anon767/swaggertest/api"
	"github.com/anon767/swaggertest/database"
	_ "github.com/jinzhu/gorm"
)

func main() {
	log.Printf("Server started")

	router := api.NewRouter(database.Bootstrap())

	log.Fatal(http.ListenAndServe(":8080", router))
}
