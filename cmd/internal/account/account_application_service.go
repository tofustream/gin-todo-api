package account

import (
	"time"

	"github.com/google/uuid"
	"github.com/tofustream/gin-todo-api/pkg/timestamp"
)

type IAccountApplicationService interface {
	// アカウントを登録
	RegisterAccount(email string, plainPassword string) (*AccountRegistrationResponseDTO, error)
}

type AccountApplicationService struct {
	repository IAccountRepository
}

func NewAccountApplicationService(repository IAccountRepository) IAccountApplicationService {
	return &AccountApplicationService{
		repository: repository,
	}
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
func (s AccountApplicationService) RegisterAccount(email string, plainPassword string) (*AccountRegistrationResponseDTO, error) {
	newAccount, err := generateAccountEntity(email, plainPassword)
	if err != nil {
		return nil, err
	}
	createdAccount, err := s.repository.Add(*newAccount)
	if err != nil {
		return nil, err
	}

	// アカウント登録時にクライアントに返却するDTOを生成
	return &AccountRegistrationResponseDTO{
		ID:        createdAccount.ID().Value(),
		Email:     createdAccount.Email().Value(),
		CreatedAt: createdAccount.CreatedAt(),
	}, nil
}
