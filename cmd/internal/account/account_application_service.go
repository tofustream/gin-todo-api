package account

import (
	"time"

	"github.com/google/uuid"
	"github.com/tofustream/gin-todo-api/pkg/timestamp"
)

type IAccountApplicationService interface {
	// アカウントを取得
	FindAccount(accountID string) (*FetchedAccountDTO, error)

	// アカウントを登録
	Signup(email string, plainPassword string) error

	// アカウント情報を更新
	UpdateAccount(command IAccountCommand) error
}

type AccountApplicationService struct {
	repository IAccountRepository
}

func NewAccountApplicationService(repository IAccountRepository) IAccountApplicationService {
	return &AccountApplicationService{
		repository: repository,
	}
}

func (s AccountApplicationService) FindAccount(accountID string) (*FetchedAccountDTO, error) {
	accountIDInstance, err := NewAccountIDFromString(accountID)
	if err != nil {
		return nil, err
	}

	fetchedAccount, err := s.repository.FindAccount(accountIDInstance)
	if err != nil {
		return nil, err
	}

	dto := fetchedAccountToFetchedAccountDTO(*fetchedAccount)

	return &dto, nil
}

// アカウントエンティティを生成
func generateAccountEntity(email string, plainPassword string) (*Account, error) {
	now := time.Now()
	newTimestamp, err := timestamp.NewTimestamp(now, now)
	if err != nil {
		return nil, err
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	newAccountID, err := NewAccountIDFromUUID(newUUID)
	if err != nil {
		return nil, err
	}

	newAccountEmail, err := NewAccountEmail(email)
	if err != nil {
		return nil, err
	}

	newAccountPassword, err := NewAccountPassword(plainPassword)
	if err != nil {
		return nil, err
	}

	newAccount := NewAccount(newAccountID, newAccountEmail, newAccountPassword, newTimestamp)
	return &newAccount, nil
}

// アカウントを登録
func (s AccountApplicationService) Signup(email string, plainPassword string) error {
	newAccount, err := generateAccountEntity(email, plainPassword)
	if err != nil {
		return err
	}
	return s.repository.AddAccount(*newAccount)
}

func (s AccountApplicationService) UpdateAccount(command IAccountCommand) error {
	return command.Execute(s.repository)
}
