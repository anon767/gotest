/*
 * Simple API
 *
 * A simple API to learn how to write OpenAPI Specification
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type Person struct {

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Username string `json:"username"`
}
