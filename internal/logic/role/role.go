package role

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	v1 "github.com/wingfeng/idxadmin/api/role/v1"
	"github.com/wingfeng/idxadmin/internal/dao"
	"github.com/wingfeng/idxadmin/internal/logic/common"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/do"
	"github.com/wingfeng/idxadmin/internal/model/entity"
	"github.com/wingfeng/idxadmin/service"
)

type sRole struct {
	crud service.ICRUD[entity.Roles]
}

func init() {
	service.RegisterRole(New())
}
func New() service.IRole {
	crud := common.NewCRUD[entity.Roles]()
	return &sRole{
		crud: crud,
	}
}

func (s *sRole) Delete(ctx context.Context, id int64) error {
	err := dao.Roles.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Ctx(ctx).Model(&entity.UserRoles{}).Delete("role_id=?", id)
		if err != nil {
			return err
		}
		_, err = tx.Ctx(ctx).Model(&entity.Roles{}).Delete("id=?", id)
		return err
	})
	return err
}

func (s *sRole) Get(ctx context.Context, id int64) (entity.Roles, error) {
	return s.crud.Get(ctx, id)
}

func (s *sRole) List(ctx context.Context, req *v1.ListReq) (*model.PageRes, error) {
	return s.crud.List(ctx, &req.PageReq)
}
func (s *sRole) Save(ctx context.Context, req *entity.Roles) error {
	req.Name = strings.ToLower(req.Name)
	return s.crud.Save(ctx, *req)
}
func (s *sRole) Members(ctx context.Context, req *v1.MembersReq) (*model.PageRes, error) {
	items := make([]entity.Users, 0)
	count := 0
	tx := dao.Users.Ctx(ctx).InnerJoin("user_roles", "user_roles.user_id=users.id").Where("user_roles.role_id=?", req.Id)

	err := tx.Handler(common.Paginate(&req.PageReq)).ScanAndCount(&items, &count, false)
	res := &model.PageRes{}
	res.PageSize = req.PageSize
	res.Page = req.Page
	res.Total = count
	res.List = items
	return res, err
}

func (s *sRole) AddMemeber(ctx context.Context, req *v1.AddMemberReq) error {
	items := make([]*do.UserRoles, 0)
	for _, v := range req.UserIds {
		ur := &do.UserRoles{UserId: v, RoleId: req.Id}
		items = append(items, ur)
	}
	_, err := dao.UserRoles.Ctx(ctx).OnConflict("role_id", "user_id").Save(items)
	return err
}

func (s *sRole) RemoveMember(ctx context.Context, req *v1.RemoveMemberReq) error {
	_, err := dao.UserRoles.Ctx(ctx).Delete("role_id=? and user_id in (?)", req.Id, req.UserIds)
	return err
}
