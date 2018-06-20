package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "Http network address")
	htmlDir := flag.String("html-dir", "./ui/html/", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static files")
	flag.Parse()

	app := &App{
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

	log.Println("Litening on ", *addr)
	log.Fatal(http.ListenAndServe(*addr, app.Routes()))

}
