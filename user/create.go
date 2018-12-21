package user

import "github.com/jinzhu/gorm"

// Create a user in the database
func Create(db *gorm.DB, user *User) (uint, error) {
	err := db.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
