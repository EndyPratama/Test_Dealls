package user

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	userDom "test_dealls/src/business/domain/user"
	"test_dealls/src/entity"
	"test_dealls/src/utils/appcontext"
	"test_dealls/src/utils/log"
)

type Interface interface {
	GetDetail(ctx context.Context, params entity.User) (entity.User, error)

	Login(ctx context.Context, params entity.User) (entity.User, string, error)
	Register(ctx context.Context, params entity.User) (entity.User, error)
}

type User struct {
	log log.Interface
	dom domain
}

type domain struct {
	user userDom.Interface
}

func Init(log log.Interface, userD userDom.Interface) Interface {
	return &User{
		log: log,
		dom: domain{
			user: userD,
		},
	}
}

func (u *User) GetDetail(ctx context.Context, params entity.User) (entity.User, error) {
	u.log.Info(ctx, "Get User by ID")

	user, err := u.dom.user.GetDetail(ctx, params)
	if err != nil {
		u.log.Error(ctx, err.Error())
		return entity.User{}, err
	}

	return user, nil
}

func (u *User) Login(ctx context.Context, params entity.User) (entity.User, string, error) {
	u.log.Info(ctx, "Login user")

	hash := sha256.New()
	hash.Write([]byte(params.Password))

	params.Password = fmt.Sprintf("%x", hash.Sum(nil))
	user, err := u.dom.user.Login(ctx, params)
	if err != nil {
		u.log.Error(ctx, err.Error())
		return entity.User{}, "", err
	}

	data := fmt.Sprintf("%v:%v", user.ID, user.Email)
	token := base64.StdEncoding.EncodeToString([]byte(data))

	return user, token, nil
}

func (u *User) Register(ctx context.Context, params entity.User) (entity.User, error) {
	u.log.Info(ctx, "Register user")

	hash := sha256.New()
	hash.Write([]byte(params.Password))

	params.Password = fmt.Sprintf("%x", hash.Sum(nil))
	params.SubscriptionID = 1
	params.CreatedAt = appcontext.GetRequestStartTime(ctx)
	user, err := u.dom.user.Create(ctx, params)
	if err != nil {
		u.log.Error(ctx, err.Error())
		return entity.User{}, err
	}

	return user, nil
}
