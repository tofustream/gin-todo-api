package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const minPasswordLength uint = 8

type UserPassword struct {
	value []byte
}

func NewUserPassword(value string) (UserPassword, error) {
	if value == "" {
		return UserPassword{}, errors.New("password cannot be empty")
	}

	if len(value) < 8 {
		return UserPassword{}, fmt.Errorf("password must be at least %d characters long", minPasswordLength)
	}

	hashedValue, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return UserPassword{}, err
	}

	return UserPassword{value: hashedValue}, nil
}

func (p UserPassword) Value() []byte {
	return p.value
}
