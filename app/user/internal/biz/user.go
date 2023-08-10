package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// model
type User struct {
	Name     string
	PassWord string
	Id       int
}

type UserRepo interface {
	//	包含db redis等接口方法
	ListUser(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id int64, user *User) error
	DeleteUser(ctx context.Context, id int64) error
}

type UserUseCase struct {
	//	接口用例	需要用方法放到此处
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

// 接口实现
func (uc *UserUseCase) Create(ctx context.Context, g *User) error {
	return uc.repo.CreateUser(ctx, g)
}
