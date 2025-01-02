package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tofustream/gin-todo-api/cmd/internal/account"
	"golang.org/x/crypto/bcrypt"
)

type IAuthApplicationService interface {
	Login(email string, password string) (*string, error)
}

type AuthApplicationService struct {
	repository account.IAccountRepository
}

func NewAuthApplicationService(repository account.IAccountRepository) IAuthApplicationService {
	return &AuthApplicationService{
		repository: repository,
	}
}

func (s *AuthApplicationService) Login(email string, password string) (*string, error) {
	emailValue, err := account.NewAccountEmail(email)
	if err != nil {
		return nil, err
	}
	dto, err := s.repository.FindByEmail(emailValue)
	if err != nil {
		return nil, err
	}

	passwordValue, err := account.NewAccountPassword(password)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dto.Password), []byte(passwordValue.Plain())); err != nil {
		return nil, errors.New("invalid credentials")
	}

	accountIDValue, err := account.NewAccountIDFromUUID(dto.ID)
	if err != nil {
		return nil, err
	}
	token, err := createToken(accountIDValue, emailValue)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// JWTトークンを生成
func createToken(accountID account.AccountID, email account.AccountEmail) (*string, error) {
	// トークンの有効期限は1時間
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   accountID.Value(),
		"email": email.Value(),
		"exp":   time.Now().Add(time.Hour).Unix(), // 1時間有効
	})

	// 環境変数から取得したSECRET_KEYを使ってトークンを署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
