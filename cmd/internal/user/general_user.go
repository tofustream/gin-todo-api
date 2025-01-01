package user

import "time"

type GeneralUser struct {
	id        UserID
	email     UserEmail
	password  UserPassword
	createAt  time.Time
	updatedAt time.Time
	isDeleted bool
}

func NewGeneralUser(id UserID, email UserEmail, password UserPassword) GeneralUser {
	now := time.Now()
	return GeneralUser{
		id:        id,
		email:     email,
		password:  password,
		createAt:  now,
		updatedAt: now,
		isDeleted: false}
}

func NewGeneralUserWithAllFields(
	id UserID,
	email UserEmail,
	password UserPassword,
	createAt, updatedAt time.Time,
	isDeleted bool) GeneralUser {
	return GeneralUser{
		id:        id,
		email:     email,
		password:  password,
		createAt:  createAt,
		updatedAt: updatedAt,
		isDeleted: isDeleted,
	}
}

func (u *GeneralUser) ID() UserID {
	return u.id
}

func (u *GeneralUser) Email() UserEmail {
	return u.email
}

func (u *GeneralUser) Password() UserPassword {
	return u.password
}

func (u *GeneralUser) CreatedAt() time.Time {
	return u.createAt
}

func (u *GeneralUser) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *GeneralUser) IsDeleted() bool {
	return u.isDeleted
}

func (u *GeneralUser) MarkAsDeleted() GeneralUser {
	return GeneralUser{
		id:        u.id,
		email:     u.email,
		createAt:  u.createAt,
		updatedAt: time.Now(),
		isDeleted: true,
	}
}
