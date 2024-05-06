package handler

import (
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

func HandlerLogin(tmpl *template.Template, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "credential")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if r.Method == "POST" {
			var email = r.FormValue("email")
			var password = r.Form.Get("password")

			if email == "admin@aj.com" && password == "admin" {

				// Set some session values.
				session.Values["email"] = email

				// Save the session.
				err = session.Save(r, w)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/index", http.StatusFound)
			}
		} else if r.Method == "GET" {

			email, ok := session.Values["email"].(string)
			if ok && email == "admin@aj.com" {
				http.Redirect(w, r, "/index", http.StatusFound)
			}

			type M map[string]interface{}
			var data = M{
				"name":  "Admin",
				"title": "Login",
			}
			err := tmpl.ExecuteTemplate(w, "login", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
