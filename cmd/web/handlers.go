package main

import (
	"net/http"
	"strconv"
)

func (a *App) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	snippets, err := a.Database.LatestSnippets()
	if err != nil {
		a.ServerError(w, err)
		return
	}

	a.RenderHTML(w, "home.page.html", &HTMLData{Snippets: snippets})
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
	a.RenderHTML(w, "show.page.html", &HTMLData{Snippet: snippet})
}

func (a *App) newSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet here"))
}
