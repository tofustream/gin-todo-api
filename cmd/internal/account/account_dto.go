package account

import (
	"time"

	"github.com/google/uuid"
)

type FindAccountByEmailResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// パスワードは公開しない
type FetchedAccountDTO struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
