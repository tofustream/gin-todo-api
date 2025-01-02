package user

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IGeneralUserApplicationService interface {
	FindAll() ([]GeneralUserDTO, error)
	Signup(email string, password string) error
	Login(email string, password string) (*string, error)
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

func (s *GeneralUserApplicationService) Signup(email string, password string) error {
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

func (s *GeneralUserApplicationService) Login(email string, password string) (*string, error) {
	emailValue, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}

	dto, err := s.repository.FindUser(emailValue.Value())
	if err != nil {
		log.Printf("find user error: %v\n", err)
		return nil, err
	}

	log.Printf("Comparing hashed password: %s with input password: %s", dto.Password, password)
	if err := bcrypt.CompareHashAndPassword([]byte(dto.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &dto.Email, nil
}
