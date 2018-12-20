package token

import (
	"testing"
	"github.com/nbriar/authenz/user"
	"github.com/stretchr/testify/assert"
)

var testUser user.User

func init() {
	testUser = user.User{ID: 432, FirstName: "Test", LastName: "User"}
}
func TestGenerate(t *testing.T) {
	newToken := Generate(&testUser)
	assert.Equal(t, testUser.FirstName, newToken.User.FirstName, "should return string version")
	assert.NotEmpty(t, newToken.JWT,  "should have a JWT in the token")
}

func TestClaims(t *testing.T) {
	newToken := Generate(&testUser)
	claims := Claims(newToken.JWT)
	assert.Equal(t, float64(testUser.ID), claims["sub"], "sub should be the user id")
}


