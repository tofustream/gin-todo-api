package account

import (
	"time"

	"github.com/tofustream/gin-todo-api/pkg/timestamp"
)

type Account struct {
	id        AccountID
	email     AccountEmail
	password  AccountPassword
	timestamp timestamp.Timestamp
	isDeleted bool
}

func NewAccount(
	id AccountID,
	email AccountEmail,
	password AccountPassword,
	timestamp timestamp.Timestamp,
) Account {
	return Account{
		id:        id,
		email:     email,
		password:  password,
		timestamp: timestamp,
		isDeleted: false,
	}
}

func (a Account) ID() AccountID {
	return a.id
}

func (a Account) Email() AccountEmail {
	return a.email
}

func (a Account) Password() AccountPassword {
	return a.password
}

func (a Account) CreatedAt() time.Time {
	return a.timestamp.CreatedAt()
}

func (a Account) UpdatedAt() time.Time {
	return a.timestamp.UpdatedAt()
}

func (a Account) IsDeleted() bool {
	return a.isDeleted
}
