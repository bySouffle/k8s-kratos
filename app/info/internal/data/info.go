package data

import (
	v1 "bysos/api/user/v1"
	"context"

	"bysos/app/info/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type InfoRepo struct {
	data *Data
	log  *log.Helper
}

// NewInfoRepo .
func NewInfoRepo(data *Data, logger log.Logger) biz.InfoRepo {
	return &InfoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *InfoRepo) Save(ctx context.Context, g *biz.Info) (*biz.Info, error) {
	req := &v1.GetUserRequest{
		Id:   "1",
		Name: "1",
	}
	resp, err := r.data.httpUser.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	log.Info(resp)
	g.Name = resp.Msg
	return g, nil
}

func (r *InfoRepo) Update(ctx context.Context, g *biz.Info) (*biz.Info, error) {
	err := r.data.redis.Set(ctx, "v1:info:update", "{update:1}", 0).Err()
	if err != nil {
		log.Error(err)
	}
	return g, nil
}

func (r *InfoRepo) FindByID(ctx context.Context, id int64) (*biz.Info, error) {
	val := r.data.redis.Get(ctx, "v1:info:update")
	info := &biz.Info{}
	if val.Err() != nil {
		info.Name = ""
		return info, nil
	}
	info.Name = val.Val()
	return info, nil
}

func (r *InfoRepo) ListByHello(context.Context, string) ([]*biz.Info, error) {
	return nil, nil
}

func (r *InfoRepo) ListAll(context.Context) ([]*biz.Info, error) {
	return nil, nil
}
