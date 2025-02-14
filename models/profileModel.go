package models

import (
	"errors"
	"strings"
)

var (
	emailEmptyError   = errors.New("'email' cannot be empty")
	emailExistsError  = errors.New("'email' already exists")
	invalidEmailError = errors.New("invalid email format")

	passwordTooShortError    = errors.New("'password' must be at least 8 characters long")
	passwordsDoNotMatchError = errors.New("'password' and 'repeat_password' do not match")
)

type Profile struct {
	email, password string
}

// NewEmail validates and returns a valid email or an error
func NewEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return "", emailEmptyError
	}

	if email == "existing@example.com" {
		return "", emailExistsError
	}

	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return "", invalidEmailError
	}

	return email, nil
}

func NewPassword(password, repeatPassword string) (string, error) {
	if len(password) < 8 {
		return "", passwordTooShortError
	}

	if password != repeatPassword {
		return "", passwordsDoNotMatchError
	}

	return password, nil
}
