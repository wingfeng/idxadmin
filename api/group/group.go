// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package group

import (
	"context"

	"github.com/wingfeng/idxadmin/api/group/v1"
)

type IGroupV1 interface {
	Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Members(ctx context.Context, req *v1.MembersReq) (res *v1.MembersRes, err error)
	AddMember(ctx context.Context, req *v1.AddMemberReq) (res *v1.AddMemberRes, err error)
	RemoveMember(ctx context.Context, req *v1.RemoveMemberReq) (res *v1.RemoveMemberRes, err error)
}
