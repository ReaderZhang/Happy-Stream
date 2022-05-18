package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/qqz/Happy-Stream/api/order/v1"
	"github.com/qqz/Happy-Stream/app/order/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewOrderService)

type OrderService struct {
	pb.UnimplementedOrderServer

	oc  *biz.OrderUseCase
	log *log.Helper
}

func NewOrderService(oc *biz.OrderUseCase, logger log.Logger) *OrderService {
	return &OrderService{
		oc:  oc,
		log: log.NewHelper(logger),
	}
}
