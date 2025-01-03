package account

import (
	"database/sql"
	"errors"
)

var ErrAccountNotFound = errors.New("account not found")

type IAccountRepository interface {
	// ユーザーIDでアカウントを取得
	FindByID(id AccountID) (*Account, error)

	// メールアドレスでアカウントを取得
	FindByEmail(email AccountEmail) (*FindByEmailResponseDTO, error)

	// 新しいユーザーを追加
	Add(account Account) error

	// 既存のユーザー情報を更新
	Update(account Account) (*Account, error)

	// ユーザーを削除
	Delete(id AccountID) (*Account, error)
}

type PostgresAccountRepository struct {
	db *sql.DB
}

func NewPostgresAccountRepository(db *sql.DB) IAccountRepository {
	return &PostgresAccountRepository{
		db: db,
	}
}

func (r PostgresAccountRepository) FindByID(id AccountID) (*Account, error) {
	return nil, nil
}

func (r PostgresAccountRepository) FindByEmail(email AccountEmail) (*FindByEmailResponseDTO, error) {
	// 専用のDTOを使ってDBから取得したデータを返却
	var dto FindByEmailResponseDTO
	query := `
		SELECT id, email, password
		FROM accounts
		WHERE email = $1
		AND is_deleted = false
	`
	err := r.db.QueryRow(query, email.Value()).Scan(
		&dto.ID,
		&dto.Email,
		&dto.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}
	return &dto, nil
}

func (r PostgresAccountRepository) Add(account Account) error {
	// パスワードをハッシュ化
	hashedPassword, err := account.Password().HashedValue()
	if err != nil {
		return err
	}

	// ユーザー情報をDBに追加
	query := `
		INSERT INTO accounts (id, email, password, created_at, updated_at, is_deleted)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = r.db.Exec(
		query,
		account.ID().Value(),
		account.Email().Value(),
		hashedPassword,
		account.CreatedAt(),
		account.UpdatedAt(),
		account.IsDeleted(),
	)
	return err
}

func (r PostgresAccountRepository) Update(account Account) (*Account, error) {
	return nil, nil
}

func (r PostgresAccountRepository) Delete(id AccountID) (*Account, error) {
	return nil, nil
}
