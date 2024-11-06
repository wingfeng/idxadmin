package user

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/user/v1"
	"github.com/wingfeng/idxadmin/internal/service"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {

	row, err := service.User().Get(ctx, req.Id)
	return &v1.GetRes{
		Users: *row,
	}, err
}
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.User().Delete(ctx, req.Id)
	return nil, err
}
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	res, err = service.User().List(ctx, req)
	return res, err
}
func (c *ControllerV1) Save(ctx context.Context, req *v1.SaveReq) (res *v1.SaveRes, err error) {
	err = service.User().Save(ctx, req)
	return nil, err
}
func (c *ControllerV1) ResetPwd(ctx context.Context, req *v1.ResetPwdReq) (res *v1.ResetPwdRes, err error) {
	newpwd, err := service.User().ResetPwd(ctx, req)
	return &v1.ResetPwdRes{
		NewPwd: newpwd,
	}, err
}
