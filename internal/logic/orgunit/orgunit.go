package orgunit

import (
	"context"

	"github.com/wingfeng/idxadmin/internal/dao"
	"github.com/wingfeng/idxadmin/internal/logic/common"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"
	"github.com/wingfeng/idxadmin/service"
)

type sOrgUnit struct {
	crud service.ICRUD[entity.OrganizationUnits]
}

func New() service.IOrgUnit {
	crud := common.NewCRUD[entity.OrganizationUnits]()
	return &sOrgUnit{
		crud: crud,
	}
}
func init() {
	service.RegisterOrgUnit(New())
}

func (s *sOrgUnit) Get(ctx context.Context, id int64) (entity.OrganizationUnits, error) {
	return s.crud.Get(ctx, id)
}
func (s *sOrgUnit) Delete(ctx context.Context, id int64) error {
	return s.crud.Delete(ctx, id)
}
func (s *sOrgUnit) Save(ctx context.Context, req entity.OrganizationUnits) error {
	return s.crud.Save(ctx, req)
}
func (s *sOrgUnit) List(ctx context.Context, req *model.PageReq) (*model.PageRes, error) {
	return s.crud.List(ctx, req)
}
func (s *sOrgUnit) Tree(ctx context.Context) ([]*entity.OrganizationUnitsEx, error) {
	result := make([]*entity.OrganizationUnitsEx, 0)
	all := make([]*entity.OrganizationUnitsEx, 0)
	re, err := dao.OrganizationUnits.Ctx(ctx).All(&all)
	if err != nil {
		return result, err
	}
	err = re.Structs(&all)
	if err != nil {
		return result, err
	}
	for i := 0; i < len(all); i++ {
		item := all[i]
		if item.ParentId == 0 {
			appendChild(all, item)
			result = append(result, item)
		}

	}
	return result, nil

}
func appendChild(all []*entity.OrganizationUnitsEx, parent *entity.OrganizationUnitsEx) {
	for i := 0; i < len(all); i++ {
		item := all[i]
		if item.ParentId == parent.Id {
			parent.Children = append(parent.Children, item)
			appendChild(all, item)
		}

	}
}
