package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user/internal/biz"

	pb "user/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserInfo) (*pb.UserInfoResponse, error) {

	user, err := s.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UserInfoResponse{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Birthday: user.Birthday,
	}, nil
}

func (s *UserService) GetUserList(ctx context.Context, req *pb.PageInfo) (*pb.UserListResponse, error) {

	list, count, err := s.uc.List(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		return nil, err
	}

	resp := &pb.UserListResponse{Count: int32(count)}
	for _, item := range list {
		resp.Data = append(resp.Data, UserResponse(item))
	}

	return resp, nil
}

func UserResponse(user *biz.User) *pb.UserInfoResponse {
	return &pb.UserInfoResponse{
		Id:       user.ID,
		NickName: user.NickName,
		Mobile:   user.Mobile,
		Birthday: user.Birthday,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
}

func (s *UserService) GetUserByMobile(ctx context.Context, req *pb.MobileRequest) (*pb.UserInfoResponse, error) {

	user, err := s.uc.UserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	return UserResponse(user), nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.IdRequest) (*pb.UserInfoResponse, error) {

	user, err := s.uc.UserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return UserResponse(user), nil

}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*emptypb.Empty, error) {

	var user *biz.User
	user = &biz.User{
		ID:       req.Id,
		NickName: req.NickName,
		Gender:   req.Gender,
		Birthday: req.Birthday,
	}

	ok, err := s.uc.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, status.Errorf(codes.Internal, "内部错误")
	}

	return &emptypb.Empty{}, nil

}
func (s *UserService) CheckPassword(ctx context.Context, req *pb.PasswordCheckRequest) (*pb.CheckPasswordResponse, error) {

	ok, err := s.uc.CheckPassword(ctx, req.Password, req.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return &pb.CheckPasswordResponse{
		Success: ok,
	}, nil
}
