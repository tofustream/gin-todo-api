package account

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
	fetchedAccount, err := repository.FindAccount(c.accountID)
	if err != nil {
		return err
	}

	updatedAccount := fetchedAccount.UpdateEmail(c.email)

	return repository.UpdateAccount(*updatedAccount)
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
	fetchedAccount, err := repository.FindAccount(c.accountID)
	if err != nil {
		return err
	}

	hashedValue, err := c.password.HashedValue()
	if err != nil {
		return err
	}
	hashedPassword, err := NewHashedAccountPassword(string(hashedValue))
	if err != nil {
		return err
	}
	updatedAccount := fetchedAccount.UpdatePassword(hashedPassword)

	return repository.UpdateAccount(*updatedAccount)
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
	fetchedAccount, err := repository.FindAccount(c.accountID)
	if err != nil {
		return err
	}

	updatedAccount := fetchedAccount.MarkAsDeleted()

	return repository.UpdateAccount(*updatedAccount)
}
