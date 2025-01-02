package account

import "github.com/tofustream/gin-todo-api/pkg/timestamp"

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

func (a Account) MarkAsDeleted() Account {
	return Account{
		id:        a.id,
		email:     a.email,
		password:  a.password,
		timestamp: a.timestamp,
		isDeleted: true,
	}
}
