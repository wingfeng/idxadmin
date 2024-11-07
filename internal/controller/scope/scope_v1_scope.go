package scope

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/scope/v1"
	"github.com/wingfeng/idxadmin/service"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	r, err := service.Scope().Get(ctx, req.Id)
	return &v1.GetRes{Scopes: *r}, err
}
func (c *ControllerV1) Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error) {
	err = service.Scope().Save(ctx, req)
	return nil, err
}
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	p, err := service.Scope().List(ctx, req)
	return &v1.ListRes{PageRes: *p}, err
}
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return nil, service.Scope().Delete(ctx, req.Id)
}
