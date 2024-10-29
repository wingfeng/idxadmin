package group

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	v1 "github.com/wingfeng/idxadmin/api/group/v1"
	"github.com/wingfeng/idxadmin/internal/dao"
	"github.com/wingfeng/idxadmin/internal/logic/common"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/do"
	"github.com/wingfeng/idxadmin/internal/model/entity"
	"github.com/wingfeng/idxadmin/internal/service"
)

type sGroup struct {
	crud service.ICRUD[entity.Groups]
}

func init() {
	service.RegisterGroup(New())
}
func New() service.IGroup {
	return &sGroup{crud: common.NewCRUD[entity.Groups]()}
}
func (s *sGroup) Get(ctx context.Context, id int64) (*entity.Groups, error) {
	r, err := s.crud.Get(ctx, id)
	return &r, err
}
func (s *sGroup) Delete(ctx context.Context, id int64) error {
	err := dao.UserGroups.DB().Transaction(context.TODO(), func(ctx context.Context, tx gdb.TX) error {
		_, err := tx.Ctx(ctx).Model(&entity.UserGroups{}).Delete("group_id=?", id)
		if err != nil {
			return err
		}
		_, err = tx.Ctx(ctx).Model(&entity.Groups{}).Delete("id=?", id)
		return err
	})
	return err
}
func (s *sGroup) List(ctx context.Context, req *v1.ListReq) (*model.PageRes, error) {
	return s.crud.List(ctx, &req.PageReq)

}
func (s *sGroup) Save(ctx context.Context, req *v1.SaveReq) error {
	return s.crud.Save(ctx, req.Groups)
}
func (s *sGroup) AddMember(ctx context.Context, req *v1.AddMemberReq) error {
	items := make([]*do.UserGroups, 0)
	for _, v := range req.UserIds {
		ur := &do.UserGroups{UserId: v, GroupId: req.Id}
		items = append(items, ur)
	}
	_, err := dao.UserGroups.Ctx(ctx).OnConflict("group_id", "user_id").Save(items)
	return err
}
func (s *sGroup) Members(ctx context.Context, req *v1.MembersReq) (*model.PageRes, error) {
	items := make([]entity.Users, 0)
	count := 0
	tx := dao.Users.Ctx(ctx).InnerJoin("user_groups", "user_groups.user_id=users.id").Where("user_groups.group_id=?", req.Id)

	err := tx.Handler(common.Paginate(&req.PageReq)).ScanAndCount(&items, &count, false)
	res := &model.PageRes{}
	res.PageSize = req.PageSize
	res.Page = req.Page
	res.Total = count
	res.List = items
	return res, err
}
func (s *sGroup) RemoveMember(ctx context.Context, req *v1.RemoveMemberReq) error {
	_, err := dao.UserGroups.Ctx(ctx).Delete("group_id=? and user_id in (?)", req.Id, req.UserIds)
	return err
}
