/*
 * Simple API
 *
 * A simple API to learn how to write OpenAPI Specification
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package restapi

import (
	"net/http"
	"github.com/anon767/swaggertest/domain"
	"encoding/json"
	"github.com/gorilla/mux"
)

func PersonsGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

func PersonsPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PersonsUsernameDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PersonsUsernameFriendsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PersonsUsernameGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var person domain.Person
	db.Find(&person, "username = ?", vars["username"])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&person)
}
