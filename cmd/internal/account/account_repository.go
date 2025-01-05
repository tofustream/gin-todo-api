package account

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

var ErrAccountNotFound = errors.New("account not found")

type IAccountRepository interface {
	// ユーザーIDでアカウントを取得
	FindAccount(id AccountID) (*AccountDTO, error)

	// メールアドレスでアカウントを取得
	FindAccountByEmail(email AccountEmail) (*FindByEmailResponseDTO, error)

	// 新しいユーザーを追加
	AddAccount(account Account) error

	// 既存のユーザー情報を更新
	UpdateAccount(updatedAccount UpdatedAccount) error
}

type PostgresAccountRepository struct {
	db *sql.DB
}

func NewPostgresAccountRepository(db *sql.DB) IAccountRepository {
	return &PostgresAccountRepository{
		db: db,
	}
}

func (r PostgresAccountRepository) FindAccount(id AccountID) (*AccountDTO, error) {
	// ユーザーIDを使ってDBから取得したデータを返却
	var dto AccountDTO
	query := `
		SELECT id, email, password, created_at, updated_at, is_deleted
		FROM accounts
		WHERE id = $1
		AND is_deleted = false
	`
	err := r.db.QueryRow(query, id.Value()).Scan(
		&dto.ID,
		&dto.Email,
		&dto.Password,
		&dto.CreatedAt,
		&dto.UpdatedAt,
		&dto.IsDeleted,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	return &dto, nil
}

func (r PostgresAccountRepository) FindAccountByEmail(email AccountEmail) (*FindByEmailResponseDTO, error) {
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
	log.Printf("account_id: %s", dto.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}
	return &dto, nil
}

func (r PostgresAccountRepository) AddAccount(account Account) error {
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

func (r PostgresAccountRepository) UpdateAccount(updatedAccount UpdatedAccount) error {
	query := `
		UPDATE accounts
		SET email = $1, password = $2, updated_at = $3, is_deleted = $4
		WHERE id = $5
	`
	_, err := r.db.Exec(
		query,
		updatedAccount.Email().Value(),
		updatedAccount.HashedPassword().Value(),
		updatedAccount.UpdatedAt().Format(time.RFC3339),
		updatedAccount.IsDeleted(),
		updatedAccount.ID().String(),
	)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}
