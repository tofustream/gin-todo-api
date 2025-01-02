package account

import (
	"time"

	"github.com/google/uuid"
)

// アカウント登録時にクライアントに返却するDTO
type AccountRegistrationResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type AccountFindByEmailResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
