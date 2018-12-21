package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"

	"github.com/nbriar/authenz-go/handlers"
	"github.com/nbriar/authenz-go/token"
)

// var db *gorm.DB

// func init() {
// 	db = repo.DBRepo
// 	defer db.Close()
// }

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
			w.Write([]byte("Welcome To Authenz Go"))
		})
		// Register a user with an auth strategy
		r.Post("/user", handlers.CreateUser)
	})

	return r
}
