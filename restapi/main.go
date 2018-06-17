package restapi

import (
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Api struct {}

func (Api) Start(db * gorm.DB) {
	router := NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", router))
}