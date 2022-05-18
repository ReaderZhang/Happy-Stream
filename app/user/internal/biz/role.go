package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/qqz/Happy-Stream/globalUtils"
)

type Role struct {
	Id         int64
	Name       string
	pid        int64
	Permission Permission
}

type ListRoleResult struct {
	Roles []*Role
	globalUtils.Pageination
}

type RoleRepo interface {
	CreateRole(ctx context.Context, role string, pid int64) error
	UpdateRole(ctx context.Context, role string, pid int64) error
	DeleteRole(ctx context.Context, id int64) error

	QueryRole(ctx context.Context, id int64) (*Role, error)
	ListQueryRole(ctx context.Context, pageNum int64, pageSize int64) (*ListRoleResult, error)
}

type RoleUseCase struct {
	repo RoleRepo
	log  *log.Helper

	permissionUc *PermissionUseCase
}

func NewRoleUseCase(repo RoleRepo, logger log.Logger) *RoleUseCase {
	return &RoleUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (rc *RoleUseCase) Save(ctx context.Context, role string, pid int64) (err error) {
	err = rc.repo.CreateRole(ctx, role, pid)
	return err
}

func (rc *RoleUseCase) Update(ctx context.Context, role string, pid int64) (err error) {
	err = rc.repo.UpdateRole(ctx, role, pid)
	return err
}

func (rc *RoleUseCase) Delete(ctx context.Context, id int64) (err error) {
	err = rc.repo.DeleteRole(ctx, id)
	return err
}

func (rc *RoleUseCase) Query(ctx context.Context, id int64) (role *Role, err error) {
	role, err = rc.repo.QueryRole(ctx, id)
	if err != nil {
		return nil, err
	}
	permission, err2 := rc.permissionUc.repo.QueryPermission(ctx, role.pid)
	if err2 != nil {
		return nil, err2
	}
	role.Permission = *permission
	return role, err
}

func (rc *RoleUseCase) List(ctx context.Context, pageInfo globalUtils.Pageination) (roles *ListRoleResult, err error) {
	roles, err = rc.repo.ListQueryRole(ctx, pageInfo.PageNum, pageInfo.PageSize)
	return roles, err
}
