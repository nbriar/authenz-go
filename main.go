package main

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/nbriar/authenz-go/token"
)

func init() {
	// TODO set these from the environment
	db, err := gorm.Open("postgresql", "host=localhost dbname=authenz_dev")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
}

func main() {
	addr := ":3333"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	r := chi.NewRouter()
	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(token.Verifier())
		r.Use(token.Authenticator)

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["sub"])))
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome Anonymous"))
		})
		// Register a user with an auth strategy
		r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(r.Body)
			var data myData
			err := decoder.Decode(&data)
			if err != nil {
				panic(err)
			}

			fmt.Println()
			w.Write([]byte("a user was created"))
		})
	})

	return r
}
