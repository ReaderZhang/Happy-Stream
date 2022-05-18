package biz

import (
	"context"
	"github.com/qqz/Happy-Stream/globalUtils"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
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
	Deleted_at bool
}

type FormParam struct {
	Search string
	Type   string
	Start  string
	End    string
	Rlow   float64
	Rhigh  float64
	Blow   float64
	Bhigh  float64
	Status string
	Page   *globalUtils.Pageination
}

type MemberTable struct {
	Members []*Member
	Page    *globalUtils.Pageination
}

type MemberResult struct {
	Members []*Member
	Page    *globalUtils.Pageination
}

type MemberRepo interface {
	CreateMember(ctx context.Context, u *Member) (bool, error)
	DeleteMember(ctx context.Context, id int64) (bool, error)
	UpdateMember(ctx context.Context, u *Member) (bool, error)
	QueryMemberById(ctx context.Context, id int64) (*Member, error)
	ListQueryMemberById(ctx context.Context, id []int64, pageSize int64, pageNum int64) (*MemberResult, error)
	ListQueryMemberByName(ctx context.Context, name string, pageSize int64, pageNum int64) (*MemberResult, error)
	ListQueryMemberByPhone(ctx context.Context, phone string, pageSize int64, pageNum int64) (*MemberResult, error)
	ListQueryMemberByEmail(ctx context.Context, email string, pageSize int64, pageNum int64) (*MemberResult, error)
	ListQueryMemberBetweenRecharge(ctx context.Context, low float64, high float64, pageSize int64, pageNum int64) (*MemberResult, error)
	ListQueryMemberBetweenBalance(ctx context.Context, low float64, high float64, pageSize int64, pageNum int64) (*MemberResult, error)
	ListQueryByTime(ctx context.Context, start string, end string, pageSize int64, pageNum int64) (*MemberResult, error)

	LoadByUsername(ctx context.Context, username string) (*Member, error)
	ListMember(ctx context.Context, param *FormParam) (*MemberTable, error)
}

type MemberUseCase struct {
	repo MemberRepo
	log  *log.Helper

	authUc *AuthUseCase
}

func NewMemberUseCase(repo MemberRepo, logger log.Logger) *MemberUseCase {
	return &MemberUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (mc *MemberUseCase) Save(ctx context.Context, member *Member) (bool, error) {
	_, err := mc.repo.CreateMember(ctx, member)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (mc *MemberUseCase) Remove(ctx context.Context, id int64) (bool, error) {
	_, err := mc.repo.DeleteMember(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (mc *MemberUseCase) Update(ctx context.Context, member *Member) (bool, error) {
	updateMember, err := mc.repo.UpdateMember(ctx, member)
	return updateMember, err
}

func (mc *MemberUseCase) Get(ctx context.Context, id int64) (*Member, error) {
	member, err := mc.repo.QueryMemberById(ctx, id)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (mc *MemberUseCase) List(ctx context.Context, pageInfo globalUtils.Pageination, params string, typeof string) (*MemberResult, error) {
	switch typeof {
	//批量id检索
	case "batch_ids":
		{
			num_ids := strings.Split(params, ",")
			length := len(num_ids)
			ids := make([]int64, length)
			for i := 0; i < length; i++ {
				num_id, _ := strconv.ParseInt(num_ids[i], 0, 10)
				ids = append(ids, num_id)
			}
			mems, err := mc.repo.ListQueryMemberById(ctx, ids, pageInfo.PageSize, pageInfo.PageNum)
			if err != nil {
				return nil, err
			}
			return mems, nil

		}
		//创建时间检索
	case "created_time":
		{
			times := strings.Split(params, ",")
			start, end := times[0], times[1]
			mems, err := mc.repo.ListQueryByTime(ctx, start, end, pageInfo.PageSize, pageInfo.PageNum)
			if err != nil {
				return nil, err
			}
			return mems, nil
		}
	case "email":
		{
			mems, err := mc.repo.ListQueryMemberByEmail(ctx, params, pageInfo.PageSize, pageInfo.PageNum)
			if err != nil {
				return nil, err
			}
			return mems, nil
		}
	case "phone":
		{
			mems, err := mc.repo.ListQueryMemberByPhone(ctx, params, pageInfo.PageSize, pageInfo.PageNum)
			if err != nil {
				return nil, err
			}
			return mems, nil
		}
	case "name":
		{
			mems, err := mc.repo.ListQueryMemberByName(ctx, params, pageInfo.PageSize, pageInfo.PageNum)
			if err != nil {
				return nil, err
			}

			return mems, nil
		}
	case "balance":
		{
			balance := strings.Split(params, ",")
			start, _ := strconv.ParseFloat(balance[0], 64)
			end, _ := strconv.ParseFloat(balance[1], 64)
			mems, err := mc.repo.ListQueryMemberBetweenBalance(ctx, start, end, pageInfo.PageSize, pageInfo.PageNum)
			if err != nil {
				return nil, err
			}
			return mems, nil
		}
	case "recharge":
		{
			recharge := strings.Split(params, ",")
			start, _ := strconv.ParseFloat(recharge[0], 64)
			end, _ := strconv.ParseFloat(recharge[1], 64)
			mems, err := mc.repo.ListQueryMemberBetweenRecharge(ctx, start, end, pageInfo.PageSize, pageInfo.PageNum)
			if err != nil {
				return nil, err
			}
			return mems, nil
		}
	}
	return nil, nil
}

func (mc *MemberUseCase) LoadByUsername(ctx context.Context, username string) (*Member, error) {
	return mc.repo.LoadByUsername(ctx, username)
}

func (mc *MemberUseCase) Table(ctx context.Context, param *FormParam) (*MemberTable, error) {
	return mc.repo.ListMember(ctx, param)
}
