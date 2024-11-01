package historywatch

import (
	"context"
	"fmt"
	"test_dealls/src/entity"
	"test_dealls/src/utils/log"
	"test_dealls/src/utils/sql"
)

type Interface interface {
	GetList(ctx context.Context, params entity.HistoryWatch) ([]entity.HistoryWatch, error)
	GetListByProfileID(ctx context.Context, params entity.HistoryWatch) ([]entity.HistoryWatch, error)

	Create(ctx context.Context, params entity.HistoryWatch) (entity.HistoryWatch, error)
	Update(ctx context.Context, params entity.HistoryWatch) (entity.HistoryWatch, error)
	Delete(ctx context.Context, params entity.HistoryWatch) (entity.HistoryWatch, error)
}

type HistoryWatch struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &HistoryWatch{
		log: log,
		db:  db,
	}
}

func (hw *HistoryWatch) GetList(ctx context.Context, params entity.HistoryWatch) ([]entity.HistoryWatch, error) {
	rows, err := hw.db.NamedQuery(ctx, "rHistoryWatchs", readHistoryWatch, params)
	if err != nil {
		return nil, err
	}

	HistoryWatchs := []entity.HistoryWatch{}
	for rows.Next() {
		HistoryWatch := entity.HistoryWatch{}
		if err := rows.StructScan(&HistoryWatch); err != nil {
			hw.log.Error(ctx, err.Error())
			continue
		}

		HistoryWatchs = append(HistoryWatchs, HistoryWatch)
	}

	return HistoryWatchs, nil
}

func (hw *HistoryWatch) GetListByProfileID(ctx context.Context, params entity.HistoryWatch) ([]entity.HistoryWatch, error) {
	rows, err := hw.db.NamedQuery(ctx, "rHistoryWatchs", readHistoryWatchByProfileID, params)
	if err != nil {
		return nil, err
	}

	HistoryWatchs := []entity.HistoryWatch{}
	for rows.Next() {
		HistoryWatch := entity.HistoryWatch{}
		if err := rows.StructScan(&HistoryWatch); err != nil {
			hw.log.Error(ctx, err.Error())
			continue
		}

		HistoryWatchs = append(HistoryWatchs, HistoryWatch)
	}

	return HistoryWatchs, nil
}

func (hw *HistoryWatch) Create(ctx context.Context, params entity.HistoryWatch) (entity.HistoryWatch, error) {
	res, err := hw.db.NamedExec(ctx, "cHistoryWatch", createHistoryWatch, params)
	if err != nil {
		return entity.HistoryWatch{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.HistoryWatch{}, fmt.Errorf("no HistoryWatch created")
	}

	params.ID, err = res.LastInsertId()
	if err != nil {
		return entity.HistoryWatch{}, err
	}

	return params, nil
}

func (hw *HistoryWatch) Update(ctx context.Context, params entity.HistoryWatch) (entity.HistoryWatch, error) {
	res, err := hw.db.NamedExec(ctx, "uHistoryWatch", updateHistoryWatch, params)
	if err != nil {
		return entity.HistoryWatch{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.HistoryWatch{}, fmt.Errorf("no HistoryWatch updated")
	}

	return params, nil
}

func (hw *HistoryWatch) Delete(ctx context.Context, params entity.HistoryWatch) (entity.HistoryWatch, error) {
	res, err := hw.db.NamedExec(ctx, "uHistoryWatch", delete, params)
	if err != nil {
		return entity.HistoryWatch{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.HistoryWatch{}, fmt.Errorf("no HistoryWatch deleted")
	}

	return params, nil
}
