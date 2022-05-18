// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/order/v1/order.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	AddOrder(ctx context.Context, in *OrderAddReq, opts ...grpc.CallOption) (*OrderAddReply, error)
	RemoveOrder(ctx context.Context, in *OrderRemoveReq, opts ...grpc.CallOption) (*OrderRemoveReply, error)
	ChangeOrder(ctx context.Context, in *OrderChangeReq, opts ...grpc.CallOption) (*OrderChangeReply, error)
	ListOrder(ctx context.Context, in *ListOrderQueryReq, opts ...grpc.CallOption) (*ListOrderQueryReply, error)
	QueryOrder(ctx context.Context, in *OrderQueryReq, opts ...grpc.CallOption) (*OrderQueryReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) AddOrder(ctx context.Context, in *OrderAddReq, opts ...grpc.CallOption) (*OrderAddReply, error) {
	out := new(OrderAddReply)
	err := c.cc.Invoke(ctx, "/order.v1.Order/AddOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) RemoveOrder(ctx context.Context, in *OrderRemoveReq, opts ...grpc.CallOption) (*OrderRemoveReply, error) {
	out := new(OrderRemoveReply)
	err := c.cc.Invoke(ctx, "/order.v1.Order/RemoveOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) ChangeOrder(ctx context.Context, in *OrderChangeReq, opts ...grpc.CallOption) (*OrderChangeReply, error) {
	out := new(OrderChangeReply)
	err := c.cc.Invoke(ctx, "/order.v1.Order/ChangeOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) ListOrder(ctx context.Context, in *ListOrderQueryReq, opts ...grpc.CallOption) (*ListOrderQueryReply, error) {
	out := new(ListOrderQueryReply)
	err := c.cc.Invoke(ctx, "/order.v1.Order/ListOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) QueryOrder(ctx context.Context, in *OrderQueryReq, opts ...grpc.CallOption) (*OrderQueryReply, error) {
	out := new(OrderQueryReply)
	err := c.cc.Invoke(ctx, "/order.v1.Order/QueryOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	AddOrder(context.Context, *OrderAddReq) (*OrderAddReply, error)
	RemoveOrder(context.Context, *OrderRemoveReq) (*OrderRemoveReply, error)
	ChangeOrder(context.Context, *OrderChangeReq) (*OrderChangeReply, error)
	ListOrder(context.Context, *ListOrderQueryReq) (*ListOrderQueryReply, error)
	QueryOrder(context.Context, *OrderQueryReq) (*OrderQueryReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) AddOrder(context.Context, *OrderAddReq) (*OrderAddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}
func (UnimplementedOrderServer) RemoveOrder(context.Context, *OrderRemoveReq) (*OrderRemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveOrder not implemented")
}
func (UnimplementedOrderServer) ChangeOrder(context.Context, *OrderChangeReq) (*OrderChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeOrder not implemented")
}
func (UnimplementedOrderServer) ListOrder(context.Context, *ListOrderQueryReq) (*ListOrderQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedOrderServer) QueryOrder(context.Context, *OrderQueryReq) (*OrderQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrder not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.v1.Order/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).AddOrder(ctx, req.(*OrderAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_RemoveOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRemoveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).RemoveOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.v1.Order/RemoveOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).RemoveOrder(ctx, req.(*OrderRemoveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_ChangeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderChangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).ChangeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.v1.Order/ChangeOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).ChangeOrder(ctx, req.(*OrderChangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.v1.Order/ListOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).ListOrder(ctx, req.(*ListOrderQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_QueryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).QueryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.v1.Order/QueryOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).QueryOrder(ctx, req.(*OrderQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.v1.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrder",
			Handler:    _Order_AddOrder_Handler,
		},
		{
			MethodName: "RemoveOrder",
			Handler:    _Order_RemoveOrder_Handler,
		},
		{
			MethodName: "ChangeOrder",
			Handler:    _Order_ChangeOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _Order_ListOrder_Handler,
		},
		{
			MethodName: "QueryOrder",
			Handler:    _Order_QueryOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/order/v1/order.proto",
}
