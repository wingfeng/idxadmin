package common

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/service"
)

type sCRUD[do interface{}] struct {
}

func New[do interface{}]() service.ICRUD[do] {

	return &sCRUD[do]{}
}

func NewCRUD[do interface{}]() *sCRUD[do] { return &sCRUD[do]{} }

func (s *sCRUD[do]) Get(ctx context.Context, id interface{}) (do, error) {
	r := new(do)
	ro, err := g.DB().Model(r).One("id=?", id)
	ro.Struct(r)
	return *r, err
}
func (s *sCRUD[do]) Delete(ctx context.Context, id interface{}) error {
	r := new(do)
	_, err := g.DB().Model(r).Delete("id=?", id)
	return err
}
func (s *sCRUD[do]) Save(ctx context.Context, row do) error {
	_, err := g.DB().Model(row).OnConflict("id").Save(row)
	return err
}
func (s *sCRUD[do]) List(ctx context.Context, req *model.PageReq) (*model.PageRes, error) {
	items := make([]do, 0)
	count := 0
	tx := g.DB().Model(new(do))

	err := tx.Handler(Paginate(req)).ScanAndCount(&items, &count, false)
	res := &model.PageRes{}
	res.PageSize = req.PageSize
	res.Page = req.Page
	res.Total = count
	res.List = items
	return res, err

}
func Paginate(req *model.PageReq) func(m *gdb.Model) *gdb.Model {
	return func(tx *gdb.Model) *gdb.Model {
		for i, v := range req.Filters {
			arg := req.Args[i]
			tx = tx.Where(v, arg)
		}
		if !strings.EqualFold(req.Fields, "") {
			tx = tx.Fields(req.Fields)
		}
		if !strings.EqualFold(req.SortField, "") {
			tx = tx.Order(req.SortField, req.GetSortOrder())
		}

		return tx.Page(req.Page, req.PageSize)
	}
}
