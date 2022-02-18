package service

import (
	"context"

	pb "github.com/qqz/Happy-Stream/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *UserService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *UserService) MemberAdd(ctx context.Context, req *pb.MemberAddReq) (*pb.MemberAddReply, error) {
	return &pb.MemberAddReply{}, nil
}
func (s *UserService) MemberRemove(ctx context.Context, req *pb.MemberRemoveReq) (*pb.MemberRemoveReply, error) {
	return &pb.MemberRemoveReply{}, nil
}
func (s *UserService) MemberChange(ctx context.Context, req *pb.MemberChangeReq) (*pb.MemberChangeReply, error) {
	return &pb.MemberChangeReply{}, nil
}
func (s *UserService) MemberGet(ctx context.Context, req *pb.MemberGetReq) (*pb.MemberGetReply, error) {
	return &pb.MemberGetReply{}, nil
}
func (s *UserService) MemberGetBetweenTime(ctx context.Context, req *pb.MemberGetBetweenTimeReq) (*pb.MemberGetBetweenTimeReply, error) {
	return &pb.MemberGetBetweenTimeReply{}, nil
}
func (s *UserService) RoleAdd(ctx context.Context, req *pb.RoleAddReq) (*pb.RoleAddReply, error) {
	return &pb.RoleAddReply{}, nil
}
func (s *UserService) RoleRemove(ctx context.Context, req *pb.RoleRemoveReq) (*pb.RoleRemoveReply, error) {
	return &pb.RoleRemoveReply{}, nil
}
func (s *UserService) RoleChange(ctx context.Context, req *pb.RoleChangeReq) (*pb.RoleChangeReply, error) {
	return &pb.RoleChangeReply{}, nil
}
func (s *UserService) RoleGet(ctx context.Context, req *pb.RoleGetReq) (*pb.RoleGetReply, error) {
	return &pb.RoleGetReply{}, nil
}
func (s *UserService) PermissionAdd(ctx context.Context, req *pb.PermissionAddReq) (*pb.PermissionAddReply, error) {
	return &pb.PermissionAddReply{}, nil
}
func (s *UserService) PermissionRemove(ctx context.Context, req *pb.PermissionRemoveReq) (*pb.PermissionRemoveReply, error) {
	return &pb.PermissionRemoveReply{}, nil
}
func (s *UserService) PermissionChange(ctx context.Context, req *pb.PermissionChangeReq) (*pb.PermissionChangeReply, error) {
	return &pb.PermissionChangeReply{}, nil
}
func (s *UserService) PermissionGet(ctx context.Context, req *pb.PermissionGetReq) (*pb.PermissionGetReply, error) {
	return &pb.PermissionGetReply{}, nil
}
