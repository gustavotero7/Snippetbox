package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"snippetbox.org/pkg/models"
)

// HTMLData _
type HTMLData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

// RenderHTML _
func (a *App) RenderHTML(w http.ResponseWriter, page string, data *HTMLData) {

	files := []string{
		filepath.Join(a.HTMLDir, "base.html"),
		filepath.Join(a.HTMLDir, page),
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		a.ServerError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)

	if err != nil {
		a.ServerError(w, err)
		return
	}
}
