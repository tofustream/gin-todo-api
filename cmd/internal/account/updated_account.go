package account

import (
	"time"
)

type UpdatedAccount struct {
	accountID      AccountID
	email          AccountEmail
	hashedPasswrod HashedAccountPassword
	updatedAt      time.Time
	isDeleted      bool
}

func NewUpdatedAccount(
	accountID AccountID,
	email AccountEmail,
	hashedPassword HashedAccountPassword,
	updatedAt time.Time,
	isDeleted bool,
) UpdatedAccount {
	return UpdatedAccount{
		accountID:      accountID,
		email:          email,
		hashedPasswrod: hashedPassword,
		updatedAt:      updatedAt,
		isDeleted:      isDeleted,
	}
}

func (a UpdatedAccount) ID() AccountID {
	return a.accountID
}

func (a UpdatedAccount) Email() AccountEmail {
	return a.email
}

func (a UpdatedAccount) HashedPassword() HashedAccountPassword {
	return a.hashedPasswrod
}

func (a UpdatedAccount) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a UpdatedAccount) IsDeleted() bool {
	return a.isDeleted
}
