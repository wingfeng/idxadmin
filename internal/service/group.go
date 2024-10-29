// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/group/v1"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type (
	IGroup interface {
		Get(ctx context.Context, id int64) (*entity.Groups, error)
		Delete(ctx context.Context, id int64) error
		List(ctx context.Context, req *v1.ListReq) (*model.PageRes, error)
		Save(ctx context.Context, req *v1.SaveReq) error
		AddMember(ctx context.Context, req *v1.AddMemberReq) error
		Members(ctx context.Context, req *v1.MembersReq) (*model.PageRes, error)
		RemoveMember(ctx context.Context, req *v1.RemoveMemberReq) error
	}
)

var (
	localGroup IGroup
)

func Group() IGroup {
	if localGroup == nil {
		panic("implement not found for interface IGroup, forgot register?")
	}
	return localGroup
}

func RegisterGroup(i IGroup) {
	localGroup = i
}
