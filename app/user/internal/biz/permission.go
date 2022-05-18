package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/qqz/Happy-Stream/globalUtils"
)

type Permission struct {
	Id       int64
	Url      string
	Children []Permission
}

type ListPermission struct {
	Permissions []Permission
	globalUtils.Pageination
}

type PermissionRepo interface {
	CreatePermission(ctx context.Context, url string, parentId int64) error
	UpdatePermission(ctx context.Context, url string, parentId int64) error
	DeletePermission(ctx context.Context, id int64) error

	QueryPermission(ctx context.Context, id int64) (*Permission, error)
	ListQueryPermission(ctx context.Context, pageNum int64, pageSize int64) (*ListPermission, error)
}

type PermissionUseCase struct {
	repo PermissionRepo
	log  *log.Helper
}

func NewPermissionUseCase(repo PermissionRepo, logger log.Logger) *PermissionUseCase {
	return &PermissionUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (pc *PermissionUseCase) Save(ctx context.Context, url string, parentId int64) (err error) {
	err = pc.repo.CreatePermission(ctx, url, parentId)
	return err
}

func (pc *PermissionUseCase) Update(ctx context.Context, url string, parentId int64) (err error) {
	err = pc.repo.UpdatePermission(ctx, url, parentId)
	return err
}

func (pc *PermissionUseCase) Delete(ctx context.Context, id int64) (err error) {
	err = pc.repo.DeletePermission(ctx, id)
	return err
}

func (pc *PermissionUseCase) Query(ctx context.Context, id int64) (permissions *Permission, err error) {
	permissions, err = pc.repo.QueryPermission(ctx, id)
	return permissions, err
}

func (pc *PermissionUseCase) List(ctx context.Context, pageInfo globalUtils.Pageination) (permissions *ListPermission, err error) {
	permissions, err = pc.repo.ListQueryPermission(ctx, pageInfo.PageNum, pageInfo.PageSize)
	return permissions, err
}
