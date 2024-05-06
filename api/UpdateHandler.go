package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/justiruel/penggajianPaperbag/config"
	"github.com/justiruel/penggajianPaperbag/model/user"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(config.DbDriver, config.DbUser+":"+config.DbPass+"@/"+config.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userID, _ := strconv.Atoi(idStr)

	var usr user.User
	err = json.NewDecoder(r.Body).Decode(&usr)

	// Call the GetUser function to fetch the user data from the database
	user.UpdateUser(db, userID, usr.Name, usr.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User updated successfully")
}
