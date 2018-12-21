package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nbriar/authenz-go/user"
)

// CreateUser handles the creation of a new user with credentials
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var newUser user.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		panic(err)
	}

	// TODO create the user and the strategy here

	w.Write([]byte("a user was created"))
}
