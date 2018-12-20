package user

// User is the primary model for authentication
type User struct {
	ID uint `gorm:"primary_key"`
	FirstName string
	LastName string
	Email string
}