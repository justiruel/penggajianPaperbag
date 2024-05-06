package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func HandlerLogout(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get a session. If the session doesn't exist, a new one will be created.
		session, err := store.Get(r, "credential")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Clear the session by removing all values.
		for key := range session.Values {
			delete(session.Values, key)
		}

		// Save the session to apply changes.
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect or respond as needed.
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
