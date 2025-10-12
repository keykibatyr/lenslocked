package controllers

import "net/http"

type Template interface {
	ExecuteTemp(w http.ResponseWriter)
}
