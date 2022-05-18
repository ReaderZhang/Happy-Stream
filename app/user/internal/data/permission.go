package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/qqz/Happy-Stream/app/user/internal/biz"
	"github.com/qqz/Happy-Stream/app/user/internal/data/model"
	"github.com/qqz/Happy-Stream/globalUtils"
	"strconv"
)

type permissionRepo struct {
	data *Data
	log  *log.Helper
}

func (pr *permissionRepo) CreatePermission(ctx context.Context, url string, parentId int64) error {
	p := &model.Permission{
		URL:      url,
		ParentID: strconv.FormatInt(parentId, 10),
	}
	result := pr.data.db.Create(&p)
	return result.Error
}

func (pr *permissionRepo) UpdatePermission(ctx context.Context, url string, parentId int64) error {
	p := &model.Permission{
		URL:      url,
		ParentID: strconv.FormatInt(parentId, 10),
	}
	result := pr.data.db.Save(&p)
	return result.Error
}

func (pr *permissionRepo) DeletePermission(ctx context.Context, id int64) error {
	p := &model.Permission{}
	result := pr.data.db.Delete(&p, id)
	return result.Error
}

func (pr *permissionRepo) QueryPermission(ctx context.Context, id int64) (*biz.Permission, error) {
	p := &model.Permission{}
	result := pr.data.db.Find(&p, id)
	if result.Error != nil {
		return nil, result.Error
	}
	permission := &biz.Permission{
		Id:  id,
		Url: p.URL,
	}
	pr.getChildren(ctx, permission, p.ParentID)
	return permission, nil
}

func (pr *permissionRepo) ListQueryPermission(ctx context.Context, pageNum int64, pageSize int64) (*biz.ListPermission, error) {
	var count int64
	pr.data.db.Count(&count)
	permissions := make([]model.Permission, pageSize)
	pers := make([]biz.Permission, pageSize)
	pr.data.db.Limit(int(pageSize)).Offset(int(pageNum-1) * int(pageNum)).Find(&permissions)
	copier.Copy(&pers, &permissions)
	for i := range permissions {
		pr.getChildren(ctx, &pers[i], permissions[i].ParentID)
	}
	return &biz.ListPermission{
		Permissions: pers,
		Pageination: globalUtils.Pageination{
			PageNum:  pageNum,
			PageSize: pageSize,
			Total:    count,
		},
	}, nil
}

func NewPermissionRepo(data *Data, logger log.Logger) biz.PermissionRepo {
	return &permissionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (pr *permissionRepo) getChildren(ctx context.Context, permission *biz.Permission, parentId string) {
	pid, err := strconv.ParseInt(parentId, 10, 64)
	if err != nil {
		return
	}
	if pid == 0 {
		return
	} else {
		children := make([]model.Permission, 10)
		result := pr.data.db.Where("parent_id = ?", parentId).Find(&children)
		if result.Error != nil {
			fmt.Println(result.Error)
			return
		}
		newc := make([]biz.Permission, 10)
		for i := range children {
			c := &biz.Permission{
				Id:  int64(children[i].ID),
				Url: children[i].URL,
			}
			newc = append(newc, *c)
		}
		permission.Children = newc
		for i := range newc {
			permission = &newc[i]
			pr.getChildren(ctx, permission, strconv.FormatInt(permission.Children[i].Id, 10))
		}
	}
}
