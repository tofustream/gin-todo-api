package user

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IGeneralUserApplicationService interface {
	FindAll() ([]GeneralUserDTO, error)
	Signup(email string, password string) error
	Login(email string, password string) (*string, error)
	GetUserFromToken(tokenString string) (*GeneralUserDTO, error)
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
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dto.Password), []byte(password)); err != nil {
		return nil, err
	}

	fetchedUserID, err := NewUserIDFromString(dto.ID)
	if err != nil {
		return nil, err
	}

	token, err := createToken(fetchedUserID, emailValue)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func createToken(userID UserID, email UserEmail) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userID.Value(),
		"email": email.Value(),
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (s *GeneralUserApplicationService) GetUserFromToken(tokenString string) (*GeneralUserDTO, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv(os.Getenv("SECRET_KEY"))), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unexpected claims type: %T", claims)
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, jwt.ErrTokenExpired
	}

	dto, err := s.repository.FindUser(claims["email"].(string))
	if err != nil {
		return nil, err
	}

	return &dto, nil
}
