package service

import (
	"bysos/app/user/internal/biz"
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	pb "bysos/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{
		Code: 200,
		Msg:  req.Name,
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{
		Code: 200,
		Msg:  req.Name,
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("taskkill", "/F", "/T", "/PID", fmt.Sprintf("%d", os.Getpid()))
	default:
		cmd = exec.Command("kill", "-9", fmt.Sprintf("%d", os.Getpid()))
	}
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error sending kill signal:", err)
	}

	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{
		Code: 200,
		Msg:  "get_user",
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{
		Code: 200,
		Msg:  req.Msg,
	}, nil
}
