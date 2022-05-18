package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/qqz/Happy-Stream/app/order/internal/biz"
	"github.com/qqz/Happy-Stream/app/order/internal/data/model"
	"github.com/qqz/Happy-Stream/globalUtils"
	"strings"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

func (or *orderRepo) CreateOrder(ctx context.Context, order *biz.Order) error {
	o := &model.Order{
		Text:   order.Text,
		UserId: order.UserId,
	}

	result := or.data.db.Create(&o)
	return result.Error
}

func (or *orderRepo) ListQuery(ctx context.Context, pageSize int64, pageNum int64, paramType string, params string) (*biz.OrderListResult, error) {
	borders := make([]*biz.Order, 0)
	orders := make([]*model.Order, 0)

	pars := strings.Split(params, ",")
	var count int64
	page := &globalUtils.Pageination{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    count,
	}
	switch paramType {
	case "time":
		{
			start, end := pars[0], pars[1]
			result := or.data.db.Where("created_at BETWEEN ? AND ?", start, end).Scopes(globalUtils.Paginate(pageNum, pageSize)).Find(&orders)
			if result.Error != nil {
				return nil, result.Error
			}
			result = or.data.db.Model(&model.Order{}).Where("created_at between ? and ?", start, end).Count(&count)
			if result.Error != nil {
				return nil, result.Error
			}
			fmt.Println("count ", count)
			page.Total = count
		}
	case "ids":
		{
			var count int64
			result := or.data.db.Where("id IN ?", pars).Scopes(globalUtils.Paginate(pageNum, pageSize)).Find(&orders)
			if result.Error != nil {
				return nil, result.Error
			}
			result = or.data.db.Model(&model.Order{}).Where("id IN ?", pars).Count(&count)
			if result.Error != nil {
				return nil, result.Error
			}
			fmt.Println("count ", count)
			page.Total = count
		}
	case "text":
		{
			var count int64
			result := or.data.db.Where("text IN ?", pars).Scopes(globalUtils.Paginate(pageNum, pageSize)).Find(&orders)
			if result.Error != nil {
				return nil, result.Error
			}
			result = or.data.db.Model(&model.Order{}).Where("text IN ?", pars).Count(&count)
			if result.Error != nil {
				return nil, result.Error
			}
			fmt.Println("count ", count)
			page.Total = count
		}
	case "user_id":
		{
			result := or.data.db.Where("user_id IN ?", pars).Scopes(globalUtils.Paginate(pageNum, pageSize)).Find(&orders)
			if result.Error != nil {
				return nil, result.Error
			}
			result = or.data.db.Model(&model.Order{}).Where("user_id IN ?", pars).Count(&count)
			if result.Error != nil {
				return nil, result.Error
			}
			fmt.Println("count ", count)
			page.Total = count
		}
	}
	fmt.Println("len is is is ", len(borders))
	for i := range orders {
		format := orders[i].CreatedAt.Format("2006-01-02 15:04:05")
		borders = append(borders, &biz.Order{
			Id:         orders[i].ID,
			Text:       orders[i].Text,
			Status:     orders[i].Status,
			UserId:     orders[i].UserId,
			CreateTime: format,
		})
	}
	fmt.Println("last  count ", count)
	return &biz.OrderListResult{
		Orders: borders,
		Page:   page,
	}, nil
}

func (or *orderRepo) DeleteOrder(ctx context.Context, id int64) error {
	result := or.data.db.Model(&model.Order{}).Where("id = ?", id).Update("status", 0)
	return result.Error
}

func (or *orderRepo) UpdateOrder(ctx context.Context, order *biz.Order) error {
	o := &model.Order{
		ID: order.Id,
	}
	result := or.data.db.Model(&o).Omit("created_at").Updates(model.Order{
		Text:   order.Text,
		Status: order.Status,
		UserId: order.UserId,
	})
	return result.Error
}

func (or *orderRepo) SelectOrder(ctx context.Context, id int64) (*biz.Order, error) {
	o := &model.Order{
		ID: id,
	}
	result := or.data.db.First(&o)
	if result.Error != nil {
		return nil, result.Error
	}
	bo := &biz.Order{}
	copier.Copy(&bo, &o)
	return bo, nil
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
