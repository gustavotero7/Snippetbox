package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (a *App) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	a.RenderHTML(w, "home.page.html")
}

func (a *App) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := a.Database.GetSnippet(id)
	if err != nil {
		a.ServerError(w, err)
	}

	if snippet == nil {
		a.NotFound(w)
	}

	fmt.Fprint(w, snippet)

	//w.Write([]byte("Show snippet here"))
}

func (a *App) newSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet here"))
}
