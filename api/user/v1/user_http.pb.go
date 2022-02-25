// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type UserHTTPServer interface {
	ListMemberGetBetweenBalance(context.Context, *ListMemberGetBetweenBalanceReq) (*ListMemberGetBetweenBalanceReply, error)
	ListMemberGetBetweenRecharge(context.Context, *ListMemberGetBetweenRechargeReq) (*ListMemberGetBetweenRechargeReply, error)
	ListMemberGetBetweenTime(context.Context, *ListMemberGetBetweenTimeReq) (*ListMemberGetBetweenTimeReply, error)
	ListMemberGetByEmail(context.Context, *ListMemberGetByEmailReq) (*ListMemberGetByEmailReply, error)
	ListMemberGetById(context.Context, *ListMemberGetByIdReq) (*ListMemberGetByIdReply, error)
	ListMemberGetByName(context.Context, *ListMemberGetByNameReq) (*ListMemberGetByNameReply, error)
	ListMemberGetByPhone(context.Context, *ListMemberGetByPhoneReq) (*ListMemberGetByPhoneReply, error)
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Logout(context.Context, *LogoutReq) (*LogoutReply, error)
	MemberAdd(context.Context, *MemberAddReq) (*MemberAddReply, error)
	MemberChange(context.Context, *MemberChangeReq) (*MemberChangeReply, error)
	MemberGet(context.Context, *MemberGetReq) (*MemberGetReply, error)
	MemberRemove(context.Context, *MemberRemoveReq) (*MemberRemoveReply, error)
	PermissionAdd(context.Context, *PermissionAddReq) (*PermissionAddReply, error)
	PermissionChange(context.Context, *PermissionChangeReq) (*PermissionChangeReply, error)
	PermissionGet(context.Context, *PermissionGetReq) (*PermissionGetReply, error)
	PermissionGetAll(context.Context, *PermissionGetAllReq) (*PermissionGetAllReply, error)
	PermissionRemove(context.Context, *PermissionRemoveReq) (*PermissionRemoveReply, error)
	Register(context.Context, *RegisterReq) (*RegisterReply, error)
	RoleAdd(context.Context, *RoleAddReq) (*RoleAddReply, error)
	RoleChange(context.Context, *RoleChangeReq) (*RoleChangeReply, error)
	RoleGet(context.Context, *RoleGetReq) (*RoleGetReply, error)
	RoleGetAll(context.Context, *RoleGetAllReq) (*RoleGetAllReply, error)
	RoleRemove(context.Context, *RoleRemoveReq) (*RoleRemoveReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/login", _User_Login0_HTTP_Handler(srv))
	r.POST("/v1/register", _User_Register0_HTTP_Handler(srv))
	r.GET("/v1/logout", _User_Logout0_HTTP_Handler(srv))
	r.POST("/v1/member/add", _User_MemberAdd0_HTTP_Handler(srv))
	r.DELETE("/v1/member/remove/{id}", _User_MemberRemove0_HTTP_Handler(srv))
	r.PATCH("/v1/member/update", _User_MemberChange0_HTTP_Handler(srv))
	r.GET("/v1/member/get/{id}", _User_MemberGet0_HTTP_Handler(srv))
	r.GET("/v1/member/list/get/{id}", _User_ListMemberGetById0_HTTP_Handler(srv))
	r.GET("/v1/member/list/get/name/{name}", _User_ListMemberGetByName0_HTTP_Handler(srv))
	r.GET("/v1/member/list/get/phone/{phone}", _User_ListMemberGetByPhone0_HTTP_Handler(srv))
	r.GET("/v1/member/list/get/email/{email}", _User_ListMemberGetByEmail0_HTTP_Handler(srv))
	r.GET("/v1/member/list/get/recharge/{low}/{high}", _User_ListMemberGetBetweenRecharge0_HTTP_Handler(srv))
	r.GET("/v1/member/list/get/balance/{low}/{high}", _User_ListMemberGetBetweenBalance0_HTTP_Handler(srv))
	r.GET("/v1/member/list/get/time/{start}/{end}", _User_ListMemberGetBetweenTime0_HTTP_Handler(srv))
	r.POST("/v1/role/add", _User_RoleAdd0_HTTP_Handler(srv))
	r.DELETE("/v1/role/remove/{id}", _User_RoleRemove0_HTTP_Handler(srv))
	r.PATCH("/v1/role/update", _User_RoleChange0_HTTP_Handler(srv))
	r.GET("/v1/role/get/{id}", _User_RoleGet0_HTTP_Handler(srv))
	r.GET("/v1/role/get/all", _User_RoleGetAll0_HTTP_Handler(srv))
	r.POST("/v1/permission/add", _User_PermissionAdd0_HTTP_Handler(srv))
	r.DELETE("/v1/permission/remove/{id}", _User_PermissionRemove0_HTTP_Handler(srv))
	r.PATCH("/v1/permission/update", _User_PermissionChange0_HTTP_Handler(srv))
	r.GET("/v1/permission/get/{id}", _User_PermissionGet0_HTTP_Handler(srv))
	r.GET("/v1/permission/get/all", _User_PermissionGetAll0_HTTP_Handler(srv))
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _User_Register0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/Register")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _User_Logout0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/Logout")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _User_MemberAdd0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MemberAddReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/MemberAdd")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MemberAdd(ctx, req.(*MemberAddReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MemberAddReply)
		return ctx.Result(200, reply)
	}
}

