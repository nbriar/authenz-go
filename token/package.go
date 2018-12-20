package token

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/nbriar/authenz/configuration"
	"github.com/nbriar/authenz/user"
)

// Auth is the model for authenticating the jwt
var Auth *jwtauth.JWTAuth

func init() {
	Auth = jwtauth.New("HS256", []byte(configuration.AuthTokenSecret()), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `sub:123` here:
	_, tokenString, _ := Auth.Encode(jwt.MapClaims{"sub": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

// Token contains a JWT which can be encoded and decoded and is attached to a session
type Token struct {
	JWT  string
	User *user.User
}

// Generate is used to generate a JWT for a specific user
func Generate(user *user.User) Token {
	_, jwt, _ := Auth.Encode(jwt.MapClaims{"sub": user.ID})
	token := Token{
		JWT:  jwt,
		User: user,
	}

	return token
}

// Claims returns the claims for the jwt
func Claims(tokenString string) jwt.MapClaims {
	claims, _ := Auth.Decode(tokenString)
	return claims.Claims.(jwt.MapClaims)
}
