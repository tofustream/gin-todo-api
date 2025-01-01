package user

import "database/sql"

type IGeneralUserRepository interface {
	FindAll() ([]GeneralUserDTO, error)
	Add(user GeneralUser) error
}

type PostgresGeneralUserRepository struct {
	db *sql.DB
}

func NewPostgresGeneralUserRepository(db *sql.DB) IGeneralUserRepository {
	return &PostgresGeneralUserRepository{
		db: db,
	}
}

func (r *PostgresGeneralUserRepository) FindAll() ([]GeneralUserDTO, error) {
	rows, err := r.db.Query("SELECT * FROM general_users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []GeneralUserDTO
	for rows.Next() {
		var dto GeneralUserDTO
		err := rows.Scan(&dto.ID, &dto.Email, &dto.Password, &dto.CreatedAt, &dto.UpdatedAt, &dto.IsDeleted)
		if err != nil {
			return nil, err
		}
		users = append(users, dto)
	}

	return users, nil
}

func (r *PostgresGeneralUserRepository) Add(user GeneralUser) error {
	_, err := r.db.Exec("INSERT INTO general_users (id, email, password, created_at, updated_at, is_deleted) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID().Value(),
		user.Email().Value(),
		user.Password().Value(),
		user.CreatedAt(),
		user.UpdatedAt(),
		user.IsDeleted())
	if err != nil {
		return err
	}

	return nil
}
