package user

import (
	"time"
)

type GeneralUser struct {
	id        UserID
	email     UserEmail
	password  UserPassword
	createdAt time.Time
	updatedAt time.Time
	isDeleted bool
}

func NewGeneralUser(id UserID, email UserEmail, password UserPassword) GeneralUser {
	now := time.Now()
	return GeneralUser{
		id:        id,
		email:     email,
		password:  password,
		createdAt: now,
		updatedAt: now,
		isDeleted: false,
	}
}

func NewGeneralUserWithAllFields(
	id UserID,
	email UserEmail,
	password UserPassword,
	createdAt, updatedAt time.Time,
	isDeleted bool) GeneralUser {
	return GeneralUser{
		id:        id,
		email:     email,
		password:  password,
		createdAt: createdAt,
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
	return u.createdAt
}

func (u *GeneralUser) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *GeneralUser) IsDeleted() bool {
	return u.isDeleted
}

func (u *GeneralUser) MarkAsDeleted() {
	u.isDeleted = true
	u.updatedAt = time.Now()
}

func (u *GeneralUser) UpdateEmail(email UserEmail) {
	u.email = email
	u.updatedAt = time.Now()
}

func (u *GeneralUser) UpdatePassword(password UserPassword) {
	u.password = password
	u.updatedAt = time.Now()
}
