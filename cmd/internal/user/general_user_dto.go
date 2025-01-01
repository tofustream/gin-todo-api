package user

type GeneralUserDTO struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsDeleted bool   `json:"is_deleted"`
}

func NewGeneralUserDTO(
	id string,
	email string,
	password string,
	createdAt string,
	updatedAt string,
	isDeleted bool) GeneralUserDTO {
	return GeneralUserDTO{
		ID:        id,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		IsDeleted: isDeleted,
	}
}
