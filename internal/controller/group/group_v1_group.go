package group

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/group/v1"
	"github.com/wingfeng/idxadmin/service"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	r, err := service.Group().Get(ctx, req.Id)
	return &v1.GetRes{
		Groups: *r,
	}, err
}
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	p, err := service.Group().List(ctx, req)
	return &v1.ListRes{
		PageRes: *p,
	}, err
}
func (c *ControllerV1) Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error) {
	return nil, service.Group().Save(ctx, req)
}
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return nil, service.Group().Delete(ctx, req.Id)
}
func (c *ControllerV1) Members(ctx context.Context, req *v1.MembersReq) (res *v1.MembersRes, err error) {
	p, err := service.Group().Members(ctx, req)
	return &v1.MembersRes{
		PageRes: *p,
	}, err
}
func (c *ControllerV1) AddMember(ctx context.Context, req *v1.AddMemberReq) (res *v1.AddMemberRes, err error) {
	return nil, service.Group().AddMember(ctx, req)
}
func (c *ControllerV1) RemoveMember(ctx context.Context, req *v1.RemoveMemberReq) (res *v1.RemoveMemberRes, err error) {
	return nil, service.Group().RemoveMember(ctx, req)
}
