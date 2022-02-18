package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/qqz/Happy-Stream/api/user/v1"
)

type Member struct {
	Id         int64
	Username   string
	Password   string
	Phone      string
	Email      string
	Avator     string
	Recharge   float64
	Balance    float64
	Created_at string
}

type MemberRepo interface {
	CreateMember(ctx context.Context, u *Member) (bool, error)
	DeleteMember(ctx context.Context, id int64) (bool, error)
	UpdateMember(ctx context.Context, u *Member) (bool, error)
	QueryMemberById(ctx context.Context, id int64) (*Member, error)
	ListQueryMemberById(ctx context.Context, id int64, page_size int64, page_num int64) ([]*Member, error)
	ListQueryMemberByName(ctx context.Context, name string, page_size int64, page_num int64) ([]*Member, error)
	ListQueryMemberByPhone(ctx context.Context, phone string, page_size int64, page_num int64) ([]*Member, error)
	ListQueryMemberByEmail(ctx context.Context, email string, page_size int64, page_num int64) ([]*Member, error)
	ListQueryMemberBetweenRecharge(ctx context.Context, low float64, high float64, page_size int64, page_num int64) ([]*Member, error)
	ListQueryMemberBetweenBalance(ctx context.Context, low float64, high float64, page_size int64, page_num int64) ([]*Member, error)
	ListQueryByTime(ctx context.Context, start string, end string, page_size int64, page_num int64) ([]*Member, error)
}

type MemberUseCase struct {
	repo MemberRepo
	log  *log.Helper
}

func NewMemberUseCase(repo MemberRepo, logger log.Logger) *MemberUseCase {
	return &MemberUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (mc *MemberUseCase) Save(ctx context.Context, in v1.MemberAddReq) (*v1.MemberAddReply, error) {
	member := &Member{
		Username: in.Member.Username,
		Phone:    in.Member.Phone,
		Password: in.Member.Password,
		Email:    in.Member.Email,
		Avator:   in.Member.Avator,
	}
	_, err := mc.repo.CreateMember(ctx, member)
	if err != nil {
		return nil, err
	}
	return &v1.MemberAddReply{
		Ok: true,
	}, nil
}

func (mc *MemberUseCase) Remove(ctx context.Context, in v1.MemberRemoveReq) (*v1.MemberRemoveReply, error) {
	_, err := mc.repo.DeleteMember(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.MemberRemoveReply{
		Ok: true,
	}, nil
}

func (mc *MemberUseCase) Update(ctx context.Context, in v1.MemberChangeReq) (*v1.MemberChangeReply, error) {
	member := &Member{
		Username: in.Member.Username,
		Phone:    in.Member.Phone,
		Password: in.Member.Password,
		Email:    in.Member.Email,
		Avator:   in.Member.Avator,
		Recharge: in.Member.Recharge,
		Balance:  in.Member.Balance,
	}
	_, err := mc.repo.UpdateMember(ctx, member)

	if err != nil {
		return nil, err
	}

	return &v1.MemberChangeReply{
		Ok: true,
	}, nil
}
