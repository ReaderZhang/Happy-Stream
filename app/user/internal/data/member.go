package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/qqz/Happy-Stream/app/user/internal/biz"
	"github.com/qqz/Happy-Stream/app/user/internal/data/model"
	"github.com/qqz/Happy-Stream/app/user/pkg/utils"
	"github.com/qqz/Happy-Stream/globalUtils"
	"strconv"
	"strings"
)

type memberRepo struct {
	data *Data
	log  *log.Helper
}

func (m *memberRepo) ListMember(ctx context.Context, param *biz.FormParam) (*biz.MemberTable, error) {
	mems := make([]*model.Member, param.Page.PageSize)
	d := m.data.db
	c := m.data.db
	var count int64
	var str strings.Builder
	str.WriteString("SELECT * FROM member WHERE ")
	if param.Status == "" {
		str.WriteString("deleted_at is null")
		d.Where("deleted_at = ?", "null")
		c.Where("deleted_at = ?", "null")
	} else if param.Status == "1" {
		str.WriteString("deleted_at != null AND")
		d.Where("deleted_at != ?", "null")
		c.Where("deleted_at != ?", "null")
	}
	if param.Bhigh != 0 || param.Blow != 0 {
		str.WriteString("(balance between ")
		a := strconv.FormatFloat(param.Bhigh, 'f', 10, 64)
		b := strconv.FormatFloat(param.Blow, 'f', 10, 64)
		str.WriteString(a)
		str.WriteString(" and ")
		str.WriteString(b)
		str.WriteString(")")
		d.Where("balance between ? and ?", param.Bhigh, param.Blow)
		c.Where("balance between ? and ?", param.Bhigh, param.Blow)
	}
	if param.Rhigh != -1 || param.Rlow != -1 {
		d.Where("balance between ? and ?", param.Rhigh, param.Rlow)
		c.Where("balance between ? and ?", param.Rhigh, param.Rlow)
	}
	if param.Start != "" && param.End != "" {
		d.Where("created_at between ? and ?", param.Start, param.End)
		c.Where("created_at between ? and ?", param.Start, param.End)
	}
	if param.Search != "" {
		str.WriteString(" AND  (username like '")
		str.WriteString(param.Search)
		str.WriteString("%' or phone like '")
		str.WriteString(param.Search)
		str.WriteString("%' or email like '")
		str.WriteString(param.Search)
		str.WriteString("%')")
		d.Where("username like ?% or phone like ?% or email like ?%", param.Search, param.Search, param.Search)
		c.Where("username like ?% or phone like ?% or email like ?%", param.Search, param.Search, param.Search)
	}
	//result := d.Limit(int(param.Page.PageSize)).Offset(int(param.Page.PageNum-1) * int(param.Page.PageNum)).Find(&mems)
	result := m.data.db.Raw(str.String()).Scan(&mems)
	c.Model(&model.Member{}).Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}
	members := make([]*biz.Member, param.Page.PageSize)

	for i := range mems {
		if mems[i] != nil {
			m := &biz.Member{
				Id:         int64(mems[i].ID),
				Username:   mems[i].Username,
				Phone:      mems[i].Phone,
				Email:      mems[i].Email,
				Recharge:   mems[i].Balance,
				Balance:    mems[i].Recharge,
				Created_at: mems[i].CreatedAt.String(),
			}
			members[i] = m
		}
	}
	return &biz.MemberTable{
		Members: members,
		Page: &globalUtils.Pageination{
			PageNum:  param.Page.PageNum,
			PageSize: param.Page.PageSize,
			Total:    count,
		},
	}, nil

}

