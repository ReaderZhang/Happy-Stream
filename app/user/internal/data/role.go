package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/qqz/Happy-Stream/app/user/internal/biz"
	"github.com/qqz/Happy-Stream/app/user/internal/data/model"
	"github.com/qqz/Happy-Stream/globalUtils"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

func (rr *roleRepo) CreateRole(ctx context.Context, role string, pid int64) error {
	r := &model.Role{
		RoleName: role,
		Pid:      int32(pid),
	}
	result := rr.data.db.Create(&r)
	return result.Error
}

func (rr *roleRepo) UpdateRole(ctx context.Context, role string, pid int64) error {
	r := &model.Role{
		RoleName: role,
		Pid:      int32(pid),
	}
	result := rr.data.db.Updates(&r)
	return result.Error
}

func (rr *roleRepo) DeleteRole(ctx context.Context, id int64) error {
	r := &model.Role{}
	result := rr.data.db.Delete(&r, id)
	return result.Error
}

func (rr *roleRepo) QueryRole(ctx context.Context, id int64) (*biz.Role, error) {
	r := &model.Role{}
	result := rr.data.db.First(&r, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Role{
		Id:   id,
		Name: r.RoleName,
	}, nil
}

func (rr *roleRepo) ListQueryRole(ctx context.Context, pageNum int64, pageSize int64) (*biz.ListRoleResult, error) {
	roles := make([]*biz.Role, pageNum)
	result := rr.data.db.Limit(int(pageSize)).Offset(int(pageNum-1) * int(pageNum)).Find(roles)
	if result.Error != nil {
		return nil, result.Error
	}
	var count int64
	rr.data.db.Count(&count)
	return &biz.ListRoleResult{
		Roles: roles,
		Pageination: globalUtils.Pageination{
			PageNum:  pageNum,
			PageSize: pageSize,
			Total:    count,
		},
	}, nil
}

func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
