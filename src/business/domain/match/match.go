package match

import (
	"context"
	"fmt"
	"test_dealls/src/entity"
	"test_dealls/src/utils/log"
	"test_dealls/src/utils/sql"
)

type Interface interface {
	GetList(ctx context.Context, params entity.Matches) ([]entity.Matches, error)
	GetListByProfile(ctx context.Context, params entity.Matches) ([]entity.Matches, error)
	GetDetail(ctx context.Context, params entity.Matches) (entity.Matches, error)

	Create(ctx context.Context, params entity.Matches) (entity.Matches, error)
	Delete(ctx context.Context, params entity.Matches) (entity.Matches, error)
}

type Matches struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &Matches{
		log: log,
		db:  db,
	}
}

func (m *Matches) GetList(ctx context.Context, params entity.Matches) ([]entity.Matches, error) {
	rows, err := m.db.NamedQuery(ctx, "rMatchess", readMatch, params)
	if err != nil {
		return nil, err
	}

	Matchess := []entity.Matches{}
	for rows.Next() {
		Matches := entity.Matches{}
		if err := rows.StructScan(&Matches); err != nil {
			m.log.Error(ctx, err.Error())
			continue
		}

		Matchess = append(Matchess, Matches)
	}

	return Matchess, nil
}

func (m *Matches) GetListByProfile(ctx context.Context, params entity.Matches) ([]entity.Matches, error) {
	rows, err := m.db.Query(ctx, "rMatches", readMatchByProfile, params.Profile1, params.Profile2)
	if err != nil {
		return nil, err
	}

	Matchess := []entity.Matches{}
	for rows.Next() {
		Matches := entity.Matches{}
		if err := rows.StructScan(&Matches); err != nil {
			m.log.Error(ctx, err.Error())
			continue
		}

		Matchess = append(Matchess, Matches)
	}

	return Matchess, nil
}

func (m *Matches) GetDetail(ctx context.Context, params entity.Matches) (entity.Matches, error) {
	row, err := m.db.QueryRow(ctx, "rMatch", detailMatch, params.ID)
	if err != nil {
		return entity.Matches{}, err
	}

	Matches := entity.Matches{}
	if err := row.StructScan(&Matches); err != nil {
		return entity.Matches{}, err
	}

	return Matches, nil
}

func (m *Matches) Create(ctx context.Context, params entity.Matches) (entity.Matches, error) {
	res, err := m.db.NamedExec(ctx, "cMatches", createMatch, params)
	if err != nil {
		return entity.Matches{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.Matches{}, fmt.Errorf("no Matches created")
	}

	params.ID, err = res.LastInsertId()
	if err != nil {
		return entity.Matches{}, err
	}

	return params, nil
}

func (m *Matches) Delete(ctx context.Context, params entity.Matches) (entity.Matches, error) {
	res, err := m.db.NamedExec(ctx, "uMatches", delete, params)
	if err != nil {
		return entity.Matches{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.Matches{}, fmt.Errorf("no Matches deleted")
	}

	return params, nil
}
