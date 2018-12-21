package credentials

// Null is used in cases where invalid credentials are passed in
type Null struct{}

// GetName for Null is used to determine which type of strategy is being used for authentication
func (strategy *Null) GetName() string {
	return "null"
}

// Register is used to register this strategy with a specific user
func (strategy *Null) Register() bool {
	// do some registration here
	return false
}

// Remove is used to delete this strategy for a specific user
func (strategy *Null) Remove() bool {
	// do some deletion here
	return false
}

// Validate is used to check if the credentials for this strategy are correct
func (strategy *Null) Validate() bool {
	// do some validation here
	return false
}
