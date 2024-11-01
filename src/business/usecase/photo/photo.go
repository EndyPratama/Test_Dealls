package photo

import (
	"context"
	"strconv"
	photoDom "test_dealls/src/business/domain/photo"
	"test_dealls/src/entity"
	"test_dealls/src/utils/appcontext"
	"test_dealls/src/utils/log"
)

type Interface interface {
	GetList(ctx context.Context, params entity.Photo) ([]entity.Photo, error)
	GetDetail(ctx context.Context, params entity.Photo) (entity.Photo, error)

	Create(ctx context.Context, params entity.Photo) (entity.Photo, error)
	Update(ctx context.Context, params entity.Photo) (entity.Photo, error)
}

type Photo struct {
	log log.Interface
	dom domain
}

type domain struct {
	photo photoDom.Interface
}

func Init(log log.Interface, photoD photoDom.Interface) Interface {
	return &Photo{
		log: log,
		dom: domain{
			photo: photoD,
		},
	}
}

func (p *Photo) GetList(ctx context.Context, params entity.Photo) ([]entity.Photo, error) {
	p.log.Info(ctx, "Get list photo")

	listPhoto, err := p.dom.photo.GetList(ctx, params)
	if err != nil {
		p.log.Error(ctx, err.Error())
		return nil, err
	}

	return listPhoto, nil
}

func (p *Photo) GetDetail(ctx context.Context, params entity.Photo) (entity.Photo, error) {
	p.log.Info(ctx, "Get Photo by ID")

	id, err := strconv.Atoi(appcontext.GetUserIDAgent(ctx))
	if err != nil {
		p.log.Error(ctx, err.Error())
	}
	params.ID = int64(id)

	photo, err := p.dom.photo.GetDetail(ctx, params)
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Photo{}, err
	}

	return photo, nil
}

func (p *Photo) Create(ctx context.Context, params entity.Photo) (entity.Photo, error) {
	p.log.Info(ctx, "Create photo")

	params.CreatedAt = appcontext.GetRequestStartTime(ctx)
	photo, err := p.dom.photo.Create(ctx, params)
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Photo{}, err
	}

	return photo, nil
}

func (p *Photo) Update(ctx context.Context, params entity.Photo) (entity.Photo, error) {
	p.log.Info(ctx, "Update photo")

	params.UpdatedAt = appcontext.GetRequestStartTime(ctx)
	photo, err := p.dom.photo.Update(ctx, params)
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Photo{}, err
	}

	return photo, nil
}
