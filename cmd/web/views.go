package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func (a *App) RenderHTML(w http.ResponseWriter, page string) {

	files := []string{
		filepath.Join(a.HTMLDir, "base.html"),
		filepath.Join(a.HTMLDir, page),
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		a.ServerError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		a.ServerError(w, err)
		return
	}
}
