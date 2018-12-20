package token

import (
	"net/http"

	"github.com/go-chi/jwtauth"
)

// Verifier is a wrapper around jwtauth.Verifier
func Verifier() func(http.Handler) http.Handler {
	verifier := jwtauth.Verifier(Auth)
	return verifier
}
