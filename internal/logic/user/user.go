package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/wingfeng/idx-oauth2/utils"
	idxConsts "github.com/wingfeng/idx/consts"
	v1 "github.com/wingfeng/idxadmin/api/user/v1"
	"github.com/wingfeng/idxadmin/internal/consts"
	"github.com/wingfeng/idxadmin/internal/dao"
	"github.com/wingfeng/idxadmin/internal/logic/common"
	"github.com/wingfeng/idxadmin/internal/model/entity"
	"github.com/wingfeng/idxadmin/internal/service"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}
func New() service.IUser {

	return &sUser{}
}

func (s *sUser) Get(ctx context.Context, id int64) (*entity.Users, error) {
	if row, err := dao.Users.Ctx(ctx).One("id", id); err != nil {
		return nil, err
	} else {

		result := &entity.Users{}
		err = row.Struct(result)
		return result, err
	}
}
func (s *sUser) Delete(ctx context.Context, id int64) error {
	_, err := dao.Users.Ctx(ctx).Delete("id=?", id)
	//todo: Clean Cache after reset;
	return err
}

func (s *sUser) Save(ctx context.Context, req *v1.SaveReq) (err error) {

	sub := ctx.Value(consts.SUBJECT_KEY)
	account := ctx.Value(consts.ACCOUNT_KEY)
	if account != nil && sub != nil {
		req.Updator = account.(string)
		req.UpdatorId = sub.(string)
	}
	req.UserName = strings.ToLower(req.UserName)
	result, err := dao.Users.Ctx(ctx).OnConflict("id").Save(req.Users)
	if err != nil {
		return err
	}
	if ra, er := result.RowsAffected(); ra == 0 || er != nil {
		return fmt.Errorf("save error ,no rows affected %s", er)
	}

	//todo: Clean Cache after reset;
	return err
}
func (s *sUser) List(ctx context.Context, req *v1.ListReq) (*v1.ListRes, error) {

	items := make([]entity.Users, 0)
	count := 0
	err := dao.Users.Ctx(ctx).Handler(common.Paginate(&req.PageReq)).ScanAndCount(&items, &count, true)
	res := &v1.ListRes{}
	res.PageSize = req.PageSize
	res.Page = req.Page
	res.Total = count
	res.List = items
	return res, err
}
func (s *sUser) ResetPwd(ctx context.Context, req *v1.ResetPwdReq) (string, error) {
	newPwd := utils.GenerateRandomString(8)
	newHash, _ := utils.HashPassword(newPwd)
	result, err := dao.Users.Ctx(ctx).Data(g.Map{"password_hash": newHash,
		"is_temporary_password": true}).Where("id=?", req.Id).Update()

	g.Log().Info(ctx, "User Reset Password", "User", req.Id, "row affected", result)
	if err == nil {

		u, _ := dao.Users.Ctx(ctx).One("id=?", req.Id)
		user := &entity.Users{}
		u.Struct(user)
		key := idxConsts.CONST_USERPWDHashKEY + user.UserName
		_, err = common.Cache.Remove(ctx, key)
		if err != nil {
			g.Log().Error(ctx, "User Reset Password", "User", req.Id, "remove pwd cache", err)
		}
		key = idxConsts.CONST_USERNAMEKEY + user.UserName
		_, err = common.Cache.Remove(ctx, key)
		if err != nil {
			g.Log().Error(ctx, "User Reset Password", "User", req.Id, "remove user cache", err)
		}

	}
	//todo: Clean Cache after reset;
	return newPwd, err
}
