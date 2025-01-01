package user

import (
	"github.com/google/uuid"
)

type IGeneralUserApplicationService interface {
	FindAll() ([]GeneralUserDTO, error)
	Register(email string, password string) error
}

type GeneralUserApplicationService struct {
	repository IGeneralUserRepository
}

func NewGeneralUserApplicationService(repository IGeneralUserRepository) IGeneralUserApplicationService {
	return &GeneralUserApplicationService{
		repository: repository,
	}
}

func (s *GeneralUserApplicationService) FindAll() ([]GeneralUserDTO, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *GeneralUserApplicationService) Register(email string, password string) error {
	emailValue, err := NewUserEmail(email)
	if err != nil {
		return err
	}

	passwordValue, err := NewUserPassword(password)
	if err != nil {
		return err
	}

	generatedUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	userID, err := NewUserID(generatedUUID)
	if err != nil {
		return err
	}

	user := NewGeneralUser(userID, emailValue, passwordValue)
	err = s.repository.Add(user)
	if err != nil {
		return err
	}

	return nil
}
