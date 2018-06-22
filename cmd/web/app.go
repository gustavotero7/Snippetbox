package main

import "snippetbox.org/pkg/models"

// App _
type App struct {
	HTMLDir   string
	StaticDir string
	Database  *models.Database
}
