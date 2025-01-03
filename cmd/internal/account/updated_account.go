package account

import (
	"time"

	"github.com/tofustream/gin-todo-api/pkg/timestamp"
)

type UpdatedAccount struct {
	accountID      AccountID
	email          AccountEmail
	hashedPasswrod HashedAccountPassword
	timeStamp      timestamp.Timestamp
	isDeleted      bool
}

func NewUpdatedAccount(
	accountID AccountID,
	email AccountEmail,
	hashedPassword HashedAccountPassword,
	timeStamp timestamp.Timestamp,
	isDeleted bool,
) UpdatedAccount {
	return UpdatedAccount{
		accountID:      accountID,
		email:          email,
		hashedPasswrod: hashedPassword,
		timeStamp:      timeStamp,
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

func (a UpdatedAccount) Timestamp() timestamp.Timestamp {
	return a.timeStamp
}

func (a UpdatedAccount) UpdatedAt() time.Time {
	return a.timeStamp.UpdatedAt()
}

func (a UpdatedAccount) IsDeleted() bool {
	return a.isDeleted
}
