package service

import (
	"context"
	"fmt"
	"github.com/qqz/Happy-Stream/app/order/internal/biz"
	"github.com/qqz/Happy-Stream/globalUtils"
	"strconv"

	pb "github.com/qqz/Happy-Stream/api/order/v1"
)

func (s *OrderService) AddOrder(ctx context.Context, req *pb.OrderAddReq) (*pb.OrderAddReply, error) {
	userid, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}
	err = s.oc.Add(ctx, &biz.Order{
		Text:   req.Text,
		UserId: int32(userid),
	})
	return &pb.OrderAddReply{}, err
}
func (s *OrderService) RemoveOrder(ctx context.Context, req *pb.OrderRemoveReq) (*pb.OrderRemoveReply, error) {
	err := s.oc.Remove(ctx, req.Id)
	return &pb.OrderRemoveReply{}, err
}
func (s *OrderService) ChangeOrder(ctx context.Context, req *pb.OrderChangeReq) (*pb.OrderChangeReply, error) {
	err := s.oc.Change(ctx, &biz.Order{
		Id:     req.Id,
		Text:   req.Text,
		Status: req.Status,
		UserId: req.UserId,
	})
	return &pb.OrderChangeReply{}, err
}
func (s *OrderService) ListOrder(ctx context.Context, req *pb.ListOrderQueryReq) (*pb.ListOrderQueryReply, error) {
	page := &globalUtils.Pageination{PageNum: req.Page.PageNum, PageSize: req.Page.PageSize}
	list, err := s.oc.List(ctx, req.QueryType, req.QueryParam, *page)
	fmt.Println("pageSize", page.PageSize, " == pageNum", page.PageNum, " == total", page.Total)
	orders := list.Orders
	infos := make([]*pb.OrderInfo, 0)
	fmt.Println("len ", len(orders))
	for i := range orders {
		p := &pb.OrderInfo{
			Id:          orders[i].Id,
			Text:        orders[i].Text,
			CreatedTime: orders[i].CreateTime,
			Status:      orders[i].Status,
		}
		infos = append(infos, p)
	}
	fmt.Print("service count", list.Page.Total)
	return &pb.ListOrderQueryReply{Orders: infos, Page: &pb.PageInfoReply{
		PageNum:  page.PageNum,
		PageSize: page.PageSize,
		Total:    list.Page.Total,
	}}, err
}
func (s *OrderService) QueryOrder(ctx context.Context, req *pb.OrderQueryReq) (*pb.OrderQueryReply, error) {
	result, err := s.oc.Get(ctx, req.Id)
	if result == nil {
		return nil, err
	}
	return &pb.OrderQueryReply{Order: &pb.OrderInfo{
		Id:          req.Id,
		Text:        result.Text,
		CreatedTime: result.CreateTime,
		Status:      result.Status,
	}}, err
}
