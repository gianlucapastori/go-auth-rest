package users

import "net/http"

type Controller interface {
	RegisterUser() http.HandlerFunc
}
