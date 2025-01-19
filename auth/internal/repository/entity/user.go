package entity

import (
	"database/sql"

	"github.com/WantBeASleep/med_ml_lib/gtc"

	"auth/internal/domain"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID      `db:"id"`
	Email    string         `db:"email"`
	Password string         `db:"password"`
	Token    sql.NullString `db:"token"`
}

func (e User) ToDomain() domain.User {
	return domain.User{
		Id:       e.Id,
		Email:    e.Email,
		Password: e.Password,
		Token:    gtc.String.SqlToPointer(e.Token),
	}
}

func (User) FromDomain(d domain.User) User {
	return User{
		Id:       d.Id,
		Email:    d.Email,
		Password: d.Password,
		Token:    gtc.String.PointerToSql(d.Token),
	}
}
