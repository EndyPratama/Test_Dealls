package user

import (
	"context"
	"fmt"
	"test_dealls/src/entity"
	"test_dealls/src/utils/log"
	"test_dealls/src/utils/sql"
)

type Interface interface {
	GetList(ctx context.Context, params entity.User) ([]entity.User, error)
	GetDetail(ctx context.Context, params entity.User) (entity.User, error)
	Login(ctx context.Context, params entity.User) (entity.User, error)
	Create(ctx context.Context, params entity.User) (entity.User, error)
}

type User struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &User{
		log: log,
		db:  db,
	}
}

func (u *User) GetList(ctx context.Context, params entity.User) ([]entity.User, error) {
	rows, err := u.db.NamedQuery(ctx, "rUsers", readUser, params)
	if err != nil {
		return nil, err
	}

	users := []entity.User{}
	for rows.Next() {
		user := entity.User{}
		if err := rows.StructScan(&user); err != nil {
			u.log.Error(ctx, err.Error())
			continue
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *User) GetDetail(ctx context.Context, params entity.User) (entity.User, error) {
	row, err := u.db.QueryRow(ctx, "rUser", detailUser, params.ID)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{}
	if err := row.StructScan(&user); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *User) Login(ctx context.Context, params entity.User) (entity.User, error) {
	row, err := u.db.QueryRow(ctx, "rUser", login, params.Email, params.Password)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{}
	if err := row.StructScan(&user); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u *User) Create(ctx context.Context, params entity.User) (entity.User, error) {
	res, err := u.db.NamedExec(ctx, "cUser", createUser, params)
	if err != nil {
		return entity.User{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.User{}, fmt.Errorf("no user created")
	}

	params.ID, err = res.LastInsertId()
	if err != nil {
		return entity.User{}, err
	}

	return params, nil
}
