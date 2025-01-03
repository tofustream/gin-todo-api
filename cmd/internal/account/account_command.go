package account

import (
	"log"
	"time"

	"github.com/tofustream/gin-todo-api/pkg/timestamp"
)

type IAccountCommand interface {
	Execute(repository IAccountRepository) (*AccountDTO, error)
}

type UpdateAccountEmailCommand struct {
	accountID AccountID
	email     AccountEmail
}

func NewUpdateAccountEmailCommand(accountID string, email string) (IAccountCommand, error) {
	accountIDInstance, err := NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}

	emailInstance, err := NewAccountEmail(email)
	if err != nil {
		return nil, err
	}

	return &UpdateAccountEmailCommand{
		accountID: accountIDInstance,
		email:     emailInstance,
	}, nil
}

func (c *UpdateAccountEmailCommand) Execute(repository IAccountRepository) (*AccountDTO, error) {
	dto, err := repository.FindAccount(c.accountID)
	if err != nil {
		log.Println("FindAccount error")
		return nil, err
	}
	updatedTimestampInstance, err := timestamp.NewTimestamp(dto.CreatedAt, time.Now())
	if err != nil {
		return nil, err
	}
	accountIDInstance, err := NewAccountIDFromUUID(dto.ID)
	if err != nil {
		return nil, err
	}
	emailInstance, err := NewAccountEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	hashedPasswordInstance, err := NewHashedAccountPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	updatedAccount := NewUpdatedAccount(
		accountIDInstance,
		emailInstance,
		hashedPasswordInstance,
		updatedTimestampInstance,
		dto.IsDeleted,
	)

	updatedDTO, err := repository.UpdateAccount(updatedAccount)
	if err != nil {
		log.Println("UpdateAccount error")
		return nil, err
	}

	return updatedDTO, nil
}
