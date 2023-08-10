package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Info is a Info model.
type Info struct {
	Name string
}

// InfoRepo is a Greater repo.
type InfoRepo interface {
	Save(context.Context, *Info) (*Info, error)
	Update(context.Context, *Info) (*Info, error)
	FindByID(context.Context, int64) (*Info, error)
	ListByHello(context.Context, string) ([]*Info, error)
	ListAll(context.Context) ([]*Info, error)
}

// InfoUsecase is a Info usecase.
type InfoUsecase struct {
	repo InfoRepo
	log  *log.Helper
}

// NewInfoUsecase new a Info usecase.
func NewInfoUsecase(repo InfoRepo, logger log.Logger) *InfoUsecase {
	return &InfoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// TestHttp creates a Info, and returns the new Info.
func (uc *InfoUsecase) TestHttp(ctx context.Context, g *Info) (*Info, error) {
	uc.log.WithContext(ctx).Infof("TestHttp: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

func (uc *InfoUsecase) TestWRedis(ctx context.Context, g *Info) (*Info, error) {
	uc.log.WithContext(ctx).Infof("TestRedis write: %v", g.Name)
	return uc.repo.Update(ctx, g)
}

func (uc *InfoUsecase) TestRRedis(ctx context.Context, g *Info) (*Info, error) {
	uc.log.WithContext(ctx).Infof("TestRedis read: %v", g.Name)
	return uc.repo.FindByID(ctx, 0)
}
