package handler

import (
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

func HandlerIndex(tmpl *template.Template, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "credential")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if user is authenticated
		email, ok := session.Values["email"].(string)
		if !ok || email != "admin@aj.com" {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		type M map[string]interface{}
		var data = M{
			"name":  "Admin",
			"title": "SISTEM PENGGAJIAN PAPERBAG",
		}
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
