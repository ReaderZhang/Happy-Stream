package biz

import (
	"context"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
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
	ListQueryMemberById(ctx context.Context, id []int64, page_size int64, page_num int64) ([]*Member, int64, error)
	ListQueryMemberByName(ctx context.Context, name string, page_size int64, page_num int64) ([]*Member, int64, error)
	ListQueryMemberByPhone(ctx context.Context, phone string, page_size int64, page_num int64) ([]*Member, int64, error)
	ListQueryMemberByEmail(ctx context.Context, email string, page_size int64, page_num int64) ([]*Member, int64, error)
	ListQueryMemberBetweenRecharge(ctx context.Context, low float64, high float64, page_size int64, page_num int64) ([]*Member, int64, error)
	ListQueryMemberBetweenBalance(ctx context.Context, low float64, high float64, page_size int64, page_num int64) ([]*Member, int64, error)
	ListQueryByTime(ctx context.Context, start string, end string, page_size int64, page_num int64) ([]*Member, int64, error)
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

func (mc *MemberUseCase) List(ctx context.Context, pageinfo v1.PageInfoReq, params string, typeof string, in v1.ListMemberGetByIdReq) (*v1.ListMemberGetBetweenBalanceReply, error) {
	switch typeof {
	//批量id检索
	case "a":
		{
			num_ids := strings.Split(params, ",")
			length := len(num_ids)
			ids := make([]int64, length)
			members := make([]*v1.MemberDTO, length)
			for i := 0; i < length; i++ {
				num_id, _ := strconv.ParseInt(num_ids[i], 0, 10)
				ids = append(ids, num_id)
			}
			mems, total, err := mc.repo.ListQueryMemberById(ctx, ids, pageinfo.PageSize, pageinfo.PageNum)
			if err != nil {
				return nil, err
			}
			copier.Copy(&members, &mems)
			return &v1.ListMemberGetBetweenBalanceReply{
				Members: members,
				Page: &v1.PageInfoReply{
					PageNum:  pageinfo.PageNum,
					PageSize: pageinfo.PageSize,
					Total:    total,
				},
			}, nil

		}
		//创建时间检索
	case "b":
		{

		}
	}
	return nil, nil
}
