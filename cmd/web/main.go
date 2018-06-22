package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"snippetbox.org/pkg/models"
)

func main() {
	addr := flag.String("addr", ":4000", "Http network address")
	dsn := flag.String("dsn", "sb:password@/snippetbox?parseTime=true", "MySQL DSN")
	htmlDir := flag.String("html-dir", "./ui/html/", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static files")
	flag.Parse()

	db := connect(*dsn)

	defer db.Close()

	app := &App{
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
		Database:  &models.Database{DB: db},
	}

	log.Println("Litening on ", *addr)
	log.Fatal(http.ListenAndServe(*addr, app.Routes()))

}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
