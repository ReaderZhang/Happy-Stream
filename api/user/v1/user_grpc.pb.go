// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/user/v1/user.proto

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterReply, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutReply, error)
	//==========================================================
	MemberAdd(ctx context.Context, in *MemberAddReq, opts ...grpc.CallOption) (*MemberAddReply, error)
	MemberRemove(ctx context.Context, in *MemberRemoveReq, opts ...grpc.CallOption) (*MemberRemoveReply, error)
	MemberChange(ctx context.Context, in *MemberChangeReq, opts ...grpc.CallOption) (*MemberChangeReply, error)
	MemberGet(ctx context.Context, in *MemberGetReq, opts ...grpc.CallOption) (*MemberGetReply, error)
	MemberGetBetweenTime(ctx context.Context, in *MemberGetBetweenTimeReq, opts ...grpc.CallOption) (*MemberGetBetweenTimeReply, error)
	//==================================================
	RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddReply, error)
	RoleRemove(ctx context.Context, in *RoleRemoveReq, opts ...grpc.CallOption) (*RoleRemoveReply, error)
	RoleChange(ctx context.Context, in *RoleChangeReq, opts ...grpc.CallOption) (*RoleChangeReply, error)
	RoleGet(ctx context.Context, in *RoleGetReq, opts ...grpc.CallOption) (*RoleGetReply, error)
	//================================================================
	PermissionAdd(ctx context.Context, in *PermissionAddReq, opts ...grpc.CallOption) (*PermissionAddReply, error)
	PermissionRemove(ctx context.Context, in *PermissionRemoveReq, opts ...grpc.CallOption) (*PermissionRemoveReply, error)
	PermissionChange(ctx context.Context, in *PermissionChangeReq, opts ...grpc.CallOption) (*PermissionChangeReply, error)
	PermissionGet(ctx context.Context, in *PermissionGetReq, opts ...grpc.CallOption) (*PermissionGetReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) MemberAdd(ctx context.Context, in *MemberAddReq, opts ...grpc.CallOption) (*MemberAddReply, error) {
	out := new(MemberAddReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/MemberAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) MemberRemove(ctx context.Context, in *MemberRemoveReq, opts ...grpc.CallOption) (*MemberRemoveReply, error) {
	out := new(MemberRemoveReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/MemberRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) MemberChange(ctx context.Context, in *MemberChangeReq, opts ...grpc.CallOption) (*MemberChangeReply, error) {
	out := new(MemberChangeReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/MemberChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) MemberGet(ctx context.Context, in *MemberGetReq, opts ...grpc.CallOption) (*MemberGetReply, error) {
	out := new(MemberGetReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/MemberGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) MemberGetBetweenTime(ctx context.Context, in *MemberGetBetweenTimeReq, opts ...grpc.CallOption) (*MemberGetBetweenTimeReply, error) {
	out := new(MemberGetBetweenTimeReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/MemberGetBetweenTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddReply, error) {
	out := new(RoleAddReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/RoleAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RoleRemove(ctx context.Context, in *RoleRemoveReq, opts ...grpc.CallOption) (*RoleRemoveReply, error) {
	out := new(RoleRemoveReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/RoleRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RoleChange(ctx context.Context, in *RoleChangeReq, opts ...grpc.CallOption) (*RoleChangeReply, error) {
	out := new(RoleChangeReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/RoleChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RoleGet(ctx context.Context, in *RoleGetReq, opts ...grpc.CallOption) (*RoleGetReply, error) {
	out := new(RoleGetReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/RoleGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) PermissionAdd(ctx context.Context, in *PermissionAddReq, opts ...grpc.CallOption) (*PermissionAddReply, error) {
	out := new(PermissionAddReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/PermissionAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) PermissionRemove(ctx context.Context, in *PermissionRemoveReq, opts ...grpc.CallOption) (*PermissionRemoveReply, error) {
	out := new(PermissionRemoveReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/PermissionRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) PermissionChange(ctx context.Context, in *PermissionChangeReq, opts ...grpc.CallOption) (*PermissionChangeReply, error) {
	out := new(PermissionChangeReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/PermissionChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) PermissionGet(ctx context.Context, in *PermissionGetReq, opts ...grpc.CallOption) (*PermissionGetReply, error) {
	out := new(PermissionGetReply)
	err := c.cc.Invoke(ctx, "/user.v1.User/PermissionGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	//==========================================================
	MemberAdd(context.Context, *MemberAddReq) (*MemberAddReply, error)
	MemberRemove(context.Context, *MemberRemoveReq) (*MemberRemoveReply, error)
	MemberChange(context.Context, *MemberChangeReq) (*MemberChangeReply, error)
	MemberGet(context.Context, *MemberGetReq) (*MemberGetReply, error)
	MemberGetBetweenTime(context.Context, *MemberGetBetweenTimeReq) (*MemberGetBetweenTimeReply, error)
	//==================================================
	RoleAdd(context.Context, *RoleAddReq) (*RoleAddReply, error)
	RoleRemove(context.Context, *RoleRemoveReq) (*RoleRemoveReply, error)
	RoleChange(context.Context, *RoleChangeReq) (*RoleChangeReply, error)
	RoleGet(context.Context, *RoleGetReq) (*RoleGetReply, error)
	//================================================================
	PermissionAdd(context.Context, *PermissionAddReq) (*PermissionAddReply, error)
	PermissionRemove(context.Context, *PermissionRemoveReq) (*PermissionRemoveReply, error)
	PermissionChange(context.Context, *PermissionChangeReq) (*PermissionChangeReply, error)
	PermissionGet(context.Context, *PermissionGetReq) (*PermissionGetReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Login(context.Context, *LoginReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) Register(context.Context, *RegisterReq) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) Logout(context.Context, *LogoutReq) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServer) MemberAdd(context.Context, *MemberAddReq) (*MemberAddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberAdd not implemented")
}
func (UnimplementedUserServer) MemberRemove(context.Context, *MemberRemoveReq) (*MemberRemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberRemove not implemented")
}
func (UnimplementedUserServer) MemberChange(context.Context, *MemberChangeReq) (*MemberChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberChange not implemented")
}
func (UnimplementedUserServer) MemberGet(context.Context, *MemberGetReq) (*MemberGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberGet not implemented")
}
func (UnimplementedUserServer) MemberGetBetweenTime(context.Context, *MemberGetBetweenTimeReq) (*MemberGetBetweenTimeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberGetBetweenTime not implemented")
}
func (UnimplementedUserServer) RoleAdd(context.Context, *RoleAddReq) (*RoleAddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleAdd not implemented")
}
func (UnimplementedUserServer) RoleRemove(context.Context, *RoleRemoveReq) (*RoleRemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleRemove not implemented")
}
func (UnimplementedUserServer) RoleChange(context.Context, *RoleChangeReq) (*RoleChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleChange not implemented")
}
func (UnimplementedUserServer) RoleGet(context.Context, *RoleGetReq) (*RoleGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleGet not implemented")
}
func (UnimplementedUserServer) PermissionAdd(context.Context, *PermissionAddReq) (*PermissionAddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionAdd not implemented")
}
func (UnimplementedUserServer) PermissionRemove(context.Context, *PermissionRemoveReq) (*PermissionRemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionRemove not implemented")
}
func (UnimplementedUserServer) PermissionChange(context.Context, *PermissionChangeReq) (*PermissionChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionChange not implemented")
}
func (UnimplementedUserServer) PermissionGet(context.Context, *PermissionGetReq) (*PermissionGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionGet not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_MemberAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).MemberAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/MemberAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).MemberAdd(ctx, req.(*MemberAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_MemberRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRemoveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).MemberRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/MemberRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).MemberRemove(ctx, req.(*MemberRemoveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_MemberChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberChangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).MemberChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/MemberChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).MemberChange(ctx, req.(*MemberChangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_MemberGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).MemberGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/MemberGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).MemberGet(ctx, req.(*MemberGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_MemberGetBetweenTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberGetBetweenTimeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).MemberGetBetweenTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/MemberGetBetweenTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).MemberGetBetweenTime(ctx, req.(*MemberGetBetweenTimeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RoleAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RoleAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/RoleAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RoleAdd(ctx, req.(*RoleAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RoleRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleRemoveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RoleRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/RoleRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RoleRemove(ctx, req.(*RoleRemoveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RoleChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleChangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RoleChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/RoleChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RoleChange(ctx, req.(*RoleChangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RoleGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RoleGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/RoleGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RoleGet(ctx, req.(*RoleGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_PermissionAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).PermissionAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/PermissionAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).PermissionAdd(ctx, req.(*PermissionAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_PermissionRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionRemoveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).PermissionRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/PermissionRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).PermissionRemove(ctx, req.(*PermissionRemoveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_PermissionChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionChangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).PermissionChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/PermissionChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).PermissionChange(ctx, req.(*PermissionChangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_PermissionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).PermissionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.v1.User/PermissionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).PermissionGet(ctx, req.(*PermissionGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _User_Logout_Handler,
		},
		{
			MethodName: "MemberAdd",
			Handler:    _User_MemberAdd_Handler,
		},
		{
			MethodName: "MemberRemove",
			Handler:    _User_MemberRemove_Handler,
		},
		{
			MethodName: "MemberChange",
			Handler:    _User_MemberChange_Handler,
		},
		{
			MethodName: "MemberGet",
			Handler:    _User_MemberGet_Handler,
		},
		{
			MethodName: "MemberGetBetweenTime",
			Handler:    _User_MemberGetBetweenTime_Handler,
		},
		{
			MethodName: "RoleAdd",
			Handler:    _User_RoleAdd_Handler,
		},
		{
			MethodName: "RoleRemove",
			Handler:    _User_RoleRemove_Handler,
		},
		{
			MethodName: "RoleChange",
			Handler:    _User_RoleChange_Handler,
		},
		{
			MethodName: "RoleGet",
			Handler:    _User_RoleGet_Handler,
		},
		{
			MethodName: "PermissionAdd",
			Handler:    _User_PermissionAdd_Handler,
		},
		{
			MethodName: "PermissionRemove",
			Handler:    _User_PermissionRemove_Handler,
		},
		{
			MethodName: "PermissionChange",
			Handler:    _User_PermissionChange_Handler,
		},
		{
			MethodName: "PermissionGet",
			Handler:    _User_PermissionGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/user.proto",
}