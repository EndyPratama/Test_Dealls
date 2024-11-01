package subscription

import (
	"context"
	"fmt"
	"test_dealls/src/entity"
	"test_dealls/src/utils/log"
	"test_dealls/src/utils/sql"
)

type Interface interface {
	GetList(ctx context.Context, params entity.Subscription) ([]entity.Subscription, error)
	GetDetail(ctx context.Context, params entity.Subscription) (entity.Subscription, error)

	Create(ctx context.Context, params entity.Subscription) (entity.Subscription, error)
	Update(ctx context.Context, params entity.Subscription) (entity.Subscription, error)
}

type Subscription struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &Subscription{
		log: log,
		db:  db,
	}
}

func (s *Subscription) GetList(ctx context.Context, params entity.Subscription) ([]entity.Subscription, error) {
	rows, err := s.db.NamedQuery(ctx, "rSubscriptions", readSubscription, params)
	if err != nil {
		return nil, err
	}

	Subscriptions := []entity.Subscription{}
	for rows.Next() {
		Subscription := entity.Subscription{}
		if err := rows.StructScan(&Subscription); err != nil {
			s.log.Error(ctx, err.Error())
			continue
		}

		Subscriptions = append(Subscriptions, Subscription)
	}

	return Subscriptions, nil
}

func (s *Subscription) GetDetail(ctx context.Context, params entity.Subscription) (entity.Subscription, error) {
	row, err := s.db.QueryRow(ctx, "rSubscription", detailSubscription, params.ID)
	if err != nil {
		return entity.Subscription{}, err
	}

	Subscription := entity.Subscription{}
	if err := row.StructScan(&Subscription); err != nil {
		return entity.Subscription{}, err
	}

	return Subscription, nil
}

func (s *Subscription) Create(ctx context.Context, params entity.Subscription) (entity.Subscription, error) {
	res, err := s.db.NamedExec(ctx, "cSubscription", createSubscription, params)
	if err != nil {
		return entity.Subscription{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.Subscription{}, fmt.Errorf("no Subscription created")
	}

	params.ID, err = res.LastInsertId()
	if err != nil {
		return entity.Subscription{}, err
	}

	return params, nil
}

func (s *Subscription) Update(ctx context.Context, params entity.Subscription) (entity.Subscription, error) {
	res, err := s.db.NamedExec(ctx, "uSubscription", updateSubscription, params)
	if err != nil {
		return entity.Subscription{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.Subscription{}, fmt.Errorf("no Subscription updated")
	}

	return params, nil
}
