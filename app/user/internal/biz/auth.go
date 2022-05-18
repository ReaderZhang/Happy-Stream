package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthMember struct {
	Username string
	Password string
	Phone    string
}

type AuthResult struct {
	Token    string
	Avator   string
	Username string
	Role     *Role
}

type AuthRepo interface {
	Login(ctx context.Context, am *AuthMember) (*AuthResult, error)
	Register(ctx context.Context, am *AuthMember) (*AuthResult, error)
	Logout(ctx context.Context) error
}

type AuthUseCase struct {
	repo AuthRepo
	log  *log.Helper

	roleUc   *RoleUseCase
	memberUc *MemberUseCase
}

func NewAuthUseCase(repo AuthRepo, logger log.Logger, roleUc *RoleUseCase, memberUc *MemberUseCase) *AuthUseCase {
	return &AuthUseCase{
		repo:     repo,
		log:      log.NewHelper(logger),
		roleUc:   roleUc,
		memberUc: memberUc,
	}
}

func (ac *AuthUseCase) Login(ctx context.Context, am *AuthMember) (*AuthResult, error) {
	a := &AuthMember{
		Username: am.Username,
		Password: am.Password,
		Phone:    am.Phone,
	}
	return ac.repo.Login(ctx, a)
}

func (ac *AuthUseCase) Register(ctx context.Context, am *AuthMember) (*AuthResult, error) {
	return ac.repo.Register(ctx, am)
}

func (as *AuthUseCase) Logout(ctx context.Context) error {
	return nil
}
