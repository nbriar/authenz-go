package credentials

import "github.com/jinzhu/gorm"

// PasswordStrategy is an authentication strategy using a username and password
type PasswordStrategy struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetName for PasswordStrategy is used to determine which type of strategy is being used for authentication
func (strategy *PasswordStrategy) GetName() string {
	return "password"
}

// Register is used to register this strategy with a specific user
func (strategy *PasswordStrategy) Register() bool {
	// do some registration here
	return false
}

// Remove is used to delete this strategy for a specific user
func (strategy *PasswordStrategy) Remove() bool {
	// do some deletion here
	return false
}

// Validate is used to check if the credentials for this strategy are correct
func (strategy *PasswordStrategy) Validate() bool {
	// do some validation here
	return false
}
