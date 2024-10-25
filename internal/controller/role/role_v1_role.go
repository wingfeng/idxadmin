package role

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/role/v1"
	"github.com/wingfeng/idxadmin/internal/service"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	r, err := service.Role().Get(ctx, req.Id)
	return &v1.GetRes{
		Roles: r,
	}, err
}
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	p, err := service.Role().List(ctx, req)
	return &v1.ListRes{
		PageRes: *p,
	}, err
}
func (c *ControllerV1) Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error) {
	err = service.Role().Save(ctx, &req.Roles)
	return nil, err
}
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	return nil, err
}
func (c *ControllerV1) Members(ctx context.Context, req *v1.MembersReq) (res *v1.MembersRes, err error) {
	p, err := service.Role().Members(ctx, req)
	return &v1.MembersRes{
		PageRes: *p,
	}, err
}
