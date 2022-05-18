package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/qqz/Happy-Stream/api/user/v1"
	"github.com/qqz/Happy-Stream/app/user/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	pb.UnimplementedUserServer

	ac *biz.AuthUseCase
	mc *biz.MemberUseCase

	rc *biz.RoleUseCase
	pc *biz.PermissionUseCase

	log *log.Helper
}

func NewUserService(ac *biz.AuthUseCase, mc *biz.MemberUseCase, pc *biz.PermissionUseCase, rc *biz.RoleUseCase, logger log.Logger) *UserService {
	return &UserService{
		ac:  ac,
		mc:  mc,
		pc:  pc,
		rc:  rc,
		log: log.NewHelper(logger),
	}
}
