package account

import (
	"errors"
	"regexp"
)

var (
	ErrEmptyEmail   = errors.New("email cannot be empty")
	ErrInvalidEmail = errors.New("invalid email format")
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type AccountEmail struct {
	value string
}

func NewAccountEmail(value string) (AccountEmail, error) {
	if value == "" {
		return AccountEmail{}, ErrEmptyEmail
	}
	if !emailRegex.MatchString(value) {
		return AccountEmail{}, ErrInvalidEmail
	}
	return AccountEmail{value: value}, nil
}

func (e AccountEmail) Value() string {
	return e.value
}

func (e AccountEmail) Equals(other AccountEmail) bool {
	return e.value == other.value
}
