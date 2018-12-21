package credentials

import "os/user"

// Strategy is the authentication interface
type Strategy interface {
	GetName() string
	Register() bool
	Remove() bool
	Validate() bool
}

// Strategies available to implement for a user's authentication
var Strategies []Strategy

func init() {
	Strategies = make([]Strategy, 0)
	Strategies = append(Strategies, &PasswordStrategy{})
}

// Create the authentication strategy for the user
func Create(user *user.User) {
	// TODO create the proper strategy
}
