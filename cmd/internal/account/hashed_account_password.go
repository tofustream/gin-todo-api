package account

import "errors"

var (
	ErrEmptyHashedAccountPassword = errors.New("hashed account password is empty")
)

// ハッシュ化されたアカウントパスワードを表現するための構造体
type HashedAccountPassword struct {
	value string
}

func NewHashedAccountPassword(value string) (HashedAccountPassword, error) {
	if len(value) == 0 {
		return HashedAccountPassword{}, ErrEmptyHashedAccountPassword
	}

	return HashedAccountPassword{value: value}, nil
}

func (h HashedAccountPassword) Value() string {
	return h.value
}
