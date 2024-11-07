// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type (
	IOrgUnit interface {
		Get(ctx context.Context, id int64) (entity.OrganizationUnits, error)
		Delete(ctx context.Context, id int64) error
		Save(ctx context.Context, req entity.OrganizationUnits) error
		List(ctx context.Context, req *model.PageReq) (*model.PageRes, error)
		Tree(ctx context.Context) ([]*entity.OrganizationUnitsEx, error)
	}
)

var (
	localOrgUnit IOrgUnit
)

func OrgUnit() IOrgUnit {
	if localOrgUnit == nil {
		panic("implement not found for interface IOrgUnit, forgot register?")
	}
	return localOrgUnit
}

func RegisterOrgUnit(i IOrgUnit) {
	localOrgUnit = i
}