func _User_MemberRemove0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MemberRemoveReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/MemberRemove")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MemberRemove(ctx, req.(*MemberRemoveReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MemberRemoveReply)
		return ctx.Result(200, reply)
	}
}

func _User_MemberChange0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MemberChangeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/MemberChange")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MemberChange(ctx, req.(*MemberChangeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MemberChangeReply)
		return ctx.Result(200, reply)
	}
}

func _User_MemberGet0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MemberGetReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/MemberGet")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MemberGet(ctx, req.(*MemberGetReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MemberGetReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListMemberGetById0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMemberGetByIdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/ListMemberGetById")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMemberGetById(ctx, req.(*ListMemberGetByIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMemberGetByIdReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListMemberGetByName0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMemberGetByNameReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/ListMemberGetByName")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMemberGetByName(ctx, req.(*ListMemberGetByNameReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMemberGetByNameReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListMemberGetByPhone0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMemberGetByPhoneReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/ListMemberGetByPhone")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMemberGetByPhone(ctx, req.(*ListMemberGetByPhoneReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMemberGetByPhoneReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListMemberGetByEmail0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMemberGetByEmailReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/ListMemberGetByEmail")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMemberGetByEmail(ctx, req.(*ListMemberGetByEmailReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMemberGetByEmailReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListMemberGetBetweenRecharge0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMemberGetBetweenRechargeReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/ListMemberGetBetweenRecharge")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMemberGetBetweenRecharge(ctx, req.(*ListMemberGetBetweenRechargeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMemberGetBetweenRechargeReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListMemberGetBetweenBalance0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMemberGetBetweenBalanceReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/ListMemberGetBetweenBalance")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMemberGetBetweenBalance(ctx, req.(*ListMemberGetBetweenBalanceReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMemberGetBetweenBalanceReply)
		return ctx.Result(200, reply)
	}
}

func _User_ListMemberGetBetweenTime0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMemberGetBetweenTimeReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/ListMemberGetBetweenTime")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMemberGetBetweenTime(ctx, req.(*ListMemberGetBetweenTimeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMemberGetBetweenTimeReply)
		return ctx.Result(200, reply)
	}
}

func _User_RoleAdd0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleAddReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/RoleAdd")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleAdd(ctx, req.(*RoleAddReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleAddReply)
		return ctx.Result(200, reply)
	}
}

func _User_RoleRemove0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleRemoveReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/RoleRemove")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleRemove(ctx, req.(*RoleRemoveReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleRemoveReply)
		return ctx.Result(200, reply)
	}
}

func _User_RoleChange0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleChangeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/RoleChange")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleChange(ctx, req.(*RoleChangeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleChangeReply)
		return ctx.Result(200, reply)
	}
}

func _User_RoleGet0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleGetReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/RoleGet")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleGet(ctx, req.(*RoleGetReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleGetReply)
		return ctx.Result(200, reply)
	}
}

func _User_RoleGetAll0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleGetAllReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/RoleGetAll")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleGetAll(ctx, req.(*RoleGetAllReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleGetAllReply)
		return ctx.Result(200, reply)
	}
}

func _User_PermissionAdd0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionAddReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/PermissionAdd")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionAdd(ctx, req.(*PermissionAddReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionAddReply)
		return ctx.Result(200, reply)
	}
}

func _User_PermissionRemove0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionRemoveReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/PermissionRemove")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionRemove(ctx, req.(*PermissionRemoveReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionRemoveReply)
		return ctx.Result(200, reply)
	}
}

func _User_PermissionChange0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionChangeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/PermissionChange")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionChange(ctx, req.(*PermissionChangeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionChangeReply)
		return ctx.Result(200, reply)
	}
}

