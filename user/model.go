package user

import (
	"encoding/json"
	"fmt"

	"github.com/nbriar/authenz-go/credentials"
)

// User is the primary model for authentication
type User struct {
	ID          uint `gorm:"primary_key"`
	FirstName   string
	LastName    string
	Email       string
	Strategy    json.RawMessage
	Credentials credentials.Strategy
}

type strategyType struct {
	Type string `json:"type"`
}

// SetCredentials is called to convert the raw json into the proper struct for the credentials
func (u *User) SetCredentials() {
	// TODO should probaby save these here too
	strat := strategyType{}
	if err := json.Unmarshal(u.Strategy, &strat); err != nil {
		return
	}

	for _, s := range credentials.Strategies {
		if strat.Type == s.GetName() {
			creds := s
			if err := json.Unmarshal(u.Strategy, &creds); err != nil {
				fmt.Println("something went horribly wrong")
				return
			}
			u.Credentials = creds
		}
	}

	if u.Credentials == nil {
		u.Credentials = &credentials.Null{}
	}
}
