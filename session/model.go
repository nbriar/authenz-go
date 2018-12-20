package session

// Session manages when and how the user is authenticated
type Session struct {
	CreatedAt string
	ExpiresAt string
	Token string
}