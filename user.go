package main

// User represents a user account
type User struct {
	Username string
	Email    string
	Password string
}

// NewUser creates a new User instance
func NewUser(username, email, password string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
	}
}