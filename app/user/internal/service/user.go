package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jinzhu/copier"
	"github.com/qqz/Happy-Stream/app/user/internal/biz"
	"github.com/qqz/Happy-Stream/globalUtils"
	"strconv"
	"strings"
	"unicode"

	pb "github.com/qqz/Happy-Stream/api/user/v1"
)

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	rv, err := s.ac.Login(ctx, &biz.AuthMember{
		Username: req.Username,
		Phone:    req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token:    rv.Token,
		Avator:   rv.Avator,
		Username: rv.Username,
	}, nil
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	if len(req.Phone) != 11 {
		return nil, errors.New(400, "手机号长度不正确", "手机号需为11位")
	}
	for _, r := range req.Phone {
		if !unicode.IsNumber(r) {
			return nil, errors.New(400, "手机号格式错误", "手机号不能存在字母")
		}

	}

	rv, err := s.ac.Register(ctx, &biz.AuthMember{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{
		Username: rv.Username,
		Token:    rv.Token,
		Avator:   rv.Avator,
	}, nil
}
func (s *UserService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{
		Ok: true,
	}, nil
}
func (s *UserService) MemberAdd(ctx context.Context, req *pb.MemberAddReq) (*pb.MemberAddReply, error) {
	save, err := s.mc.Save(ctx, &biz.Member{
		Username: req.Member.Username,
		Password: req.Member.Password,
		Phone:    req.Member.Phone,
		Email:    req.Member.Email,
		Avator:   req.Member.Avator,
		Recharge: req.Member.Recharge,
		Balance:  req.Member.Balance,
	})
	if err != nil {
		return nil, err
	}
	return &pb.MemberAddReply{
		Ok: save,
	}, nil
}
func (s *UserService) MemberRemove(ctx context.Context, req *pb.MemberRemoveReq) (*pb.MemberRemoveReply, error) {
	remove, err := s.mc.Remove(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.MemberRemoveReply{Ok: remove}, nil
}
func (s *UserService) MemberChange(ctx context.Context, req *pb.MemberChangeReq) (*pb.MemberChangeReply, error) {
	update, err := s.mc.Update(ctx, &biz.Member{
		Username: req.Member.Username,
		Password: req.Member.Password,
		Phone:    req.Member.Phone,
		Email:    req.Member.Email,
		Avator:   req.Member.Avator,
		Recharge: req.Member.Recharge,
		Balance:  req.Member.Balance,
	})
	if err != nil {
		return nil, err
	}
	return &pb.MemberChangeReply{Ok: update}, nil
}
func (s *UserService) MemberGet(ctx context.Context, req *pb.MemberGetReq) (*pb.MemberGetReply, error) {
	id, serr := strconv.ParseInt(req.Id, 10, 64)
	if serr != nil {
		return nil, serr
	}
	member, err := s.mc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	memberDTO := &pb.MemberDTO{}
	copier.Copy(&memberDTO, &member)
	return &pb.MemberGetReply{Members: memberDTO}, nil
}
func (s *UserService) ListMemberGetById(ctx context.Context, req *pb.ListMemberGetByIdReq) (*pb.ListMemberGetByIdReply, error) {
	pageinfo := &globalUtils.Pageination{
		PageNum:  req.Page.PageNum,
		PageSize: req.Page.PageSize,
	}
	id := strconv.FormatInt(req.Id, 64)
	list, err := s.mc.List(ctx, *pageinfo, id, "batch_ids")
	if err != nil {
		return nil, err
	}
	mems := make([]*pb.MemberDTO, pageinfo.PageSize)

	members := list.Members
	copier.Copy(&mems, &members)
	return &pb.ListMemberGetByIdReply{
		Members: mems,
		Page: &pb.PageInfoReply{
			PageNum:  list.Page.PageNum,
			PageSize: list.Page.PageSize,
			Total:    list.Page.Total,
		},
	}, nil
}
func (s *UserService) ListMemberGetByName(ctx context.Context, req *pb.ListMemberGetByNameReq) (*pb.ListMemberGetByNameReply, error) {
	pageinfo := &globalUtils.Pageination{
		PageNum:  req.Page.PageNum,
		PageSize: req.Page.PageSize,
	}
	list, err := s.mc.List(ctx, *pageinfo, req.Name, "name")
	if err != nil {
		return nil, err
	}
	mems := make([]*pb.MemberDTO, pageinfo.PageSize)
	members := list.Members
	copier.Copy(&mems, &members)
	return &pb.ListMemberGetByNameReply{
		Members: mems,
		Page: &pb.PageInfoReply{
			PageNum:  list.Page.PageNum,
			PageSize: list.Page.PageSize,
			Total:    list.Page.Total,
		},
	}, nil
}
func (s *UserService) ListMemberGetByPhone(ctx context.Context, req *pb.ListMemberGetByPhoneReq) (*pb.ListMemberGetByPhoneReply, error) {
	pageinfo := &globalUtils.Pageination{
		PageNum:  req.Page.PageNum,
		PageSize: req.Page.PageSize,
	}
	list, err := s.mc.List(ctx, *pageinfo, req.Phone, "phone")
	if err != nil {
		return nil, err
	}
	mems := make([]*pb.MemberDTO, pageinfo.PageSize)
	members := list.Members
	copier.Copy(&mems, &members)
	return &pb.ListMemberGetByPhoneReply{
		Members: mems,
		Page: &pb.PageInfoReply{
			PageNum:  list.Page.PageNum,
			PageSize: list.Page.PageSize,
			Total:    list.Page.Total,
		},
	}, nil
}
func (s *UserService) ListMemberGetByEmail(ctx context.Context, req *pb.ListMemberGetByEmailReq) (*pb.ListMemberGetByEmailReply, error) {
	pageinfo := &globalUtils.Pageination{
		PageNum:  req.Page.PageNum,
		PageSize: req.Page.PageSize,
	}
	list, err := s.mc.List(ctx, *pageinfo, req.Email, "phone")
	if err != nil {
		return nil, err
	}
	mems := make([]*pb.MemberDTO, pageinfo.PageSize)
	members := list.Members
	copier.Copy(&mems, &members)
	return &pb.ListMemberGetByEmailReply{
		Members: mems,
		Page: &pb.PageInfoReply{
			PageNum:  list.Page.PageNum,
			PageSize: list.Page.PageSize,
			Total:    list.Page.Total,
		},
	}, nil
}
func (s *UserService) ListMemberGetBetweenRecharge(ctx context.Context, req *pb.ListMemberGetBetweenRechargeReq) (*pb.ListMemberGetBetweenRechargeReply, error) {
	pageinfo := &globalUtils.Pageination{
		PageNum:  req.Page.PageNum,
		PageSize: req.Page.PageSize,
	}
	var builder strings.Builder
	builder.WriteString(strconv.FormatFloat(req.Low, 'f', 2, 64))
	builder.WriteString(",")
	builder.WriteString(strconv.FormatFloat(req.High, 'f', 2, 64))
	list, err := s.mc.List(ctx, *pageinfo, builder.String(), "recharge")
	if err != nil {
		return nil, err
	}
	mems := make([]*pb.MemberDTO, pageinfo.PageSize)
	members := list.Members
	copier.Copy(&mems, &members)
	return &pb.ListMemberGetBetweenRechargeReply{
		Members: mems,
		Page: &pb.PageInfoReply{
			PageNum:  list.Page.PageNum,
			PageSize: list.Page.PageSize,
			Total:    list.Page.Total,
		},
	}, nil
}
func (s *UserService) ListMemberGetBetweenBalance(ctx context.Context, req *pb.ListMemberGetBetweenBalanceReq) (*pb.ListMemberGetBetweenBalanceReply, error) {
	pageinfo := &globalUtils.Pageination{
		PageNum:  req.Page.PageNum,
		PageSize: req.Page.PageSize,
	}
	var builder strings.Builder
	builder.WriteString(strconv.FormatFloat(req.Low, 'f', 2, 64))
	builder.WriteString(",")
	builder.WriteString(strconv.FormatFloat(req.High, 'f', 2, 64))
	list, err := s.mc.List(ctx, *pageinfo, builder.String(), "balance")
	if err != nil {
		return nil, err
	}
	mems := make([]*pb.MemberDTO, pageinfo.PageSize)
	members := list.Members
	copier.Copy(&mems, &members)
	return &pb.ListMemberGetBetweenBalanceReply{
		Members: mems,
		Page: &pb.PageInfoReply{
			PageNum:  list.Page.PageNum,
			PageSize: list.Page.PageSize,
			Total:    list.Page.Total,
		},
	}, nil
}
func (s *UserService) ListMemberGetBetweenTime(ctx context.Context, req *pb.ListMemberGetBetweenTimeReq) (*pb.ListMemberGetBetweenTimeReply, error) {
	pageinfo := &globalUtils.Pageination{
		PageNum:  int64(req.PageNum),
		PageSize: int64(req.PageSize),
	}
	var builder strings.Builder
	builder.WriteString(req.Start)
	builder.WriteString(",")
	builder.WriteString(req.End)
	list, err := s.mc.List(ctx, *pageinfo, builder.String(), "created_time")
	if err != nil {
		return nil, err
	}

	members := list.Members
	mems := make([]*pb.MemberDTO, len(members))
	fmt.Println(len(members))
	fmt.Println(list.Page.Total)
	for i := range members {
		m := &pb.MemberDTO{
			Id:        members[i].Id,
			Username:  members[i].Username,
			Phone:     members[i].Phone,
			Email:     members[i].Email,
			Recharge:  members[i].Recharge,
			Balance:   members[i].Balance,
			CreatedAt: members[i].Created_at,
		}
		mems[i] = m
	}
	return &pb.ListMemberGetBetweenTimeReply{
		Members: mems,
		Page: &pb.PageInfoReply{
			PageNum:  list.Page.PageNum,
			PageSize: list.Page.PageSize,
			Total:    list.Page.Total,
		},
	}, nil
}

func (s *UserService) ListMembers(ctx context.Context, req *pb.MemberForm) (*pb.MemberTableReply, error) {
	fmt.Println(req.Search)
	if req.PageSize == 0 {
		req.PageSize = 5
	}
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	pageinfo := &globalUtils.Pageination{PageNum: req.PageNum, PageSize: req.PageSize}
	list, err := s.mc.Table(ctx, &biz.FormParam{
		Search: req.Search,
		Type:   req.Type,
		Start:  req.Start,
		End:    req.End,
		Rlow:   req.Rlow,
		Rhigh:  req.Rhigh,
		Blow:   req.Blow,
		Bhigh:  req.Bhight,
		Status: req.Status,
		Page:   pageinfo,
	})
	if err != nil {
		return nil, err
	}
	members := list.Members
	count := 0
	for i := range members {
		if members[i] != nil {
			count++
		}
	}
	mems := make([]*pb.MemberTable, count)
	for i := range members {
		if members[i] == nil {
			break
		}
		var flag int32
		if members[i].Deleted_at == true {
			flag = 1
		} else {
			flag = 0
		}
		m := &pb.MemberTable{
			Id:        members[i].Id,
			Username:  members[i].Username,
			Email:     members[i].Email,
			Phone:     members[i].Phone,
			Balance:   members[i].Balance,
			Recharge:  members[i].Recharge,
			CreatedAt: members[i].Created_at,
			Status:    flag,
		}
		mems[i] = m
	}
	reply := &pb.MemberTableReply{
		Table: mems,
		Page: &pb.PageInfoReply{
			PageNum:  pageinfo.PageNum,
			PageSize: pageinfo.PageSize,
			Total:    list.Page.Total,
		},
	}
	return reply, nil
}

func (s *UserService) RoleAdd(ctx context.Context, req *pb.RoleAddReq) (*pb.RoleAddReply, error) {
	err := s.rc.Save(ctx, req.Role.RoleName, 0)
	if err != nil {
		return nil, err
	}
	return &pb.RoleAddReply{Ok: true}, nil
}
func (s *UserService) RoleRemove(ctx context.Context, req *pb.RoleRemoveReq) (*pb.RoleRemoveReply, error) {
	err := s.rc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.RoleRemoveReply{Ok: true}, nil
}
func (s *UserService) RoleChange(ctx context.Context, req *pb.RoleChangeReq) (*pb.RoleChangeReply, error) {
	err := s.rc.Update(ctx, req.Role.RoleName, 0)
	if err != nil {
		return nil, err
	}
	return &pb.RoleChangeReply{Ok: true}, nil
}
func (s *UserService) RoleGet(ctx context.Context, req *pb.RoleGetReq) (*pb.RoleGetReply, error) {
	//query, err := s.rc.Query(ctx, req.Id)
	//if err != nil{
	//	return nil,err
	//}
	//roles := make([]pb.RoleDTO,1)
	//strs := make([]string,1)
	//strs = append(strs, query.Permission.Url)
	//p := &pb.RoleDTO{
	//	Id:          query.Id,
	//	RoleName:    query.Name,
	//	Permissions: query.Permission,
	//}
	//roles = append(roles,)
	return &pb.RoleGetReply{}, nil
}
func (s *UserService) RoleGetAll(ctx context.Context, req *pb.RoleGetAllReq) (*pb.RoleGetAllReply, error) {
	return &pb.RoleGetAllReply{}, nil
}
func (s *UserService) PermissionAdd(ctx context.Context, req *pb.PermissionAddReq) (*pb.PermissionAddReply, error) {
	return &pb.PermissionAddReply{}, nil
}
func (s *UserService) PermissionRemove(ctx context.Context, req *pb.PermissionRemoveReq) (*pb.PermissionRemoveReply, error) {
	return &pb.PermissionRemoveReply{}, nil
}
func (s *UserService) PermissionChange(ctx context.Context, req *pb.PermissionChangeReq) (*pb.PermissionChangeReply, error) {
	return &pb.PermissionChangeReply{}, nil
}
func (s *UserService) PermissionGet(ctx context.Context, req *pb.PermissionGetReq) (*pb.PermissionGetReply, error) {
	return &pb.PermissionGetReply{}, nil
}
func (s *UserService) PermissionGetAll(ctx context.Context, req *pb.PermissionGetAllReq) (*pb.PermissionGetAllReply, error) {
	return &pb.PermissionGetAllReply{}, nil
}
