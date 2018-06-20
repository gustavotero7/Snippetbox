package main

import (
	"net/http"
)

func (a *App) Routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet", a.showSnippet)
	mux.HandleFunc("/snippet/new", a.newSnippet)

	fileServer := http.FileServer(http.Dir(a.StaticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
