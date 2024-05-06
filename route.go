package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/justiruel/penggajianPaperbag/handler"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func RegisterWebRoute(tmpl *template.Template) {
	http.HandleFunc("/", handler.HandlerIndex(tmpl, store))
	http.HandleFunc("/index", handler.HandlerIndex(tmpl, store))
	http.HandleFunc("/login", handler.HandlerLogin(tmpl, store))
	http.HandleFunc("/logout", handler.HandlerLogout(store))
}
