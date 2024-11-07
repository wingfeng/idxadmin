// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/wingfeng/idxadmin/api/scope/v1"
	"github.com/wingfeng/idxadmin/internal/model"
	"github.com/wingfeng/idxadmin/internal/model/entity"
)

type (
	IScope interface {
		Get(ctx context.Context, id int64) (res *entity.Scopes, err error)
		Delete(ctx context.Context, id int64) error
		List(ctx context.Context, req *v1.ListReq) (*model.PageRes, error)
		Save(ctx context.Context, req *v1.SaveReq) error
	}
)

var (
	localScope IScope
)

func Scope() IScope {
	if localScope == nil {
		panic("implement not found for interface IScope, forgot register?")
	}
	return localScope
}

func RegisterScope(i IScope) {
	localScope = i
}
