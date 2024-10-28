package scope

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/wingfeng/idxadmin/api/scope/v1"
	"github.com/wingfeng/idxadmin/internal/consts"
	"github.com/wingfeng/idxadmin/internal/dao"
	"github.com/wingfeng/idxadmin/internal/logic/common"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/do"
	"github.com/wingfeng/idxadmin/internal/model/entity"
	"github.com/wingfeng/idxadmin/internal/service"
)

type sScope struct {
	crud service.ICRUD[entity.Scopes]
}

func init() {
	service.RegisterScope(New())
}
func New() service.IScope {
	return &sScope{crud: common.NewCRUD[entity.Scopes]()}
}
func (s *sScope) Get(ctx context.Context, id int64) (res *entity.Scopes, err error) {
	r, err := s.crud.Get(ctx, id)
	return &r, err
}
func (s *sScope) Delete(ctx context.Context, id int64) error {
	r, err := s.crud.Get(ctx, id)
	if err != nil {
		return err
	}
	if isEssential(&r) {
		return fmt.Errorf("essential scope can not be deleted")
	}
	return s.crud.Delete(ctx, id)
}
func (s *sScope) List(ctx context.Context, req *v1.ListReq) (*model.PageRes, error) {
	items := make([]entity.Scopes, 0)
	count := 0
	tx := dao.Scopes.Ctx(ctx)

	err := tx.Handler(common.Paginate(&req.PageReq)).ScanAndCount(&items, &count, false)

	res := &model.PageRes{}
	res.PageSize = req.PageSize
	res.Page = req.Page
	res.Total = count
	res.List = items
	return res, err
}
func (s *sScope) Save(ctx context.Context, req *v1.SaveReq) error {
	req.Name = strings.ToLower(req.Name)
	if req.Id == 0 {
		d := &do.Scopes{}

		gconv.Struct(req.Scopes, d)
		d.Id = nil
		_, err := dao.Scopes.Ctx(ctx).Insert(d)
		return err
	} else {
		if isEssential(&req.Scopes) {
			return fmt.Errorf("essential scope can not be modified")
		}
		return s.crud.Save(ctx, req.Scopes)
	}

}

func isEssential(s *entity.Scopes) bool {

	r := s.Properties.Get(consts.Essential_KEY).Bool()
	return r
}
