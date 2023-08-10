package service

import (
	"bysos/app/info/internal/biz"
	"context"

	pb "bysos/api/info/v1"
)

type InfoService struct {
	pb.UnimplementedInfoServer
	uc *biz.InfoUsecase
}

func NewInfoService(uc *biz.InfoUsecase) *InfoService {
	return &InfoService{
		uc: uc,
	}
}

func (s *InfoService) CreateInfo(ctx context.Context, req *pb.CreateInfoRequest) (*pb.CreateInfoReply, error) {
	info := biz.Info{Name: "zzz"}
	query, err := s.uc.TestHttp(ctx, &info)
	if err != nil {
		return &pb.CreateInfoReply{
			Data: &pb.Resp{
				Code: 200,
				Msg:  req.Param.Msg,
			},
		}, err
	}
	return &pb.CreateInfoReply{
		Data: &pb.Resp{
			Code: 200,
			Msg:  query.Name,
		},
	}, nil
}
func (s *InfoService) UpdateInfo(ctx context.Context, req *pb.UpdateInfoRequest) (*pb.UpdateInfoReply, error) {
	info := biz.Info{Name: "zzz"}
	redis, err := s.uc.TestWRedis(ctx, &info)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateInfoReply{
		Data: &pb.Resp{
			Code: 200,
			Msg:  redis.Name,
		},
	}, nil
}
func (s *InfoService) DeleteInfo(ctx context.Context, req *pb.DeleteInfoRequest) (*pb.DeleteInfoReply, error) {
	panic("zz")
	return &pb.DeleteInfoReply{
		Data: &pb.Resp{
			Code: 200,
			Msg:  req.Param.Msg,
		},
	}, nil
}
func (s *InfoService) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoReply, error) {
	info := biz.Info{Name: "zzz"}
	redis, err := s.uc.TestRRedis(ctx, &info)
	if err != nil {
		return nil, err
	}

	return &pb.GetInfoReply{
		Data: &pb.Resp{
			Code: 200,
			Msg:  redis.Name,
		},
	}, nil
}
func (s *InfoService) ListInfo(ctx context.Context, req *pb.ListInfoRequest) (*pb.ListInfoReply, error) {
	return &pb.ListInfoReply{
		Data: &pb.Resp{
			Code: 200,
			Msg:  req.Param.Msg,
		},
	}, nil
}
