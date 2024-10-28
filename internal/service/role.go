// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/role/v1"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type (
	IRole interface {
		Delete(ctx context.Context, id int64) error
		Get(ctx context.Context, id int64) (entity.Roles, error)
		List(ctx context.Context, req *v1.ListReq) (*model.PageRes, error)
		Save(ctx context.Context, req *entity.Roles) error
		Members(ctx context.Context, req *v1.MembersReq) (*model.PageRes, error)
		AddMemeber(ctx context.Context, req *v1.AddMemberReq) error
		RemoveMember(ctx context.Context, req *v1.RemoveMemberReq) error
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
