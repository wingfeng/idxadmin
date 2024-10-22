package user

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/wingfeng/idx-oauth2/utils"
	v1 "github.com/wingfeng/idxadmin/api/user/v1"
	"github.com/wingfeng/idxadmin/internal/dao"
	"github.com/wingfeng/idxadmin/internal/logic/common"
	"github.com/wingfeng/idxadmin/internal/model/do"
	"github.com/wingfeng/idxadmin/internal/service"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}
func New() service.IUser {
	slog.Info("Client service register")
	return &sUser{}
}

func (s *sUser) Get(ctx context.Context, id int64) (*do.Users, error) {
	if row, err := dao.Users.Ctx(ctx).One("id", id); err != nil {
		return nil, err
	} else {

		result := &do.Users{}
		err = row.Struct(result)
		return result, err
	}
}
func (s *sUser) Delete(ctx context.Context, id int64) error {
	_, err := dao.Users.Ctx(ctx).Delete("id=?", id)
	return err
}

func (s *sUser) Save(ctx context.Context, req v1.SaveReq) (err error) {

	result, err := dao.Users.Ctx(ctx).OnConflict("id").OmitEmptyData().Save(req.Users)
	if ra, er := result.RowsAffected(); ra == 0 || er != nil {
		return fmt.Errorf("save error ,no rows affected %s", er)
	}
	return err
}
func (s *sUser) List(ctx context.Context, req v1.PageReq) (*v1.PageRes, error) {

	items := make([]do.Users, 0)
	count := 0
	err := dao.Users.Ctx(ctx).Handler(common.Paginate(&req.PageReq)).ScanAndCount(&items, &count, true)
	res := &v1.PageRes{}
	res.PageSize = req.PageSize
	res.Page = req.Page
	res.Total = count
	res.List = items
	return res, err
}
func (s *sUser) ResetPwd(ctx context.Context, req v1.ResetPwdReq) (string, error) {
	newPwd := utils.GenerateRandomString(8)
	newHash, _ := utils.HashPassword(newPwd)
	result, err := dao.Users.Ctx(ctx).Data(g.Map{"password_hash": newHash,
		"is_temporary_password": true}).Where("id=?", req.Id).Update()

	g.Log().Info(ctx, "User Reset Password", "User", req.Id, "row affected", result)
	return newPwd, err
}
