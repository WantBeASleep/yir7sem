package repository

import (
	"github.com/WantBeASleep/goooool/daolib"

	"auth/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const userTable = "\"user\""

type UserQuery interface {
	InsertUser(user entity.User) error
	GetUserByPK(id uuid.UUID) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	UpdateUser(user entity.User) (int64, error)
}

type userQuery struct {
	*daolib.BaseQuery
}

func (q *userQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *userQuery) InsertUser(user entity.User) error {
	query := q.QueryBuilder().
		Insert(userTable).
		Columns(
			"id",
			"email",
			"password",
			"token",
		).
		Values(
			user.Id,
			user.Email,
			user.Password,
			user.Token,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *userQuery) GetUserByPK(id uuid.UUID) (entity.User, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"email",
			"password",
			"token",
		).
		From(userTable).
		Where(sq.Eq{
			"id": id,
		})

	var user entity.User
	if err := q.Runner().Getx(q.Context(), &user, query); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (q *userQuery) GetUserByEmail(email string) (entity.User, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"email",
			"password",
			"token",
		).
		From(userTable).
		Where(sq.Eq{
			"email": email,
		})

	var user entity.User
	if err := q.Runner().Getx(q.Context(), &user, query); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (q *userQuery) UpdateUser(user entity.User) (int64, error) {
	query := q.QueryBuilder().
		Update(userTable).
		SetMap(sq.Eq{
			"password": user.Password,
			"token":    user.Token,
		}).
		Where(sq.Eq{
			"id": user.Id,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
