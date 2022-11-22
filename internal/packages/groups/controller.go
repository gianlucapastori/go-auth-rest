package groups

import "net/http"

type Controller interface {
	CreateGroup() http.HandlerFunc
	RemoveGroup() http.HandlerFunc
	UpdateGroup() http.HandlerFunc
	CreateTask() http.HandlerFunc
	RemoveTask() http.HandlerFunc
}
