package token

import (
	"fmt"

	"github.com/go-chi/jwtauth"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nbriar/authenz/configuration"
	"github.com/nbriar/authenz/user"

)

// TokenAuth is the model for authenticating the jwt
var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte(configuration.AuthTokenSecret()), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `sub:123` here:
	_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{"sub": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

// Token contains a JWT which can be encoded and decoded and is attached to a session
type Token struct {
	JWT string
	User *user.User
}

// GenerateToken is used to generate a JWT for a specific user
func GenerateToken(user *user.User) Token {
	_, jwt, _ := TokenAuth.Encode(jwt.MapClaims{"sub": user.ID})
	token := Token{
		JWT: jwt,
		User: user,
	}

	return token
}
