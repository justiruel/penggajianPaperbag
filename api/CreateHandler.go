package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/justiruel/penggajianPaperbag/config"
	"github.com/justiruel/penggajianPaperbag/model/user"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(config.DbDriver, config.DbUser+":"+config.DbPass+"@/"+config.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Parse JSON data from the request body
	var usr user.User
	json.NewDecoder(r.Body).Decode(&usr)

	user.CreateUser(db, usr.Name, usr.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}