func (m *memberRepo) CreateMember(ctx context.Context, u *biz.Member) (bool, error) {
	member := &model.Member{
		ID:       int32(u.Id),
		Username: u.Username,
		Password: Sha256(u.Password),
		Phone:    u.Phone,
		Email:    u.Email,
		Avator:   u.Avator,
	}
	result := m.data.db.Create(member)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *memberRepo) DeleteMember(ctx context.Context, id int64) (bool, error) {
	result := m.data.db.Delete(&model.Member{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (m *memberRepo) UpdateMember(ctx context.Context, u *biz.Member) (bool, error) {
	node, err := utils.NewWorker(1)
	if err != nil {
		return false, err
	}
	id := node.GetId()
	member := &model.Member{
		ID:       int32(id),
		Username: u.Username,
		Password: Sha256(u.Password),
		Phone:    u.Phone,
		Email:    u.Email,
		Avator:   u.Avator,
	}
	result := m.data.db.Save(&member)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (m *memberRepo) QueryMemberById(ctx context.Context, id int64) (*biz.Member, error) {
	member := &model.Member{}
	m.data.db.First(member, id)
	return &biz.Member{
		Id:       id,
		Username: member.Username,
		Phone:    member.Phone,
		Email:    member.Email,
		Avator:   member.Avator,
	}, nil
}

func (m *memberRepo) ListQueryMemberById(ctx context.Context, id []int64, pageSize int64, pageNum int64) (*biz.MemberResult, error) {
	members := make([]*biz.Member, pageSize)
	mems := make([]*model.Member, pageSize)
	result := m.data.db.Limit(int(pageSize)).Offset(int(pageNum-1)*int(pageNum)).Find(&mems, id)
	copier.Copy(&members, &mems)
	if result.Error != nil {
		return nil, result.Error
	}
	var count int64
	m.data.db.Count(&count)
	page := &globalUtils.Pageination{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    count,
	}
	rs := &biz.MemberResult{
		Members: members,
		Page:    page,
	}
	return rs, nil
}

func (m *memberRepo) ListQueryMemberByName(ctx context.Context, name string, pageSize int64, pageNum int64) (*biz.MemberResult, error) {
	members := make([]*biz.Member, pageSize)
	mems := make([]*model.Member, pageSize)
	result := m.data.db.Where("username LIKE ?", "%"+name+"%").Limit(int(pageSize)).Offset(int(pageNum-1) * int(pageNum)).Find(&mems)
	copier.Copy(&members, &mems)
	if result.Error != nil {
		return nil, result.Error
	}
	var count int64
	m.data.db.Count(&count)
	page := &globalUtils.Pageination{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    count,
	}
	rs := &biz.MemberResult{
		Members: members,
		Page:    page,
	}
	return rs, nil
}

func (m *memberRepo) ListQueryMemberByPhone(ctx context.Context, phone string, pageSize int64, pageNum int64) (*biz.MemberResult, error) {
	members := make([]*biz.Member, pageSize)
	mems := make([]*model.Member, pageSize)
	result := m.data.db.Where("phone LIKE ?", "%"+phone+"%").Limit(int(pageSize)).Offset(int(pageNum-1) * int(pageNum)).Find(&mems)
	copier.Copy(&members, &mems)
	if result.Error != nil {
		return nil, result.Error
	}
	var count int64
	m.data.db.Count(&count)
	page := &globalUtils.Pageination{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    count,
	}
	rs := &biz.MemberResult{
		Members: members,
		Page:    page,
	}
	return rs, nil
}

func (m *memberRepo) ListQueryMemberByEmail(ctx context.Context, email string, pageSize int64, pageNum int64) (*biz.MemberResult, error) {
	members := make([]*biz.Member, pageSize)
	mems := make([]*model.Member, pageSize)
	result := m.data.db.Where("email LIKE ?", "%"+email+"%").Limit(int(pageSize)).Offset(int(pageNum-1) * int(pageNum)).Find(&mems)
	copier.Copy(&members, &mems)
	if result.Error != nil {
		return nil, result.Error
	}
	var count int64
	m.data.db.Count(&count)
	page := &globalUtils.Pageination{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    count,
	}
	rs := &biz.MemberResult{
		Members: members,
		Page:    page,
	}
	return rs, nil
}

func (m *memberRepo) ListQueryMemberBetweenRecharge(ctx context.Context, low float64, high float64, pageSize int64, pageNum int64) (*biz.MemberResult, error) {
	members := make([]*biz.Member, pageSize)
	mems := make([]*model.Member, pageSize)
	result := m.data.db.Where("recharge BETWEEN ? AND ?", low, high).Limit(int(pageSize)).Offset(int(pageNum-1) * int(pageNum)).Find(&mems)
	copier.Copy(&members, &mems)
	if result.Error != nil {
		return nil, result.Error
	}
	var count int64
	m.data.db.Count(&count)
	page := &globalUtils.Pageination{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    count,
	}
	rs := &biz.MemberResult{
		Members: members,
		Page:    page,
	}
	return rs, nil
}

func (m *memberRepo) ListQueryMemberBetweenBalance(ctx context.Context, low float64, high float64, pageSize int64, pageNum int64) (*biz.MemberResult, error) {
	members := make([]*biz.Member, pageSize)
	mems := make([]*model.Member, pageSize)
	result := m.data.db.Where("balance BETWEEN ? AND ?", low, high).Limit(int(pageSize)).Offset(int(pageNum-1) * int(pageNum)).Find(&mems)
	copier.Copy(&members, &mems)
	if result.Error != nil {
		return nil, result.Error
	}
	var count int64
	m.data.db.Count(&count)
	page := &globalUtils.Pageination{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    count,
	}
	rs := &biz.MemberResult{
		Members: members,
		Page:    page,
	}
	return rs, nil
}

func (m memberRepo) ListQueryByTime(ctx context.Context, start string, end string, pageSize int64, pageNum int64) (*biz.MemberResult, error) {
	members := make([]*biz.Member, pageSize)
	mems := make([]*model.Member, pageSize)
	result := m.data.db.Where("created_at BETWEEN ? AND ?", start, end).Limit(int(pageSize)).Offset(int(pageNum-1) * int(pageNum)).Find(&mems)
	if result.Error != nil {
		return nil, result.Error
	}
	var count int64
	m.data.db.Model(&model.Member{}).Count(&count)
	for i := range mems {
		m := &biz.Member{
			Id:         int64(mems[i].ID),
			Username:   mems[i].Username,
			Phone:      mems[i].Phone,
			Email:      mems[i].Email,
			Recharge:   mems[i].Balance,
			Balance:    mems[i].Recharge,
			Created_at: mems[i].CreatedAt.String(),
		}
		members[i] = m
	}
	page := &globalUtils.Pageination{
		PageNum:  pageNum,
		PageSize: pageSize,
		Total:    count,
	}
	rs := &biz.MemberResult{
		Members: members,
		Page:    page,
	}
	return rs, nil
}

func (m *memberRepo) LoadByUsername(ctx context.Context, username string) (*biz.Member, error) {
	member := &biz.Member{}
	mem := &model.Member{}
	result := m.data.db.Where("username = ?", username).First(&mem)
	copier.Copy(&member, &mem)
	if result.Error != nil {
		return nil, result.Error
	}
	return member, nil
}
func NewMemberRepo(data *Data, logger log.Logger) biz.MemberRepo {
	return &memberRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
