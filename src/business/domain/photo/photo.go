package photo

import (
	"context"
	"fmt"
	"test_dealls/src/entity"
	"test_dealls/src/utils/log"
	"test_dealls/src/utils/sql"
)

type Interface interface {
	GetList(ctx context.Context, params entity.Photo) ([]entity.Photo, error)
	GetDetail(ctx context.Context, params entity.Photo) (entity.Photo, error)

	Create(ctx context.Context, params entity.Photo) (entity.Photo, error)
	Update(ctx context.Context, params entity.Photo) (entity.Photo, error)
}

type Photo struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &Photo{
		log: log,
		db:  db,
	}
}

func (p *Photo) GetList(ctx context.Context, params entity.Photo) ([]entity.Photo, error) {
	rows, err := p.db.NamedQuery(ctx, "rPhotos", readPhoto, params)
	if err != nil {
		return nil, err
	}

	Photos := []entity.Photo{}
	for rows.Next() {
		Photo := entity.Photo{}
		if err := rows.StructScan(&Photo); err != nil {
			p.log.Error(ctx, err.Error())
			continue
		}

		Photos = append(Photos, Photo)
	}

	return Photos, nil
}

func (p *Photo) GetDetail(ctx context.Context, params entity.Photo) (entity.Photo, error) {
	row, err := p.db.QueryRow(ctx, "rPhoto", detailPhoto, params.ID)
	if err != nil {
		return entity.Photo{}, err
	}

	Photo := entity.Photo{}
	if err := row.StructScan(&Photo); err != nil {
		return entity.Photo{}, err
	}

	return Photo, nil
}

func (p *Photo) Create(ctx context.Context, params entity.Photo) (entity.Photo, error) {
	res, err := p.db.NamedExec(ctx, "cPhoto", createPhoto, params)
	if err != nil {
		return entity.Photo{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.Photo{}, fmt.Errorf("no Photo created")
	}

	params.ID, err = res.LastInsertId()
	if err != nil {
		return entity.Photo{}, err
	}

	return params, nil
}

func (p *Photo) Update(ctx context.Context, params entity.Photo) (entity.Photo, error) {
	res, err := p.db.NamedExec(ctx, "uPhoto", updatePhoto, params)
	if err != nil {
		return entity.Photo{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.Photo{}, fmt.Errorf("no Photo updated")
	}

	return params, nil
}
