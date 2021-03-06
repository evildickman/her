package main

import _ "github.com/mattn/go-sqlite3"

import (
	"github.com/go-code/her"
	"github.com/go-code/her/samples/book/handler"
)

var (
	config = map[string]interface{}{
		"TemplatePath": "templates",
		"Address":      "0.0.0.0",
		"XSRFCookies":  true,
		"CookieSecret": "book_secert",
		"Port":         8080,
		"Debug":        true,
	}
)

func main() {
	app := her.NewApplication(config)

	app.Database.Connection("sqlite", "sqlite3", "./book.s3db")

	app.Route.Handle("/", handler.Book.HomeHandler)
	app.Route.Handle("/create", handler.Book.CreateHandler)
	app.Route.Handle("/del/{id:[0-9]+}", handler.Book.DeleteHandler)
	app.Route.Handle("/static/{path:.*}", her.StaticFileHandler("static"))

	app.Start()
}
