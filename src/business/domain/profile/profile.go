package profile

import (
	"context"
	"fmt"
	"test_dealls/src/entity"
	"test_dealls/src/utils/log"
	"test_dealls/src/utils/sql"
)

type Interface interface {
	GetList(ctx context.Context, params entity.Profile) ([]entity.Profile, error)
	SearchProfilePeople(ctx context.Context, params entity.Profile) ([]entity.Profile, error)
	GetDetail(ctx context.Context, params entity.Profile) (entity.Profile, error)
	GetDetailByUserID(ctx context.Context, params entity.Profile) (entity.Profile, error)

	Create(ctx context.Context, params entity.Profile) (entity.Profile, error)
	Update(ctx context.Context, params entity.Profile) (entity.Profile, error)
}

type Profile struct {
	log log.Interface
	db  sql.Interface
}

func Init(log log.Interface, db sql.Interface) Interface {
	return &Profile{
		log: log,
		db:  db,
	}
}

func (p *Profile) GetList(ctx context.Context, params entity.Profile) ([]entity.Profile, error) {
	rows, err := p.db.NamedQuery(ctx, "rProfiles", readProfile, params)
	if err != nil {
		return nil, err
	}

	Profiles := []entity.Profile{}
	for rows.Next() {
		Profile := entity.Profile{}
		if err := rows.StructScan(&Profile); err != nil {
			p.log.Error(ctx, err.Error())
			continue
		}

		Profiles = append(Profiles, Profile)
	}

	return Profiles, nil
}

func (p *Profile) SearchProfilePeople(ctx context.Context, params entity.Profile) ([]entity.Profile, error) {
	rows, err := p.db.Query(ctx, "rProfilesPeople", SearchProfilePeople, params.UserID, params.Gender)
	if err != nil {
		return nil, err
	}

	Profiles := []entity.Profile{}
	for rows.Next() {
		Profile := entity.Profile{}
		if err := rows.StructScan(&Profile); err != nil {
			p.log.Error(ctx, err.Error())
			continue
		}

		Profiles = append(Profiles, Profile)
	}

	return Profiles, nil
}

func (p *Profile) GetDetail(ctx context.Context, params entity.Profile) (entity.Profile, error) {
	row, err := p.db.QueryRow(ctx, "rProfile", detailProfile, params.ID)
	if err != nil {
		return entity.Profile{}, err
	}

	Profile := entity.Profile{}
	if err := row.StructScan(&Profile); err != nil {
		return entity.Profile{}, err
	}

	return Profile, nil
}

func (p *Profile) GetDetailByUserID(ctx context.Context, params entity.Profile) (entity.Profile, error) {
	row, err := p.db.QueryRow(ctx, "rProfile", detailProfileByUserID, params.UserID)
	if err != nil {
		return entity.Profile{}, err
	}

	Profile := entity.Profile{}
	if err := row.StructScan(&Profile); err != nil {
		return entity.Profile{}, err
	}

	return Profile, nil
}

func (p *Profile) Create(ctx context.Context, params entity.Profile) (entity.Profile, error) {
	res, err := p.db.NamedExec(ctx, "cProfile", createProfile, params)
	if err != nil {
		return entity.Profile{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.Profile{}, fmt.Errorf("no Profile created")
	}

	params.ID, err = res.LastInsertId()
	if err != nil {
		return entity.Profile{}, err
	}

	return params, nil
}

func (p *Profile) Update(ctx context.Context, params entity.Profile) (entity.Profile, error) {
	res, err := p.db.NamedExec(ctx, "uProfile", updateProfile, params)
	if err != nil {
		return entity.Profile{}, err
	}

	num, err := res.RowsAffected()
	if err != nil || num < 1 {
		return entity.Profile{}, fmt.Errorf("no Profile updated")
	}

	return params, nil
}
