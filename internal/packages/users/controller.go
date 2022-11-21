package users

import "net/http"

type Controller interface {
	RegisterUser() http.HandlerFunc
	LoginUser() http.HandlerFunc
	Protected() http.HandlerFunc
	ChangePassword() http.HandlerFunc
	RequestNewPassword() http.HandlerFunc
	RequestNewAccess() http.HandlerFunc
}
