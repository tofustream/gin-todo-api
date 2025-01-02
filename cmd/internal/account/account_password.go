package account

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const minPasswordLength int = 8

var (
	ErrEmptyAccountPassword   = errors.New("password cannot be empty")
	ErrInvalidAccountPassword = fmt.Errorf("password must be at least %d characters long", minPasswordLength)
)

type AccountPassword struct {
	plain string
}

func NewAccountPassword(plain string) (AccountPassword, error) {
	if plain == "" {
		return AccountPassword{}, ErrEmptyAccountPassword
	}
	if len(plain) < minPasswordLength {
		return AccountPassword{}, ErrInvalidAccountPassword
	}

	return AccountPassword{plain: plain}, nil
}

func (p AccountPassword) Value() string {
	return p.plain
}

func (p AccountPassword) HashedValue() ([]byte, error) {
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(p.plain), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedValue, nil
}
