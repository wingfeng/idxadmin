package orgunit

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/wingfeng/idxadmin/api/orgunit/v1"
	"github.com/wingfeng/idxadmin/internal/model/entity"
	"github.com/wingfeng/idxadmin/service"
)

func (c *ControllerV1) Tree(ctx context.Context, req *v1.TreeReq) (res *v1.TreeRes, err error) {
	items, err := service.OrgUnit().Tree(ctx)
	list := make([]*entity.OrganizationUnitsEx, 0)
	gconv.Scan(items, &list)
	res = &v1.TreeRes{
		Items: list,
	}
	return res, err
}
func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	r, err := service.OrgUnit().List(ctx, &req.PageReq)
	//	gconv.Scan(res, r)
	return &v1.ListRes{
		PageRes: *r,
	}, err
}
