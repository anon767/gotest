package main

import (
	"log"
	"github.com/anon767/swaggertest/restapi"
	"github.com/anon767/swaggertest/database"
	_ "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
)

type API interface{
	Start(db * gorm.DB)
}

func main() {

	log.Printf("Server started")
 	db := database.Bootstrap()
	var api API = restapi.Api{}
	api.Start(db)


	defer db.Close()
}
