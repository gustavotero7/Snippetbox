package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

func (a *App) ServerError(w http.ResponseWriter, err error) {
	log.Printf("%s\n%s\n", err.Error(), debug.Stack())
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func (a *App) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (a *App) NotFound(w http.ResponseWriter) {
	a.ClientError(w, http.StatusNotFound)
}
