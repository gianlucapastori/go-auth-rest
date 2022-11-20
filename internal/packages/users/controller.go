package users

import "net/http"

type Controller interface {
	RegisterUser() http.HandlerFunc
	Protected() http.HandlerFunc
	RequestNewAccess() http.HandlerFunc
}
