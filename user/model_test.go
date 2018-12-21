package user

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_SetCredentials(t *testing.T) {
	newUser := User{Strategy: json.RawMessage(`{
			"type": "password",
			"username": "test23",
			"password": "password23"
		}`)}

	newUser.SetCredentials()

	assert.Equal(t, "password", newUser.Credentials.GetName(), "should set the credentials properly on the user")

}
func TestUser_SetCredentialsFailCase(t *testing.T) {
	newUser := User{Strategy: json.RawMessage(`{
			"type": "bogus"
		}`)}

	newUser.SetCredentials()

	assert.Equal(t, "null", newUser.Credentials.GetName(), "should set the credentials properly on the user")

}
