package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/nbriar/authenz-go/user"
)
// CreateUser handles the creation of a new user with credentials
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println("A user was created")
	w.Write([]byte("a user was created"))
}