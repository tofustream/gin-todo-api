package account

import (
	"time"

	"github.com/tofustream/gin-todo-api/pkg/timestamp"
)

type FetchedAccount struct {
	id        AccountID
	email     AccountEmail
	password  HashedAccountPassword
	timestamp timestamp.Timestamp
}

func NewFetchedAccount(
	id AccountID,
	email AccountEmail,
	password HashedAccountPassword,
	timestamp timestamp.Timestamp,
) FetchedAccount {
	return FetchedAccount{
		id:        id,
		email:     email,
		password:  password,
		timestamp: timestamp,
	}
}

func (a FetchedAccount) ID() AccountID {
	return a.id
}

func (a FetchedAccount) Email() AccountEmail {
	return a.email
}

func (a FetchedAccount) HashedPassword() HashedAccountPassword {
	return a.password
}

func (a FetchedAccount) CreatedAt() time.Time {
	return a.timestamp.CreatedAt()
}

func (a FetchedAccount) UpdatedAt() time.Time {
	return a.timestamp.UpdatedAt()
}

func (a FetchedAccount) UpdateEmail(newEmail AccountEmail) *UpdatedAccount {
	updatedAccount := NewUpdatedAccount(
		a.id,
		newEmail,
		a.password,
		time.Now(),
		false,
	)
	return &updatedAccount
}

func (a FetchedAccount) UpdatePassword(newHashedPassword HashedAccountPassword) *UpdatedAccount {
	updatedAccount := NewUpdatedAccount(
		a.id,
		a.email,
		newHashedPassword,
		time.Now(),
		false,
	)
	return &updatedAccount
}

func (a FetchedAccount) MarkAsDeleted() *UpdatedAccount {
	updatedAccount := NewUpdatedAccount(
		a.id,
		a.email,
		a.password,
		time.Now(),
		true,
	)
	return &updatedAccount
}
