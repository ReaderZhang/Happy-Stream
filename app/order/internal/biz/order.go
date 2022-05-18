package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/qqz/Happy-Stream/globalUtils"
)

type Order struct {
	Id         int64
	Text       string
	Status     int32
	UserId     int32
	CreateTime string
}

type OrderListResult struct {
	Orders []*Order
	Page   *globalUtils.Pageination
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, order *Order) error
	ListQuery(ctx context.Context, pageSize int64, pageNum int64, paramType string, params string) (*OrderListResult, error)
	DeleteOrder(ctx context.Context, id int64) error
	UpdateOrder(ctx context.Context, order *Order) error
	SelectOrder(ctx context.Context, id int64) (*Order, error)
}

type OrderUseCase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUseCase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (oc *OrderUseCase) Get(ctx context.Context, id int64) (*Order, error) {
	return oc.repo.SelectOrder(ctx, id)
}

func (oc *OrderUseCase) Remove(ctx context.Context, id int64) error {
	return oc.repo.DeleteOrder(ctx, id)
}

func (oc *OrderUseCase) Add(ctx context.Context, order *Order) error {
	return oc.repo.CreateOrder(ctx, order)
}

func (oc *OrderUseCase) Change(ctx context.Context, order *Order) error {
	return oc.repo.UpdateOrder(ctx, order)
}

func (oc *OrderUseCase) List(ctx context.Context, paramType string, params string, pageination globalUtils.Pageination) (*OrderListResult, error) {
	return oc.repo.ListQuery(ctx, pageination.PageSize, pageination.PageNum, paramType, params)
}
