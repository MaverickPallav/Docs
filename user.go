// user.go
package main

type User struct {
	username string
}

// NewUser constructor
func NewUser(username string) *User {
	return &User{username: username}
}

// GetUsername returns the username of the user
func (u *User) GetUsername() string {
	return u.username
}
