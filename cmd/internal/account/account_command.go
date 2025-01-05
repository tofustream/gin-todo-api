package account

import (
	"time"
)

type IAccountCommand interface {
	Execute(repository IAccountRepository) error
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

func (c UpdateAccountEmailCommand) Execute(repository IAccountRepository) error {
	dto, err := repository.FindAccount(c.accountID)
	if err != nil {
		return err
	}

	accountIDInstance, err := NewAccountIDFromUUID(dto.ID)
	if err != nil {
		return err
	}

	newEmailInstance, err := NewAccountEmail(c.email.Value())
	if err != nil {
		return err
	}

	hashedPasswordInstance, err := NewHashedAccountPassword(dto.Password)
	if err != nil {
		return err
	}

	updatedAccount := NewUpdatedAccount(
		accountIDInstance,
		newEmailInstance,
		hashedPasswordInstance,
		time.Now(),
		dto.IsDeleted,
	)

	return repository.UpdateAccount(updatedAccount)
}

type UpdateAccountPasswordCommand struct {
	accountID AccountID
	password  AccountPassword
}

func NewUpdateAccountPasswordCommand(accountID string, password string) (IAccountCommand, error) {
	accountIDInstance, err := NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}

	passwordInstance, err := NewAccountPassword(password)
	if err != nil {
		return nil, err
	}

	return &UpdateAccountPasswordCommand{
		accountID: accountIDInstance,
		password:  passwordInstance,
	}, nil
}

func (c UpdateAccountPasswordCommand) Execute(repository IAccountRepository) error {
	dto, err := repository.FindAccount(c.accountID)
	if err != nil {
		return err
	}

	accountIDInstance, err := NewAccountIDFromUUID(dto.ID)
	if err != nil {
		return err
	}

	emailInstance, err := NewAccountEmail(dto.Email)
	if err != nil {
		return err
	}

	hashedPassword, err := c.password.HashedValue()
	if err != nil {
		return err
	}
	newHashedPasswordInstance, err := NewHashedAccountPassword(string(hashedPassword))
	if err != nil {
		return err
	}

	updatedAccount := NewUpdatedAccount(
		accountIDInstance,
		emailInstance,
		newHashedPasswordInstance,
		time.Now(),
		dto.IsDeleted,
	)

	return repository.UpdateAccount(updatedAccount)
}

type MarkAsDeletedCommand struct {
	accountID AccountID
}

func NewMarkAsDeletedCommand(accountID string) (IAccountCommand, error) {
	accountIDInstance, err := NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}

	return &MarkAsDeletedCommand{
		accountID: accountIDInstance,
	}, nil
}

func (c MarkAsDeletedCommand) Execute(repository IAccountRepository) error {
	dto, err := repository.FindAccount(c.accountID)
	if err != nil {
		return err
	}

	accountIDInstance, err := NewAccountIDFromUUID(dto.ID)
	if err != nil {
		return err
	}

	emailInstance, err := NewAccountEmail(dto.Email)
	if err != nil {
		return err
	}

	hashedPasswordInstance, err := NewHashedAccountPassword(dto.Password)
	if err != nil {
		return err
	}

	updatedAccount := NewUpdatedAccount(
		accountIDInstance,
		emailInstance,
		hashedPasswordInstance,
		time.Now(),
		true, // Mark as deleted
	)

	return repository.UpdateAccount(updatedAccount)
}
