package database

import "strings"

// BeforeSave for User.
func (u *User) BeforeSave() {
	u.Email = strings.ToLower(u.Email)
}
