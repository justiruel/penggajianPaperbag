package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/justiruel/penggajianPaperbag/config"
	"github.com/justiruel/penggajianPaperbag/model/user"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(config.DbDriver, config.DbUser+":"+config.DbPass+"@/"+config.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userID, err := strconv.Atoi(idStr)

	// Call the GetUser function to fetch the user data from the database
	usr, err := user.GetUser(db, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr)
}