func _User_PermissionGet0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionGetReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/PermissionGet")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionGet(ctx, req.(*PermissionGetReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionGetReply)
		return ctx.Result(200, reply)
	}
}

func _User_PermissionGetAll0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionGetAllReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.v1.User/PermissionGetAll")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionGetAll(ctx, req.(*PermissionGetAllReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionGetAllReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	ListMemberGetBetweenBalance(ctx context.Context, req *ListMemberGetBetweenBalanceReq, opts ...http.CallOption) (rsp *ListMemberGetBetweenBalanceReply, err error)
	ListMemberGetBetweenRecharge(ctx context.Context, req *ListMemberGetBetweenRechargeReq, opts ...http.CallOption) (rsp *ListMemberGetBetweenRechargeReply, err error)
	ListMemberGetBetweenTime(ctx context.Context, req *ListMemberGetBetweenTimeReq, opts ...http.CallOption) (rsp *ListMemberGetBetweenTimeReply, err error)
	ListMemberGetByEmail(ctx context.Context, req *ListMemberGetByEmailReq, opts ...http.CallOption) (rsp *ListMemberGetByEmailReply, err error)
	ListMemberGetById(ctx context.Context, req *ListMemberGetByIdReq, opts ...http.CallOption) (rsp *ListMemberGetByIdReply, err error)
	ListMemberGetByName(ctx context.Context, req *ListMemberGetByNameReq, opts ...http.CallOption) (rsp *ListMemberGetByNameReply, err error)
	ListMemberGetByPhone(ctx context.Context, req *ListMemberGetByPhoneReq, opts ...http.CallOption) (rsp *ListMemberGetByPhoneReply, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutReq, opts ...http.CallOption) (rsp *LogoutReply, err error)
	MemberAdd(ctx context.Context, req *MemberAddReq, opts ...http.CallOption) (rsp *MemberAddReply, err error)
	MemberChange(ctx context.Context, req *MemberChangeReq, opts ...http.CallOption) (rsp *MemberChangeReply, err error)
	MemberGet(ctx context.Context, req *MemberGetReq, opts ...http.CallOption) (rsp *MemberGetReply, err error)
	MemberRemove(ctx context.Context, req *MemberRemoveReq, opts ...http.CallOption) (rsp *MemberRemoveReply, err error)
	PermissionAdd(ctx context.Context, req *PermissionAddReq, opts ...http.CallOption) (rsp *PermissionAddReply, err error)
	PermissionChange(ctx context.Context, req *PermissionChangeReq, opts ...http.CallOption) (rsp *PermissionChangeReply, err error)
	PermissionGet(ctx context.Context, req *PermissionGetReq, opts ...http.CallOption) (rsp *PermissionGetReply, err error)
	PermissionGetAll(ctx context.Context, req *PermissionGetAllReq, opts ...http.CallOption) (rsp *PermissionGetAllReply, err error)
	PermissionRemove(ctx context.Context, req *PermissionRemoveReq, opts ...http.CallOption) (rsp *PermissionRemoveReply, err error)
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterReply, err error)
	RoleAdd(ctx context.Context, req *RoleAddReq, opts ...http.CallOption) (rsp *RoleAddReply, err error)
	RoleChange(ctx context.Context, req *RoleChangeReq, opts ...http.CallOption) (rsp *RoleChangeReply, err error)
	RoleGet(ctx context.Context, req *RoleGetReq, opts ...http.CallOption) (rsp *RoleGetReply, err error)
	RoleGetAll(ctx context.Context, req *RoleGetAllReq, opts ...http.CallOption) (rsp *RoleGetAllReply, err error)
	RoleRemove(ctx context.Context, req *RoleRemoveReq, opts ...http.CallOption) (rsp *RoleRemoveReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) ListMemberGetBetweenBalance(ctx context.Context, in *ListMemberGetBetweenBalanceReq, opts ...http.CallOption) (*ListMemberGetBetweenBalanceReply, error) {
	var out ListMemberGetBetweenBalanceReply
	pattern := "/v1/member/list/get/balance/{low}/{high}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/ListMemberGetBetweenBalance"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListMemberGetBetweenRecharge(ctx context.Context, in *ListMemberGetBetweenRechargeReq, opts ...http.CallOption) (*ListMemberGetBetweenRechargeReply, error) {
	var out ListMemberGetBetweenRechargeReply
	pattern := "/v1/member/list/get/recharge/{low}/{high}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/ListMemberGetBetweenRecharge"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListMemberGetBetweenTime(ctx context.Context, in *ListMemberGetBetweenTimeReq, opts ...http.CallOption) (*ListMemberGetBetweenTimeReply, error) {
	var out ListMemberGetBetweenTimeReply
	pattern := "/v1/member/list/get/time/{start}/{end}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/ListMemberGetBetweenTime"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListMemberGetByEmail(ctx context.Context, in *ListMemberGetByEmailReq, opts ...http.CallOption) (*ListMemberGetByEmailReply, error) {
	var out ListMemberGetByEmailReply
	pattern := "/v1/member/list/get/email/{email}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/ListMemberGetByEmail"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListMemberGetById(ctx context.Context, in *ListMemberGetByIdReq, opts ...http.CallOption) (*ListMemberGetByIdReply, error) {
	var out ListMemberGetByIdReply
	pattern := "/v1/member/list/get/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/ListMemberGetById"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListMemberGetByName(ctx context.Context, in *ListMemberGetByNameReq, opts ...http.CallOption) (*ListMemberGetByNameReply, error) {
	var out ListMemberGetByNameReply
	pattern := "/v1/member/list/get/name/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/ListMemberGetByName"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ListMemberGetByPhone(ctx context.Context, in *ListMemberGetByPhoneReq, opts ...http.CallOption) (*ListMemberGetByPhoneReply, error) {
	var out ListMemberGetByPhoneReply
	pattern := "/v1/member/list/get/phone/{phone}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/ListMemberGetByPhone"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Logout(ctx context.Context, in *LogoutReq, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/v1/logout"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/Logout"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) MemberAdd(ctx context.Context, in *MemberAddReq, opts ...http.CallOption) (*MemberAddReply, error) {
	var out MemberAddReply
	pattern := "/v1/member/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/MemberAdd"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) MemberChange(ctx context.Context, in *MemberChangeReq, opts ...http.CallOption) (*MemberChangeReply, error) {
	var out MemberChangeReply
	pattern := "/v1/member/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/MemberChange"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) MemberGet(ctx context.Context, in *MemberGetReq, opts ...http.CallOption) (*MemberGetReply, error) {
	var out MemberGetReply
	pattern := "/v1/member/get/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/MemberGet"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) MemberRemove(ctx context.Context, in *MemberRemoveReq, opts ...http.CallOption) (*MemberRemoveReply, error) {
	var out MemberRemoveReply
	pattern := "/v1/member/remove/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/MemberRemove"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) PermissionAdd(ctx context.Context, in *PermissionAddReq, opts ...http.CallOption) (*PermissionAddReply, error) {
	var out PermissionAddReply
	pattern := "/v1/permission/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/PermissionAdd"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) PermissionChange(ctx context.Context, in *PermissionChangeReq, opts ...http.CallOption) (*PermissionChangeReply, error) {
	var out PermissionChangeReply
	pattern := "/v1/permission/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/PermissionChange"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) PermissionGet(ctx context.Context, in *PermissionGetReq, opts ...http.CallOption) (*PermissionGetReply, error) {
	var out PermissionGetReply
	pattern := "/v1/permission/get/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/PermissionGet"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) PermissionGetAll(ctx context.Context, in *PermissionGetAllReq, opts ...http.CallOption) (*PermissionGetAllReply, error) {
	var out PermissionGetAllReply
	pattern := "/v1/permission/get/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/PermissionGetAll"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) PermissionRemove(ctx context.Context, in *PermissionRemoveReq, opts ...http.CallOption) (*PermissionRemoveReply, error) {
	var out PermissionRemoveReply
	pattern := "/v1/permission/remove/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/PermissionRemove"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/Register"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RoleAdd(ctx context.Context, in *RoleAddReq, opts ...http.CallOption) (*RoleAddReply, error) {
	var out RoleAddReply
	pattern := "/v1/role/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/RoleAdd"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RoleChange(ctx context.Context, in *RoleChangeReq, opts ...http.CallOption) (*RoleChangeReply, error) {
	var out RoleChangeReply
	pattern := "/v1/role/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.v1.User/RoleChange"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RoleGet(ctx context.Context, in *RoleGetReq, opts ...http.CallOption) (*RoleGetReply, error) {
	var out RoleGetReply
	pattern := "/v1/role/get/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/RoleGet"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RoleGetAll(ctx context.Context, in *RoleGetAllReq, opts ...http.CallOption) (*RoleGetAllReply, error) {
	var out RoleGetAllReply
	pattern := "/v1/role/get/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/RoleGetAll"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RoleRemove(ctx context.Context, in *RoleRemoveReq, opts ...http.CallOption) (*RoleRemoveReply, error) {
	var out RoleRemoveReply
	pattern := "/v1/role/remove/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.v1.User/RoleRemove"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}