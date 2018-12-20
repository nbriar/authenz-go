package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/nbriar/authenz/user"
)

func TestToken_JWT(t *testing.T) {
	user := user.User{FirstName: "Test", LastName: "User"}
	assert.Equal(t, "Test", user.FirstName, "should return string version")
}