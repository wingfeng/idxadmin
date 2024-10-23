// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/user/v1"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type (
	IUser interface {
		Get(ctx context.Context, id int64) (*entity.Users, error)
		Delete(ctx context.Context, id int64) error
		Save(ctx context.Context, req v1.SaveReq) (err error)
		List(ctx context.Context, req v1.PageReq) (*v1.PageRes, error)
		ResetPwd(ctx context.Context, req v1.ResetPwdReq) (string, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
