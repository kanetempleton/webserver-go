package main

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	username := "john_doe"
	email := "john@example.com"
	password := "secretpassword"

	u := NewUser(username, email, password)

	if u.Username != username {
		t.Errorf("Expected username %s, but got %s", username, u.Username)
	}

	if u.Email != email {
		t.Errorf("Expected email %s, but got %s", email, u.Email)
	}

	if u.Password != password {
		t.Errorf("Expected password %s, but got %s", password, u.Password)
	}
}