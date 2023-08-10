package data

import (
	"bysos/app/user/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) ListUser(ctx context.Context) ([]*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) UpdateUser(ctx context.Context, id int64, user *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) DeleteUser(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
