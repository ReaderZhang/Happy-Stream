package data

import (
	"context"
	sha2562 "crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/qqz/Happy-Stream/app/user/internal/biz"
	"github.com/qqz/Happy-Stream/app/user/internal/data/model"
	"github.com/qqz/Happy-Stream/app/user/pkg/utils"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func (a authRepo) Login(ctx context.Context, am *biz.AuthMember) (*biz.AuthResult, error) {
	password := Sha256(am.Password)
	member := &model.Member{}
	result := a.data.db.Where("password=?", password).Or(model.Member{Username: am.Username, Phone: am.Username}).Find(member)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Print(member.ID)
	if member.ID == 0 {
		return nil, errors.New(400, "用户不存在", "帐号或密码错误")
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": am.Username,
		"id":       member.ID,
	})
	signedString, err := claims.SignedString([]byte("qqz"))

	if err != nil {
		return nil, errors.Errorf(401, "认证失败", "签名出错")
	}
	return &biz.AuthResult{
		Avator:   member.Avator,
		Username: member.Username,
		Token:    signedString,
	}, nil
}

func (a authRepo) Register(ctx context.Context, am *biz.AuthMember) (*biz.AuthResult, error) {
	password := Sha256(am.Password)
	node, err := utils.NewWorker(1)
	id := node.GetId()
	member := &model.Member{
		ID:       int32(id),
		Username: am.Username,
		Phone:    am.Phone,
		Password: password,
	}
	result := a.data.db.Create(&member)
	if result.Error != nil {
		return nil, result.Error
	}
	info, err := a.Login(ctx, am)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (a authRepo) Logout(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func Sha256(data string) string {
	sha256 := sha2562.New()
	sha256.Write([]byte(data))
	return hex.EncodeToString(sha256.Sum([]byte("")))
}
