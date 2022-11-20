package users

import "net/http"

type Controller interface {
	RegisterUser() http.HandlerFunc
	LoginUser() http.HandlerFunc
	Protected() http.HandlerFunc
	RequestNewAccess() http.HandlerFunc
}
