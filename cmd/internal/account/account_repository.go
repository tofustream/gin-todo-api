package account

import "database/sql"

type IAccountRepository interface {
	// ユーザーIDでアカウントを取得
	FindByID(id AccountID) (*Account, error)

	// メールアドレスでアカウントを取得
	FindByEmail(email AccountEmail) (*Account, error)

	// 新しいユーザーを追加
	Add(account Account) (*Account, error)

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

func (r PostgresAccountRepository) FindByEmail(email AccountEmail) (*Account, error) {
	return nil, nil
}

func (r PostgresAccountRepository) Add(account Account) (*Account, error) {
	// パスワードをハッシュ化
	hashedPassword, err := account.Password().HashedValue()
	if err != nil {
		return nil, err
	}

	// ユーザー情報をDBに追加
	_, err = r.db.Exec("INSERT INTO accounts (id, email, password, created_at, updated_at, is_deleted) VALUES ($1, $2, $3, $4, $5, $6)",
		account.ID().Value(),
		account.Email().Value(),
		hashedPassword,
		account.CreatedAt(),
		account.UpdatedAt(),
		account.IsDeleted(),
	)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r PostgresAccountRepository) Update(account Account) (*Account, error) {
	return nil, nil
}

func (r PostgresAccountRepository) Delete(id AccountID) (*Account, error) {
	return nil, nil
}
