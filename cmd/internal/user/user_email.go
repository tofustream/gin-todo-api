package user

import (
	"errors"
	"regexp"
)

type UserEmail struct {
	value string
}

func NewUserEmail(value string) (UserEmail, error) {
	if value == "" {
		return UserEmail{}, errors.New("email cannot be empty")
	}

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(value) {
		return UserEmail{}, errors.New("invalid email format")
	}

	return UserEmail{value: value}, nil
}

func (e UserEmail) Value() string {
	return e.value
}

func (e UserEmail) Equals(other UserEmail) bool {
	return e.value == other.value
}
