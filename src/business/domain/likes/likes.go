package likes

import (
	"context"
	"fmt"
	"test_dealls/src/entity"
	"test_dealls/src/utils/log"
	"test_dealls/src/utils/sql"
)

type Interface interface {
	GetList(ctx context.Context, params entity.Likes) ([]entity.Likes, error)
	GetDetail(ctx context.Context, params entity.Likes) (entity.Likes, error)
	GetListByLikerID(ctx context.Context, params entity.Likes) ([]entity.Likes, error)
	CheckLikes(ctx context.Context, params entity.Likes) (entity.Likes, error)

	Create(ctx context.Context, params entity.Likes) (entity.Likes, error)
	Delete(ctx context.Context, params entity.Likes) (entity.Likes, error)
}

type Likes struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &Likes{
		log: log,
		db:  db,
	}
}

func (l *Likes) GetList(ctx context.Context, params entity.Likes) ([]entity.Likes, error) {
	rows, err := l.db.NamedQuery(ctx, "rLikess", readLikes, params)
	if err != nil {
		return nil, err
	}

	Likess := []entity.Likes{}
	for rows.Next() {
		Likes := entity.Likes{}
		if err := rows.StructScan(&Likes); err != nil {
			l.log.Error(ctx, err.Error())
			continue
		}

		Likess = append(Likess, Likes)
	}

	return Likess, nil
}

func (l *Likes) GetDetail(ctx context.Context, params entity.Likes) (entity.Likes, error) {
	row, err := l.db.QueryRow(ctx, "rLikes", detailLikes, params.ID)
	if err != nil {
		return entity.Likes{}, err
	}

	Likes := entity.Likes{}
	if err := row.StructScan(&Likes); err != nil {
		return entity.Likes{}, err
	}

	return Likes, nil
}

func (l *Likes) CheckLikes(ctx context.Context, params entity.Likes) (entity.Likes, error) {
	row, err := l.db.QueryRow(ctx, "rLikes", checkLikes, params.LikerID, params.LikedID)
	if err != nil {
		return entity.Likes{}, err
	}

	Likes := entity.Likes{}
	if err := row.StructScan(&Likes); err != nil {
		return entity.Likes{}, err
	}

	return Likes, nil
}

func (l *Likes) GetListByLikerID(ctx context.Context, params entity.Likes) ([]entity.Likes, error) {
	rows, err := l.db.Query(ctx, "rLikeMatch", readLikesByLikerID, params.LikerID)
	if err != nil {
		return nil, err
	}

	likes := []entity.Likes{}
	for rows.Next() {
		like := entity.Likes{}
		if err := rows.StructScan(&like); err != nil {
			l.log.Error(ctx, err.Error())
			continue
		}

		likes = append(likes, like)
	}

	return likes, nil
}

func (l *Likes) Create(ctx context.Context, params entity.Likes) (entity.Likes, error) {
	res, err := l.db.NamedExec(ctx, "cLikes", createLikes, params)
	if err != nil {
		return entity.Likes{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.Likes{}, fmt.Errorf("no Likes created")
	}

	params.ID, err = res.LastInsertId()
	if err != nil {
		return entity.Likes{}, err
	}

	return params, nil
}

func (l *Likes) Delete(ctx context.Context, params entity.Likes) (entity.Likes, error) {
	res, err := l.db.NamedExec(ctx, "dLikes", delete, params)
	if err != nil {
		return entity.Likes{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.Likes{}, fmt.Errorf("no Likes deleted")
	}

	return params, nil
}
